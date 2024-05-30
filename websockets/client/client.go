package client

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

// WebSocketClient wraps a WebSocket connection.
type WebSocketClient struct {
	conn *websocket.Conn
}

// Connect establishes a WebSocket connection to the given URL.
func Connect(url string) (*WebSocketClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	return &WebSocketClient{conn: conn}, nil
}

// SendJSON sends a JSON-encoded message over the WebSocket connection.
func (c *WebSocketClient) SendJSON(message interface{}) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if err := c.conn.WriteMessage(websocket.TextMessage, messageJSON); err != nil {
		return err
	}

	return nil
}

// ReceiveJSON receives a JSON-encoded message from the WebSocket connection.
func (c *WebSocketClient) ReceiveJSON(v interface{}) error {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(message, v); err != nil {
		return err
	}

	return nil
}

// Close closes the WebSocket connection.
func (c *WebSocketClient) Close() error {
	return c.conn.Close()
}
