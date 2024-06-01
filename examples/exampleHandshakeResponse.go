//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func main() {
	handshake, err := common.NewHandshake(generated.RM, []string{"1.0"})
	if err != nil {
		log.Fatalf("Error creating the handhake %v", err)
	}

	handshakeResponse, err := common.NewHandshakeResponse(handshake)
	if err != nil {
		log.Fatalf("Error creatiing the handshake response %v", err)
	}

	// Serialize the message

	handshakeResponseJson, err := json.MarshalIndent(handshakeResponse, "", " ")
	if err != nil {
		log.Fatalf("Error creating the JSON handshake response: %v", err)
	}

	fmt.Printf("Handshake Response: %s\n", handshakeResponseJson)

}

/* OUTPUT:

Handshake Response: {
 "message_id": {                                                                                         "value": "0a8da1624ba44447aaab8ed6ec0f4bdf"                                                           },
"message_type": "HandshakeResponse",                                                                   "selected_protocol_version": [                                                                          "1.0"                                                                                                 ]                                                                                                     }
*/
