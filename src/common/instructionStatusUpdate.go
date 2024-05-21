package common

import (
	"github.com/BuriP/S2-Go/src/generated"
	"time"
)

// InstructionStatusUpdate represents an instruction status update message

type InstructionStatusUpdate struct {
	InstructionID *generated.ID                `json:"instruction_id" description:"ID of this instruction (as provided by the CEM)"`
	MessageID     *generated.ID                `json:"message_id" description:"ID of this message"`
	MessageType   string                       `json:"message_type" description:"Type of the message" const:"true"`
	StatusType    *generated.InstructionStatus `json:"status_type" description:"Present status of this instruction."`
	Timestamp     time.Time                    `json:"timestamp" description:"Timestamp when status_type has changed the last time."`
}

// NewInstructionStatusUpdate creates a new instance of InstructionStatusUpdate.
func GetInstructionStatusUpdate(messageType string, statusType *generated.InstructionStatus) (*InstructionStatusUpdate, error) {
	id, err := generated.NewID()
	if err != nil {
		return nil, err
	}
	insid, err := generated.NewID() // check if it has ro be the handshel id of the CEM
	if err != nil {
		return nil, err
	}
	time := time.Now()
	return &InstructionStatusUpdate{
		MessageID:     id,
		InstructionID: insid,
		MessageType:   messageType,
		StatusType:    statusType,
		Timestamp:     time,
	}, nil
}

// GetMessageID returns the message ID field info
func (is *InstructionStatusUpdate) GetMessageID() *generated.ID {
	return is.MessageID
}

// InstructionID returns the instruction ID field info
func (is *InstructionStatusUpdate) SetInstructionID() *generated.ID {
	return is.InstructionID
}
