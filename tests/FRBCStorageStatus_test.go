package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/frbc"
)

func TestNewFRBCStorageStatus(t *testing.T) {
	// Test valid FRBCStorageStatus creation
	presentFillLevel := 75.0
	status, err := frbc.NewFRBCStorageStatus(presentFillLevel)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if status.GetMessageID() == nil || status.GetMessageID().Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", status.GetMessageID())
	}
	if status.GetMessageType() != "FRBC.StorageStatus" {
		t.Errorf("expected message type FRBC.StorageStatus, got %v", status.GetMessageType())
	}
	if status.GetPresentFillLevel() != presentFillLevel {
		t.Errorf("expected present fill level %v, got %v", presentFillLevel, status.GetPresentFillLevel())
	}
}

func TestFRBCStorageStatusMethods(t *testing.T) {
	presentFillLevel := 75.0
	status, err := frbc.NewFRBCStorageStatus(presentFillLevel)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetMessageID method
	if status.GetMessageID() == nil || status.GetMessageID().Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", status.GetMessageID())
	}

	// Test GetMessageType method
	if status.GetMessageType() != "FRBC.StorageStatus" {
		t.Errorf("expected message type FRBC.StorageStatus, got %v", status.GetMessageType())
	}

	// Test GetPresentFillLevel method
	if status.GetPresentFillLevel() != presentFillLevel {
		t.Errorf("expected present fill level %v, got %v", presentFillLevel, status.GetPresentFillLevel())
	}
}
