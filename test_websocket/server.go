package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	gorilla "github.com/gorilla/websocket"
	golang "golang.org/x/net/websocket"
)

type msg struct {
	Num int
}

func main() {
	http.HandleFunc("/gorilla", gorillaWSHandler)
	http.Handle("/golang", golang.Handler(echoGolang))
	http.HandleFunc("/", handler)

	panic(http.ListenAndServe(":8010", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Open file failed: %s", err)
	} else {
		fmt.Fprintf(w, "%s", content)
	}
}

// x/net/websocket handler func
func echoGolang(conn *golang.Conn) {
	io.Copy(conn, conn)
}

// gorilla/websocket handler func
func gorillaWSHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}

	conn, err := gorilla.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	go echoGorilla(conn)
}

func echoGorilla(conn *gorilla.Conn) {
	for {
		m := msg{}
		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Read json failed: %s", err)
		}

		fmt.Printf("Got message: %#v\n", m)

		if err = conn.WriteJSON(m); err != nil {
			fmt.Println("Write json failed: %s", err)
		}
	}
}
