package S2client

import (
	"encoding/json"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
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
		return "", err
	}
	return string(message), nil
}

// Close method closes the connection for a given client.
func (client *WebSocketClient) Close() {
	if err := client.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		log.Printf("Error sending close message: %v", err)
	}
	if err := client.conn.Close(); err != nil {
		log.Printf("Error closing connection: %v", err)
	}
}

// GetConn provides access to the WebSocket connection (for testing purposes)
func (client *WebSocketClient) GetConn() *websocket.Conn {
	return client.conn
}
