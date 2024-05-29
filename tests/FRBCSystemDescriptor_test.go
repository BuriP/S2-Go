package tests

import (
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/frbc"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewFRBCSystemDescription(t *testing.T) {
	actuatorID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	actuator := &frbc.FRBCActuatorDescriptor{
		ID: actuatorID,
	}

	storage := &frbc.FRBCStorageDescription{}
	actuators := []*frbc.FRBCActuatorDescriptor{actuator}
	validFrom := time.Now()

	// Test valid FRBCSystemDescription creation
	systemDescription, err := frbc.NewFRBCSystemDescription(actuators, storage, validFrom)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(systemDescription.GetActuators()) != len(actuators) {
		t.Errorf("expected %d actuators, got %d", len(actuators), len(systemDescription.GetActuators()))
	}
	if systemDescription.GetMessageID() == nil || systemDescription.GetMessageID().Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", systemDescription.GetMessageID())
	}
	if systemDescription.GetMessageType() != "FRBC.SystemDescription" {
		t.Errorf("expected message type FRBC.SystemDescription, got %v", systemDescription.GetMessageType())
	}
	if systemDescription.GetStorage() != storage {
		t.Errorf("expected storage %v, got %v", storage, systemDescription.GetStorage())
	}
	if !systemDescription.GetValidFrom().Equal(validFrom) {
		t.Errorf("expected valid from time %v, got %v", validFrom, systemDescription.GetValidFrom())
	}

	// Test creation with nil actuators
	_, err = frbc.NewFRBCSystemDescription(nil, storage, validFrom)
	if err == nil {
		t.Fatalf("expected error for nil actuators, got nil")
	}

	// Test creation with empty actuators
	_, err = frbc.NewFRBCSystemDescription([]*frbc.FRBCActuatorDescriptor{}, storage, validFrom)
	if err == nil {
		t.Fatalf("expected error for empty actuators, got nil")
	}

	// Test creation with more than 10 actuators
	excessiveActuators := make([]*frbc.FRBCActuatorDescriptor, 11)
	for i := range excessiveActuators {
		excessiveActuators[i] = actuator
	}
	_, err = frbc.NewFRBCSystemDescription(excessiveActuators, storage, validFrom)
	if err == nil {
		t.Fatalf("expected error for more than 10 actuators, got nil")
	}
}

func TestFRBCSystemDescriptionMethods(t *testing.T) {
	actuatorID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	actuator := &frbc.FRBCActuatorDescriptor{
		ID: actuatorID,
	}

	storage := &frbc.FRBCStorageDescription{}
	actuators := []*frbc.FRBCActuatorDescriptor{actuator}
	validFrom := time.Now()

	systemDescription, err := frbc.NewFRBCSystemDescription(actuators, storage, validFrom)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetActuators method
	if len(systemDescription.GetActuators()) != len(actuators) {
		t.Errorf("expected %d actuators, got %d", len(actuators), len(systemDescription.GetActuators()))
	}

	// Test GetMessageID method
	if systemDescription.GetMessageID() == nil || systemDescription.GetMessageID().Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", systemDescription.GetMessageID())
	}

	// Test GetMessageType method
	if systemDescription.GetMessageType() != "FRBC.SystemDescription" {
		t.Errorf("expected message type FRBC.SystemDescription, got %v", systemDescription.GetMessageType())
	}

	// Test GetStorage method
	if systemDescription.GetStorage() != storage {
		t.Errorf("expected storage %v, got %v", storage, systemDescription.GetStorage())
	}

	// Test GetValidFrom method
	if !systemDescription.GetValidFrom().Equal(validFrom) {
		t.Errorf("expected valid from time %v, got %v", validFrom, systemDescription.GetValidFrom())
	}
}
