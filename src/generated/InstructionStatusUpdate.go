package generated

import (
	"time"
)

// InstructionStatusUpdate represents an update to the status of an instruction.
type InstructionStatusUpdate struct {
	InstructionID ID        `json:"instruction_id" description:"ID of this instruction (as provided by the CEM)"`
	MessageID     ID        `json:"message_id" description:"ID of this message"`
	MessageType   string    `json:"message_type" description:"Type of the message" const:"true"`
	StatusType    string    `json:"status_type" description:"Present status of this instruction."`
	Timestamp     time.Time `json:"timestamp" description:"Timestamp when status_type has changed the last time."`
}
