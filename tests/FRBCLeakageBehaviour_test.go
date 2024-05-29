package tests

import (
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
)

func TestNewFRBCLeakageBehaviour(t *testing.T) {
	fillLevelRange, err := common.NewNumberRange(0.0, 100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	leakageRate := 0.1
	element, err := frbc.NewFRBCLeakageBehaviourElement(fillLevelRange, leakageRate)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	elements := []*frbc.FRBCLeakageBehaviourElement{element}
	validFrom := time.Now()

	// Test valid FRBCLeakageBehaviour creation
	behaviour, err := frbc.NewFRBCLeakageBehaviour(elements, validFrom)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(behaviour.GetElements()) != len(elements) {
		t.Errorf("expected %d elements, got %d", len(elements), len(behaviour.GetElements()))
	}
	if behaviour.GetMessageID() == nil || behaviour.GetMessageID().Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", behaviour.GetMessageID())
	}
	if behaviour.GetMessageType() != "FRBC.LeakageBehaviour" {
		t.Errorf("expected message type FRBC.LeakageBehaviour, got %v", behaviour.GetMessageType())
	}
	if !behaviour.GetValidFrom().Equal(validFrom) {
		t.Errorf("expected valid from time %v, got %v", validFrom, behaviour.GetValidFrom())
	}

	// Test creation with nil elements
	_, err = frbc.NewFRBCLeakageBehaviour(nil, validFrom)
	if err == nil {
		t.Fatalf("expected error for nil elements, got nil")
	}

	// Test creation with empty elements
	_, err = frbc.NewFRBCLeakageBehaviour([]*frbc.FRBCLeakageBehaviourElement{}, validFrom)
	if err == nil {
		t.Fatalf("expected error for empty elements, got nil")
	}
}

func TestFRBCLeakageBehaviourMethods(t *testing.T) {
	fillLevelRange, err := common.NewNumberRange(0.0, 100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	leakageRate := 0.1
	element, err := frbc.NewFRBCLeakageBehaviourElement(fillLevelRange, leakageRate)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	elements := []*frbc.FRBCLeakageBehaviourElement{element}
	validFrom := time.Now()

	behaviour, err := frbc.NewFRBCLeakageBehaviour(elements, validFrom)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetElements method
	if len(behaviour.GetElements()) != len(elements) {
		t.Errorf("expected %d elements, got %d", len(elements), len(behaviour.GetElements()))
	}

	// Test GetMessageID method
	if behaviour.GetMessageID() == nil || behaviour.GetMessageID().Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", behaviour.GetMessageID())
	}

	// Test GetMessageType method
	if behaviour.GetMessageType() != "FRBC.LeakageBehaviour" {
		t.Errorf("expected message type FRBC.LeakageBehaviour, got %v", behaviour.GetMessageType())
	}

	// Test GetValidFrom method
	if !behaviour.GetValidFrom().Equal(validFrom) {
		t.Errorf("expected valid from time %v, got %v", validFrom, behaviour.GetValidFrom())
	}
}
