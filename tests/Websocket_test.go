package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/frbc"
	"github.com/BuriP/S2-Go/src/generated"
	"github.com/BuriP/S2-Go/websockets/client"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Error upgrading connection: %v", err)
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		// Deserialize JSON message to FRBCActuatorStatus
		var actuatorStatus frbc.FRBCActuatorStatus
		if err := json.Unmarshal(message, &actuatorStatus); err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			continue
		}

		log.Printf("Received message: %+v", actuatorStatus)

		// Create a response
		response := struct {
			Message         string                  `json:"message"`
			ReceivedMessage frbc.FRBCActuatorStatus `json:"received_message"`
		}{
			Message:         "message received",
			ReceivedMessage: actuatorStatus,
		}

		responseMessage, _ := json.Marshal(response)
		if err := conn.WriteMessage(messageType, responseMessage); err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
	}
}

func TestWebSocketCommunication(t *testing.T) {
	// Start a mock WebSocket server
	server := httptest.NewServer(http.HandlerFunc(handleConnections))
	defer server.Close()

	// Convert the server URL to WebSocket URL
	u := "ws" + server.URL[4:] + "/ws"

	// Connect to the mock server
	client, err := client.Connect(u)
	if err != nil {
		t.Fatalf("Error connecting to WebSocket server: %v", err)
	}
	defer client.Close()

	// Create a sample FRBCActuatorStatus message
	activeOperationModeID, _ := generated.NewID()
	actuatorID, _ := generated.NewID()
	previousOperationModeID, _ := generated.NewID()
	transitionTimestamp := time.Now()
	messageID, _ := generated.NewID()

	message := &frbc.FRBCActuatorStatus{
		ActiveOperationModeID:   activeOperationModeID,
		ActuatorID:              actuatorID,
		MessageID:               messageID,
		MessageType:             "FRBCActuatorStatus",
		OperationModeFactor:     1.0,
		PreviousOperationModeID: previousOperationModeID,
		TransitionTimestamp:     &transitionTimestamp,
	}

	// Send the message
	if err := client.SendJSON(message); err != nil {
		t.Fatalf("Error sending message: %v", err)
	}

	// Receive a response
	var response struct {
		Message         string                  `json:"message"`
		ReceivedMessage frbc.FRBCActuatorStatus `json:"received_message"`
	}
	if err := client.ReceiveJSON(&response); err != nil {
		t.Fatalf("Error receiving response: %v", err)
	}

	// Validate the response
	if response.Message != "message received" {
		t.Errorf("Expected message 'message received', got %v", response.Message)
	}
	if response.ReceivedMessage != *message {
		t.Errorf("Expected received message %+v, got %+v", *message, response.ReceivedMessage)
	}
}

