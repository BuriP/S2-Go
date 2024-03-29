package common

import (
	"../generated"
	"github.com/google/uuid"
)

// InstructionStatusUpdate represents an instruction status update message
type InstructionStatusUpdate struct {
	*generated.InstructionStatusUpdate
}

// Config represents configuration for InstructionStatusUpdate
type Config struct {
	*generated.InstructionStatusUpdateConfig
	ValidateAssignment bool // Additional configuration
}

// NewInstructionStatusUpdate creates a new InstructionStatusUpdate instance
func NewInstructionStatusUpdate() *InstructionStatusUpdate {
	return &InstructionStatusUpdate{
		InstructionStatusUpdate: &generated.InstructionStatusUpdate{},
	}
}

// MessageID returns the message ID field info
func (is *InstructionStatusUpdate) MessageID() uuid.UUID {
	return is.InstructionStatusUpdate.MessageID
}

// InstructionID returns the instruction ID field info
func (is *InstructionStatusUpdate) InstructionID() uuid.UUID {
	return is.InstructionStatusUpdate.InstructionID
}

