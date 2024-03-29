package common

import (

	"../generated"
	"github.com/google/uuid"
)


func NewHandshake() *generated.Handshake {
	return &generated.Handshake{}
}

func (h *generated.Handshake) ValidateAssignment() bool {
	return true
}

func (h *generated.Handshake) MessageID() uuid.UUID {
	return uuid.New()
}

func main() {
	// Sample usage
	handshake := NewHandshake()
	messageID := handshake.MessageID()
	println("Message ID:", messageID.String())

	// Validate assignment
	validateAssignment := handshake.ValidateAssignment()
	println("Validate Assignment:", validateAssignment)
}
