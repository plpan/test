package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

func main() {
	const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)
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
}
