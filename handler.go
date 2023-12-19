package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  readBufferSize,
	WriteBufferSize: writeBufferSize,
}

func createConnection(hub *Hub, w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return err
	}

	client := newClient(hub, conn)
	hub.register <- client

	go client.writePump()
	go client.readPump()

	return nil
}
