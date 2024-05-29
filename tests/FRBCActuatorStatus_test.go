package tests

import (
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/frbc"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewFRBCActuatorStatus(t *testing.T) {
	activeModeID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	actuatorID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	previousModeID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	transitionTime := time.Now()

	// Test valid FRBCActuatorStatus creation
	status, err := frbc.NewFRBCActuatorStatus(activeModeID, actuatorID, 0.75, previousModeID, &transitionTime)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if status.GetActiveOperationModeID() == nil || status.GetActiveOperationModeID().Value != activeModeID.Value {
		t.Errorf("expected active operation mode ID %v, got %v", activeModeID.Value, status.GetActiveOperationModeID().Value)
	}
	if status.GetActuatorID() == nil || status.GetActuatorID().Value != actuatorID.Value {
		t.Errorf("expected actuator ID %v, got %v", actuatorID.Value, status.GetActuatorID().Value)
	}
	if status.GetMessageID() == nil || status.GetMessageID().Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", status.GetMessageID())
	}
	if status.GetMessageType() != "FRBCActuatorStatus" {
		t.Errorf("expected message type FRBCActuatorStatus, got %v", status.GetMessageType())
	}
	if status.GetOperationModeFactor() != 0.75 {
		t.Errorf("expected operation mode factor 0.75, got %v", status.GetOperationModeFactor())
	}
	if status.GetPreviousOperationModeID() == nil || status.GetPreviousOperationModeID().Value != previousModeID.Value {
		t.Errorf("expected previous operation mode ID %v, got %v", previousModeID.Value, status.GetPreviousOperationModeID().Value)
	}
	if status.GetTransitionTimestamp() == nil || !status.GetTransitionTimestamp().Equal(transitionTime) {
		t.Errorf("expected transition timestamp %v, got %v", transitionTime, status.GetTransitionTimestamp())
	}

	// Test creation without previous operation mode ID and transition timestamp
	status, err = frbc.NewFRBCActuatorStatus(activeModeID, actuatorID, 0.75, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if status.GetPreviousOperationModeID() != nil {
		t.Errorf("expected previous operation mode ID to be nil, got %v", status.GetPreviousOperationModeID())
	}
	if status.GetTransitionTimestamp() != nil {
		t.Errorf("expected transition timestamp to be nil, got %v", status.GetTransitionTimestamp())
	}
}
