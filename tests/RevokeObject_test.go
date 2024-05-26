package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewRevokeObject(t *testing.T) {
	// Test valid RevokeObject creation
	objectID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	revokeObject, err := common.NewRevokeObject(objectID, generated.PEBC_PowerConstraints)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if revokeObject.MessageType != "Revoke Object" {
		t.Errorf("expected MessageType 'Revoke Object', got %v", revokeObject.MessageType)
	}
	if revokeObject.MessageID == nil {
		t.Errorf("expected MessageID to be non-nil")
	}
	if revokeObject.ObjectID == nil || revokeObject.ObjectID.Value != objectID.Value {
		t.Errorf("expected ObjectID %v, got %v", objectID.Value, revokeObject.ObjectID.Value)
	}
	if revokeObject.ObjectType != generated.PEBC_PowerConstraints {
		t.Errorf("expected ObjectType %v, got %v", generated.PEBC_PowerConstraints, revokeObject.ObjectType)
	}

	// Test creation with nil objectID
	_, err = common.NewRevokeObject(nil, generated.PEBC_PowerConstraints)
	if err == nil {
		t.Fatalf("expected error for nil objectID, got nil")
	}

	// Additional tests for other RevokableObject types
	revokeObject, err = common.NewRevokeObject(objectID, generated.FRBC_Instruction)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if revokeObject.ObjectType != generated.FRBC_Instruction {
		t.Errorf("expected ObjectType %v, got %v", generated.FRBC_Instruction, revokeObject.ObjectType)
	}
}
