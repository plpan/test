package main

import "net/http"

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func main() {
	http.Handle("/", &Handler{})
	http.ListenAndServe(":8080", nil)
}
