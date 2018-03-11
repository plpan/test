package main

import (
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile(`stupig`)
	name := r.URL.Path[1:]
	res := pattern.FindSubmatch([]byte(name))

	if len(res) > 0 {
		w.Write([]byte("Hello, World"))
	} else {
		w.Write([]byte("None"))
	}
}

func main() {
	http.Handle("/", &Handler{})
	http.ListenAndServe(":8080", nil)
}
