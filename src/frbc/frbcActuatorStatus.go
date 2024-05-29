package frbc

import (
	"time"

	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCActuatorStatus represents a FRBC actuator status model.
type FRBCActuatorStatus struct {
	ActiveOperationModeID   *generated.ID `json:"active_operation_mode_id" description:"ID of the FRBC.OperationMode that is presently active."`
	ActuatorID              *generated.ID `json:"actuator_id" description:"ID of the actuator this message refers to"`
	MessageID               *generated.ID `json:"message_id" description:"ID of this message"`
	MessageType             string        `json:"message_type" description:"Type of the message"`
	OperationModeFactor     float64       `json:"operation_mode_factor" description:"The number indicates the factor with which the FRBC.OperationMode is configured"`
	PreviousOperationModeID *generated.ID `json:"previous_operation_mode_id,omitempty" description:"ID of the FRBC.OperationMode that was active before the present one"`
	TransitionTimestamp     *time.Time    `json:"transition_timestamp,omitempty" description:"Time at which the transition from the previous FRBC.OperationMode to the active FRBC.OperationMode was initiated"`
}

// NewFRBCActuatorStatus creates a new instance of FRBCActuatorStatus.
func NewFRBCActuatorStatus(activeOperationModeID, actuatorID *generated.ID, operationModeFactor float64, previousOperationModeID *generated.ID, transitionTimestamp *time.Time) (*FRBCActuatorStatus, error) {
	messageID, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &FRBCActuatorStatus{
		ActiveOperationModeID:   activeOperationModeID,
		ActuatorID:              actuatorID,
		MessageID:               messageID,
		MessageType:             "FRBCActuatorStatus",
		OperationModeFactor:     operationModeFactor,
		PreviousOperationModeID: previousOperationModeID,
		TransitionTimestamp:     transitionTimestamp,
	}, nil
}

// GetActiveOperationModeID returns the active operation mode ID.
func (s *FRBCActuatorStatus) GetActiveOperationModeID() *generated.ID {
	return s.ActiveOperationModeID
}

// GetActuatorID returns the actuator ID.
func (s *FRBCActuatorStatus) GetActuatorID() *generated.ID {
	return s.ActuatorID
}

// GetMessageID returns the message ID.
func (s *FRBCActuatorStatus) GetMessageID() *generated.ID {
	return s.MessageID
}

// GetMessageType returns the message type.
func (s *FRBCActuatorStatus) GetMessageType() string {
	return s.MessageType
}

// GetOperationModeFactor returns the operation mode factor.
func (s *FRBCActuatorStatus) GetOperationModeFactor() float64 {
	return s.OperationModeFactor
}

// GetPreviousOperationModeID returns the previous operation mode ID.
func (s *FRBCActuatorStatus) GetPreviousOperationModeID() *generated.ID {
	return s.PreviousOperationModeID
}

// GetTransitionTimestamp returns the transition timestamp.
func (s *FRBCActuatorStatus) GetTransitionTimestamp() *time.Time {
	return s.TransitionTimestamp
}
