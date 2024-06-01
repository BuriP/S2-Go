package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
	"github.com/BuriP/S2-Go/websockets/S2client"
	"github.com/BuriP/S2-Go/websockets/S2server"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func mockServer() *httptest.Server {
	server := S2server.NewWebSocketServer()
	return httptest.NewServer(http.HandlerFunc(server.HandleConnections))
}

func TestWebSocketClient(t *testing.T) {
	server := mockServer()
	defer server.Close()

	addr := server.Listener.Addr().String()

	client, err := S2client.NewWebSocketClient(addr, "/ws")
	assert.NoError(t, err, "Should create a new WebSocket client")

	defer client.Close()

	// Create an instance of a Handshake message
	role := generated.CEM
	supportedProtocols := []string{"1.0", "1.1"}
	handshake, err := common.NewHandshake(role, supportedProtocols)
	assert.NoError(t, err, "Should create a new Handshake message")

	// Send the Handshake message
	err = client.SendMessage(handshake)
	assert.NoError(t, err, "Should send Handshake message without error")

	// Receive a response
	receivedMessage, err := client.ReceiveMessage()
	assert.NoError(t, err, "Should receive message without error")
	assert.Contains(t, receivedMessage, "Handshake", "Received message should contain 'Handshake'")

	client.Close()
}

func TestWebSocketClientTimeout(t *testing.T) {
	server := mockServer()
	defer server.Close()

	addr := server.Listener.Addr().String()

	client, err := S2client.NewWebSocketClient(addr, "/ws")
	assert.NoError(t, err, "Should create a new WebSocket client")

	defer client.Close()

	client.GetConn().SetReadDeadline(time.Now().Add(1 * time.Second))

	_, err = client.ReceiveMessage()
	assert.Error(t, err, "Should receive a timeout error")
}
