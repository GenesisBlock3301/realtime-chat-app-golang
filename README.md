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

# Some understanding on Package:

`http.ResponseWriter`  This is the mechanism used for sending responses to any connected HTTP clients. 
It's also how response headers are set<br>
`w.Write([]byte("<h1>Welcome to my web server!</h1>"))`

The ``http.ResponseWriter`` interface has a Write method which accepts a byte slice and writes the data to the connection as part of an HTTP response. Converting a string to a byte slice is as easy as using []byte(str), and that's how we're able to respond to HTTP requests.
