package S2server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	upgrader websocket.Upgrader
	clients  map[*websocket.Conn]bool
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		clients: make(map[*websocket.Conn]bool),
	}
}

func (server *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	server.clients[conn] = true

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error: %v", err)
			}
			delete(server.clients, conn)
			break
		}
		log.Printf("Received message: %s", message)

		for client := range server.clients {
			if err := client.WriteMessage(messageType, message); err != nil {
				log.Printf("Error writing message: %v", err)
				client.Close()
				delete(server.clients, client)
			}
		}
	}
}

func (server *WebSocketServer) Run(addr string) {
	http.HandleFunc("/ws", server.HandleConnections)
	log.Printf("WebSocket server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error starting WebSocket server: %v", err)
	}
}
