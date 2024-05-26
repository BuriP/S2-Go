package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewHandshake(t *testing.T) {
	// Test valid Handshake creation for RM
	role := generated.RM
	supportedProtocols := []string{"1.0", "2.0"}

	handshake, err := common.NewHandshake(role, supportedProtocols)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if handshake.MessageType != "Handshake" {
		t.Errorf("expected MessageType 'Handshake', got %v", handshake.MessageType)
	}
	if handshake.MessageID == nil {
		t.Errorf("expected MessageID to be non-nil")
	}
	if handshake.Role != role {
		t.Errorf("expected Role %v, got %v", role, handshake.Role)
	}
	if handshake.SupportedProtocolVersions == nil || len(*handshake.SupportedProtocolVersions) != len(supportedProtocols) {
		t.Errorf("expected SupportedProtocolVersions %v, got %v", supportedProtocols, handshake.SupportedProtocolVersions)
	}

	// Test valid Handshake creation for CEM
	role = generated.CEM

	handshake, err = common.NewHandshake(role, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if handshake.SupportedProtocolVersions != nil {
		t.Errorf("expected SupportedProtocolVersions to be nil, got %v", handshake.SupportedProtocolVersions)
	}

	// Test invalid role
	_, err = common.NewHandshake("INVALID_ROLE", supportedProtocols)
	if err == nil {
		t.Fatalf("expected error for invalid role, got nil")
	}

	// Test RM with empty supportedProtocols
	_, err = common.NewHandshake(generated.RM, nil)
	if err == nil {
		t.Fatalf("expected error for RM with empty supportedProtocols, got nil")
	}
}

func TestSetMessageID(t *testing.T) {
	role := generated.RM
	supportedProtocols := []string{"1.0", "2.0"}
	handshake, err := common.NewHandshake(role, supportedProtocols)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	newPattern := "new-pattern"
	handshake.SetMessageID(newPattern)

	if handshake.MessageID == nil || handshake.MessageID.Value != generated.TransformToValidID(newPattern) {
		t.Errorf("expected MessageID to be set to transformed pattern, got %v", handshake.MessageID)
	}
}

func TestSetSupportedProtocolVersions(t *testing.T) {
	role := generated.RM
	supportedProtocols := []string{"1.0", "2.0"}
	handshake, err := common.NewHandshake(role, supportedProtocols)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	newSupportedProtocols := []string{"3.0", "4.0"}
	handshake.SetSupportedProtocolVersions(newSupportedProtocols)

	if handshake.SupportedProtocolVersions == nil || len(*handshake.SupportedProtocolVersions) != len(newSupportedProtocols) {
		t.Errorf("expected SupportedProtocolVersions to be set to %v, got %v", newSupportedProtocols, handshake.SupportedProtocolVersions)
	}
}
