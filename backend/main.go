package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Message string `json:"message"`
}

func main() {
	router := gin.Default()
	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Good morning"})
	})

	//Create a hub
	hub := NewHub()

	go hub.run()

	router.GET("/ws", func(context *gin.Context) {
		// A CheckOrigin function should carefully validate the request origin to
		// prevent cross-site request forgery.
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		// If the upgrade fails, then Upgrade replies to the client with an HTTP error response.
		ws, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer func(ws *websocket.Conn) {
			delete(hub.clients, ws)
			err := ws.Close()
			log.Println("closed")
			if err != nil {
				log.Println(err)
			}
		}(ws)

		//Add client
		hub.clients[ws] = true

		log.Println("Connected")
		read(hub, ws)
		// THis is for only one client
		//for {
		//	var message Message
		//	err := ws.ReadJSON(&message)
		//	if err != nil {
		//		log.Printf("Error occured in : %v\n", err)
		//		break
		//	}
		//	log.Println(message)
		//
		//	//	send message from server
		//	if err := ws.WriteJSON(message); err != nil {
		//		log.Printf("Error occured in : %v\n", err)
		//	}
		//}
	})
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
func read(hub *Hub, client *websocket.Conn) {
	for {
		var message Message
		err := client.ReadJSON(&message)
		if err != nil {
			log.Println(err)
			delete(hub.clients, client)
			break
		}
		log.Println(message)
		hub.broadcast <- message
	}
}
