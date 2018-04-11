package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	"time"
)

const (
	master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
	overlay = `{{define "list"}} {{join . ", " | println}}{{end}}`
	email   = `{{with .Account -}}
Dear {{.FirstName}} {{.LastName}},
{{- end}}

Below are your account statement details for period from {{.FromDate | formatAsDate}} to {{.ToDate | formatAsDate}}.

{{if .Purchases -}}
    Your purchases:
    {{- range .Purchases }}
        {{ .Date | formatAsDate}} {{ printf "%-20s" .Description }} {{.AmountInCents | formatAsDollars -}}
    {{- end}}
{{- else}}
You didn't make any purchases during the period.
{{- end}}

{{$note := urgentNote .Account -}}
{{if $note -}}
Note: {{$note}}
{{- end}}

Best Wishes,
Customer Service`
)

type Account struct {
	FirstName string
	LastName  string
}

type Purchase struct {
	Date          time.Time
	Description   string
	AmountInCents int
}

type Statement struct {
	FromDate  time.Time
	ToDate    time.Time
	Account   Account
	Purchases []Purchase
}

func formatAsDollars(valueInCents int) (string, error) {
	dollars := valueInCents / 100
	cents := valueInCents % 100
	return fmt.Sprintf("$%d.%2d", dollars, cents), nil
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%d/%d", day, month, year)
}

func urgentNote(acc Account) string {
	return fmt.Sprintf("You have earned 100 VIP points that can be used for purchases")
}

func createMockStatement() Statement {
	return Statement{
		FromDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
		ToDate:   time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC),
		Account: Account{
			FirstName: "John",
			LastName:  "Dow",
		},
		Purchases: []Purchase{
			Purchase{
				Date:          time.Date(2016, 1, 3, 0, 0, 0, 0, time.UTC),
				Description:   "Shovel",
				AmountInCents: 2326,
			},
			Purchase{
				Date:          time.Date(2016, 1, 8, 0, 0, 0, 0, time.UTC),
				Description:   "Staple remover",
				AmountInCents: 5432,
			},
		},
	}
}

func main() {
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	// parse normal template
	masterTmpl, err := template.New("master").Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	// parse template with functions, the newer parser will overwrite the older parser
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Funcs(funcs).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	// execute normal template
	fmt.Println("-- normal template --")
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- template with function --")
	// execute template with functions
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- complicated template --")
	fmap := template.FuncMap{
		"formatAsDollars": formatAsDollars,
		"formatAsDate":    formatAsDate,
		"urgentNote":      urgentNote,
	}
	t := template.Must(template.New("email.tmpl").Funcs(fmap).Parse(email))
	err = t.Execute(os.Stdout, createMockStatement())
	if err != nil {
		panic(err)
	}
}
