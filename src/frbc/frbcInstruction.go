package frbc

import (
	"errors"
	"time"

	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCInstruction represents an instruction in FRBC.
type FRBCInstruction struct {
	AbnormalCondition   bool          `json:"abnormal_condition" description:"Indicates if this is an instruction during an abnormal condition"`
	ActuatorID          *generated.ID `json:"actuator_id" description:"ID of the actuator this instruction belongs to"`
	ExecutionTime       time.Time     `json:"execution_time" description:"Time when instruction should be executed"`
	ID                  *generated.ID `json:"id" description:"ID of the instruction. Must be unique in the scope of the Resource Manager, for at least the duration of the session between Resource Manager and CEM"`
	MessageID           *generated.ID `json:"message_id" description:"ID of this message"`
	MessageType         string        `json:"message_type" description:"FRBC.Instruction"`
	OperationMode       *generated.ID `json:"operation_mode" description:"ID of the FRBC.OperationMode that should be activated"`
	OperationModeFactor float64       `json:"operation_mode_factor" description:"The number indicates the factor with which the FRBC.OperationMode should be configured. The factor should be greater than or equal to 0 and less than or equal to 1"`
}

// NewFRBCInstruction creates a new instance of FRBCInstruction if validated, otherwise returns an error.
func NewFRBCInstruction(abnormalCondition bool, actuatorID, operationMode *generated.ID, executionTime time.Time, operationModeFactor float64) (*FRBCInstruction, error) {
	if actuatorID == nil {
		return nil, errors.New("actuatorID must be provided")
	}
	if operationMode == nil {
		return nil, errors.New("operationMode must be provided")
	}
	if operationModeFactor < 0 || operationModeFactor > 1 {
		return nil, errors.New("operationModeFactor must be between 0 and 1")
	}

	messageID, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	instructionID, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &FRBCInstruction{
		AbnormalCondition:   abnormalCondition,
		ActuatorID:          actuatorID,
		ExecutionTime:       executionTime,
		ID:                  instructionID,
		MessageID:           messageID,
		MessageType:         "FRBC.Instruction",
		OperationMode:       operationMode,
		OperationModeFactor: operationModeFactor,
	}, nil
}

// GetAbnormalCondition returns if the instruction is for an abnormal condition.
func (i *FRBCInstruction) GetAbnormalCondition() bool {
	return i.AbnormalCondition
}

// GetActuatorID returns the actuator ID.
func (i *FRBCInstruction) GetActuatorID() *generated.ID {
	return i.ActuatorID
}

// GetExecutionTime returns the execution time.
func (i *FRBCInstruction) GetExecutionTime() time.Time {
	return i.ExecutionTime
}

// GetID returns the instruction ID.
func (i *FRBCInstruction) GetID() *generated.ID {
	return i.ID
}

// GetMessageID returns the message ID.
func (i *FRBCInstruction) GetMessageID() *generated.ID {
	return i.MessageID
}

// GetMessageType returns the message type.
func (i *FRBCInstruction) GetMessageType() string {
	return i.MessageType
}

// GetOperationMode returns the operation mode ID.
func (i *FRBCInstruction) GetOperationMode() *generated.ID {
	return i.OperationMode
}

// GetOperationModeFactor returns the operation mode factor.
func (i *FRBCInstruction) GetOperationModeFactor() float64 {
	return i.OperationModeFactor
}
