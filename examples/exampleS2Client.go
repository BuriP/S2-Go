//go:build ignore

// + build ignore
package main

import (
	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
	"github.com/BuriP/S2-Go/websockets/s2client"
	"log"
	"time"
)

func main() {
	// Replace with the address and path of your WebSocket server
	addr := "localhost:8080" // For the purpose of the example localhost is used.
	path := "/ws"

	client, err := s2client.NewWebSocketClient(addr, path)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer client.Close()

	// Create a sample message to send
	messageID, _ := generated.NewID()
	statusValue := generated.OK
	receptionStatus, err := common.NewReceptionStatus(statusValue, messageID, nil)
	if err != nil {
		log.Fatalf("Failed to create reception status message: %v", err)
	}

	// Send the message
	if err := client.SendMessage(receptionStatus); err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Receive a response
	response, err := client.ReceiveMessage()
	if err != nil {
		log.Fatalf("Failed to receive message: %v", err)
	}

	log.Printf("Received response: %s", response)

	// Add a delay to ensure the connection stays open long enough to receive a response
	time.Sleep(2 * time.Second)
}
