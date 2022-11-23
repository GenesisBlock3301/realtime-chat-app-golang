# Realtime Chatapp Golang
# Setup & Instalization:
```
mkdir backend
cd backend
go mod init github.com/GenesisBlock3301/realtime-chat-app-golang/

create main.go
go run main.go
mkdir frontend

```
Packages:
```
"github.com/gin-gonic/gin"
"github.com/gorilla/websocket"
"log"
"net/http"
```
# For handling Single client:
```go
//THis is for only one client
for {
	var message Message
	err := ws.ReadJSON(&message)
	if err != nil {
		log.Printf("Error occured in : %v\n", err)
		break
	}
	log.Println(message)
	//	send message from server
	if err := ws.WriteJSON(message); err != nil {
		log.Printf("Error occured in : %v\n", err)
		}
}
```

# For handling multiple client:
```go
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


```
# Some understanding on Package:

`http.ResponseWriter`  This is the mechanism used for sending responses to any connected HTTP clients. 
It's also how response headers are set<br>
`w.Write([]byte("<h1>Welcome to my web server!</h1>"))`

The ``http.ResponseWriter`` interface has a Write method which accepts a byte slice and writes the data to the connection as part of an HTTP response. Converting a string to a byte slice is as easy as using []byte(str), and that's how we're able to respond to HTTP requests.
