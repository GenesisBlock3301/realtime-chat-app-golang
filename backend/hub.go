package main

// create this file for holds all clients information and send a message to them.
//The hub struct stores a client and sends a message through channels.
//As the channel gets the message, it will write back to all the clients.
import (
	"github.com/gorilla/websocket"
	"log"
)

type Hub struct {
	clients   map[*websocket.Conn]bool
	broadcast chan Message
}

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan Message),
	}
}

func (h *Hub) run() {
	for {
		select {
		case message := <-h.broadcast:
			for client := range h.clients {
				if err := client.WriteJSON(message); err != nil {
					log.Printf("error occurred: %v", err)
				}
			}

		}
	}
}
