package s2client

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

type WebSocketClient struct {
	conn *websocket.Conn
}

// NewWebSocketClient creates a new instance of the client.
func NewWebSocketClient(addr, path string) (*WebSocketClient, error) {
	u := url.URL{Scheme: "ws", Host: addr, Path: path}
	log.Printf("Connecting to %s", u.String())

	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	c, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	return &WebSocketClient{conn: c}, nil
}

// SendMessage method sends a given message for the selected client.
func (client *WebSocketClient) SendMessage(message interface{}) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	log.Printf("Sending message: %s", string(messageJSON))

	err = client.conn.WriteMessage(websocket.TextMessage, messageJSON)
	if err != nil {
		return err
	}

	return nil
}

// ReceiveMessage method receives a response message from the server.
func (client *WebSocketClient) ReceiveMessage() (string, error) {
	client.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	_, message, err := client.conn.ReadMessage()
	if err != nil {
		if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
			log.Printf("Connection closed normally")
			return "", nil
		}
		return "", err
	}
	return string(message), nil
}

// Close method closes the connection for a given client.
func (client *WebSocketClient) Close() {
	// Send a close message to the server
	closeMessage := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Closing the connection")
	if err := client.conn.WriteMessage(websocket.CloseMessage, closeMessage); err != nil {
		log.Printf("Error sending close message: %v", err)
	}
	time.Sleep(1 * time.Second)
	if err := client.conn.Close(); err != nil {

		log.Printf("Error closing connection: %v", err)
	}
}

// GetConn provides access to the WebSocket connection (for testing purposes)
func (client *WebSocketClient) GetConn() *websocket.Conn {
	return client.conn
}
