package common

import (

	"S2-Go/src/generated"
	"github.com/google/uuid"
)

// NewHandshake created a new Handshake instance.
func NewHandshake() *generated.Handshake {
	return &generated.Handshake{}
}


func SetMessageID(h *generated.Handshake) {
	// Generating a new UUID
	newUUID := uuid.New()

	// Assigning the new UUID to the MessageId field.
	h.MessageID = generated.ID(newUUID.String())
}

