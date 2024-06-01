package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func main() {
	handshake, err := common.NewHandshake(generated.RM, []string{"1.0", "2.0"})
	if err != nil {
		log.Fatalf("Error creating handshake: %v", err)
	}

	// Print the handshake as JSON for better readability
	handshakeJSON, err := json.MarshalIndent(handshake, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling handshake to JSON: %v", err)
	}

	fmt.Printf("Handshake: %s\n", handshakeJSON)
}
