package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	hub     *Hub
	conn    *websocket.Conn
	inbound chan []byte
}

func newClient(h *Hub, c *websocket.Conn) *Client {
	return &Client{
		h,
		c,
		make(chan []byte),
	}
}

func (c *Client) writePump() error {
	for {
		select {
		case message := <-c.inbound:
			writer, err := c.conn.NextWriter(websocket.TextMessage)

			if err != nil {
				return err
			}

			writer.Write(message)

			if len(c.inbound) > 0 {
				for m := range c.inbound {
					writer.Write(m)
				}
			}
			if err := writer.Close(); err != nil {
				return err
			}
		}
	}
}

func (c *Client) readPump() error {
	for {
		_, p, err := c.conn.ReadMessage()

		if err != nil {
			return err
		}
		c.hub.broadcast <- p
	}
}
