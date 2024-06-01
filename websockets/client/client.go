package client

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	conn *websocket.Conn
}

// NewWebSocketClient creates a new WebSocket client and connects to the given URL.
func NewWebSocketClient(url string) (*WebSocketClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return &WebSocketClient{conn: conn}, nil
}

// SendMessage sends a message over the WebSocket connection.
func (c *WebSocketClient) SendMessage(message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return c.conn.WriteMessage(websocket.TextMessage, data)
}

// ReceiveMessage receives a message from the WebSocket connection and unmarshals it into the provided target.
func (c *WebSocketClient) ReceiveMessage(target interface{}) error {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return err
	}
	return json.Unmarshal(message, target)
}

// Close closes the WebSocket connection.
func (c *WebSocketClient) Close() error {
	return c.conn.Close()
}
