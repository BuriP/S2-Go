package s2server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

// WebSocketServer represents a WebSocket server.
type WebSocketServer struct {
	upgrader websocket.Upgrader
	clients  map[*websocket.Conn]bool
	mu       sync.Mutex // Mutex to synchronize access to clients
}

// NewWebSocketServer creates and returns a new WebSocketServer instance.
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

// HandleConnections handles incoming WebSocket connections.
func (server *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer func() {
		server.mu.Lock()
		defer server.mu.Unlock()
		conn.Close()
		delete(server.clients, conn)
	}()

	server.mu.Lock()
	server.clients[conn] = true
	server.mu.Unlock()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Printf("Connection closed normally: %v", err)
			} else if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error: %v", err)
			} else {
				log.Printf("Error reading message: %v", err)
			}
			break
		}
		log.Printf("Received message: %s", message)

		// If the message is a close message, log it
		if messageType == websocket.CloseMessage {
			log.Printf("Received close message")
			break
		}

		server.mu.Lock()
		for client := range server.clients {
			if err := client.WriteMessage(messageType, message); err != nil {
				log.Printf("Error writing message: %v", err)
				client.Close()
				delete(server.clients, client)
			}
		}
		server.mu.Unlock()
	}
}

// Run starts the WebSocket server and listens on the specified address.
func (server *WebSocketServer) Run(addr string) {
	http.HandleFunc("/ws", server.HandleConnections)
	log.Printf("WebSocket server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error starting WebSocket server: %v", err)
	}
}
