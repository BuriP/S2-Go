package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewHandshakeResponse(t *testing.T) {
	supportedProtocols := []string{"1.0", "2.0"}
	handshake, err := common.NewHandshake(generated.CEM, supportedProtocols)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test valid HandshakeResponse creation
	handshakeResponse, err := common.NewHandshakeResponse(handshake)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if handshakeResponse.MessageType != "HandshakeResponse" {
		t.Errorf("expected MessageType 'HandshakeResponse', got %v", handshakeResponse.MessageType)
	}
	if handshakeResponse.MessageID == nil {
		t.Errorf("expected MessageID to be non-nil")
	}
	if handshakeResponse.SelectedProtocolVersion == nil || len(*handshakeResponse.SelectedProtocolVersion) != len(supportedProtocols) {
		t.Errorf("expected SelectedProtocolVersion %v, got %v", supportedProtocols, handshakeResponse.SelectedProtocolVersion)
	}

	// Test HandshakeResponse creation with nil handshake
	_, err = common.NewHandshakeResponse(nil)
	if err == nil {
		t.Fatalf("expected error for nil handshake, got nil")
	}

	// Test HandshakeResponse creation with nil SupportedProtocolVersions
	handshake.SupportedProtocolVersions = nil
	_, err = common.NewHandshakeResponse(handshake)
	if err == nil {
		t.Fatalf("expected error for handshake with nil SupportedProtocolVersions, got nil")
	}
}

func TestGetSelectedProtocolVersions(t *testing.T) {
	supportedProtocols := []string{"1.0", "2.0"}
	handshake, err := common.NewHandshake(generated.CEM, supportedProtocols)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	handshakeResponse, err := common.NewHandshakeResponse(handshake)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	selectedProtocols := handshakeResponse.GetSelectedProtocolVersions()
	if selectedProtocols == nil || len(*selectedProtocols) != len(supportedProtocols) {
		t.Errorf("expected SelectedProtocolVersions %v, got %v", supportedProtocols, selectedProtocols)
	}
}
