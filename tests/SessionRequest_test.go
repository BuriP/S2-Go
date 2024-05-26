package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewSessionRequest(t *testing.T) {
	diagnosticLabel := "Test Diagnostic Label"
	requestType := generated.RECONNECT

	// Test valid SessionRequest creation with diagnostic label
	sessionRequest, err := common.NewSessionRequest(requestType, diagnosticLabel)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if sessionRequest.MessageType != "Session Request" {
		t.Errorf("expected MessageType 'Session Request', got %v", sessionRequest.MessageType)
	}
	if sessionRequest.MessageID == nil {
		t.Errorf("expected MessageID to be non-nil")
	}
	if sessionRequest.Request != requestType {
		t.Errorf("expected Request %v, got %v", requestType, sessionRequest.Request)
	}
	if sessionRequest.DiagnosticLabel == nil || *sessionRequest.DiagnosticLabel != diagnosticLabel {
		t.Errorf("expected DiagnosticLabel %v, got %v", diagnosticLabel, *sessionRequest.DiagnosticLabel)
	}

	// Test valid SessionRequest creation without diagnostic label
	sessionRequest, err = common.NewSessionRequest(requestType)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if sessionRequest.DiagnosticLabel != nil {
		t.Errorf("expected DiagnosticLabel to be nil, got %v", *sessionRequest.DiagnosticLabel)
	}

	// Test creation with empty request
	_, err = common.NewSessionRequest("")
	if err == nil {
		t.Fatalf("expected error for empty request, got nil")
	}

	// Additional tests for other SessionRequestType constants
	requestTypes := []generated.SessionRequestType{
		generated.RECONNECT,
		generated.TERMINATE,
	}

	for _, rt := range requestTypes {
		sessionRequest, err := common.NewSessionRequest(rt, diagnosticLabel)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if sessionRequest.Request != rt {
			t.Errorf("expected Request %v, got %v", rt, sessionRequest.Request)
		}
	}
}
