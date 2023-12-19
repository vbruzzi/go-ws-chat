package main

import (
	"log"
	"net/http"
)

var addr = ":8080"

func serveHome(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "pages/home.html")
}

func serveClient(hub *Hub, w http.ResponseWriter, r *http.Request) {
	err := createConnection(hub, w, r)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	hub := newHub()
	go hub.init()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveClient(hub, w, r)
	})

	server := http.Server{
		Addr: addr,
	}

	log.Printf("Listening on localhost%v...", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("%v", err)
	}
}
