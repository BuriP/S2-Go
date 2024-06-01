package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common" // Ensure this is the correct import path
	"github.com/BuriP/S2-Go/src/frbc"
	"github.com/BuriP/S2-Go/src/generated"
	"github.com/BuriP/S2-Go/websockets/client"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

// mockServer handles the WebSocket connection for testing.
func mockServer(t *testing.T, handler func(*websocket.Conn)) *httptest.Server {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade connection: %v", err)
		}
		handler(conn)
	}))
}

// TestWebSocketClient tests the WebSocketClient send and receive functionality.
func TestWebSocketClient(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	server := mockServer(t, func(conn *websocket.Conn) {
		defer conn.Close()
		defer wg.Done()

		for {
			// Expect to receive a message
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}

			var messageType struct {
				MessageType string `json:"message_type"`
			}
			err = json.Unmarshal(message, &messageType)
			assert.NoError(t, err)

			switch messageType.MessageType {
			case "Handshake":
				var handshake common.Handshake
				err = json.Unmarshal(message, &handshake)
				assert.NoError(t, err)

				// Create a response and send it back to the client
				response, err := common.NewHandshakeResponse(&handshake)
				assert.NoError(t, err)

				responseData, err := json.Marshal(response)
				assert.NoError(t, err)
				err = conn.WriteMessage(websocket.TextMessage, responseData)
				assert.NoError(t, err)

			case "FRBC.Instruction":
				var instruction frbc.FRBCInstruction
				err = json.Unmarshal(message, &instruction)
				assert.NoError(t, err)

				// Echo the instruction back to the client
				err = conn.WriteMessage(websocket.TextMessage, message)
				assert.NoError(t, err)

			case "FRBCActuatorStatus":
				var status frbc.FRBCActuatorStatus
				err = json.Unmarshal(message, &status)
				assert.NoError(t, err)

				// Echo the status back to the client
				err = conn.WriteMessage(websocket.TextMessage, message)
				assert.NoError(t, err)
			}
		}
	})
	defer server.Close()

	// Correctly format the WebSocket URL
	wsURL := "ws" + server.URL[len("http"):]

	// Create WebSocket client
	wsClient, err := client.NewWebSocketClient(wsURL)
	assert.NoError(t, err)
	if wsClient == nil {
		t.Fatalf("Failed to create WebSocket client")
	}
	defer wsClient.Close()

	// Create and send a Handshake message
	role := generated.CEM
	supportedProtocols := []string{"1.0", "2.0"}
	handshake, err := common.NewHandshake(role, supportedProtocols)
	assert.NoError(t, err)

	err = wsClient.SendMessage(handshake)
	assert.NoError(t, err)

	var handshakeResponse common.HandshakeResponse
	err = wsClient.ReceiveMessage(&handshakeResponse)
	assert.NoError(t, err)
	assert.Equal(t, "HandshakeResponse", handshakeResponse.MessageType)
	assert.Equal(t, supportedProtocols, *handshakeResponse.SelectedProtocolVersion)

	// Create and send a FRBCInstruction message
	actuatorID, _ := generated.NewID()
	operationModeID, _ := generated.NewID()
	executionTime := time.Now()
	instruction, err := frbc.NewFRBCInstruction(false, actuatorID, operationModeID, executionTime, 0.5)
	assert.NoError(t, err)

	err = wsClient.SendMessage(instruction)
	assert.NoError(t, err)

	var receivedInstruction frbc.FRBCInstruction
	err = wsClient.ReceiveMessage(&receivedInstruction)
	assert.NoError(t, err)
	assert.Equal(t, instruction.MessageType, receivedInstruction.MessageType)
	assert.Equal(t, instruction.ActuatorID, receivedInstruction.ActuatorID)
	assert.Equal(t, instruction.OperationMode, receivedInstruction.OperationMode)
	assert.Equal(t, instruction.OperationModeFactor, receivedInstruction.OperationModeFactor)

	// Create and send a FRBCActuatorStatus message
	status, err := frbc.NewFRBCActuatorStatus(actuatorID, actuatorID, 0.8, operationModeID, &executionTime)
	assert.NoError(t, err)

	err = wsClient.SendMessage(status)
	assert.NoError(t, err)

	var receivedStatus frbc.FRBCActuatorStatus
	err = wsClient.ReceiveMessage(&receivedStatus)
	assert.NoError(t, err)
	assert.Equal(t, status.MessageType, receivedStatus.MessageType)
	assert.Equal(t, status.ActuatorID, receivedStatus.ActuatorID)
	assert.Equal(t, status.OperationModeFactor, receivedStatus.OperationModeFactor)
	assert.Equal(t, status.ActiveOperationModeID, receivedStatus.ActiveOperationModeID)

	wg.Wait() // Wait for the server to complete handling the connections
}
