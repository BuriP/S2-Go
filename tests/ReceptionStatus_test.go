package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewReceptionStatus(t *testing.T) {
	statusValue := generated.OK
	subjectMessageID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	diagnosticLabel := "TestLabel"

	// Test valid ReceptionStatus creation
	rs, err := common.NewReceptionStatus(statusValue, subjectMessageID, &diagnosticLabel)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if rs.MessageType != "ReceptionStatus" {
		t.Errorf("expected MessageType ReceptionStatus, got %v", rs.MessageType)
	}
	if rs.Status == "" || rs.Status != statusValue {
		t.Errorf("expected Status %v, got %v", statusValue, rs.Status)
	}
	if rs.SubjectMessageID == nil || rs.SubjectMessageID.Value != subjectMessageID.Value {
		t.Errorf("expected SubjectMessageID %v, got %v", subjectMessageID.Value, rs.SubjectMessageID.Value)
	}
	if rs.DiagnosticLabel == nil || *rs.DiagnosticLabel != diagnosticLabel {
		t.Errorf("expected DiagnosticLabel %v, got %v", diagnosticLabel, rs.DiagnosticLabel)
	}

	// Test creation with nil status
	_, err = common.NewReceptionStatus("", subjectMessageID, &diagnosticLabel)
	if err == nil {
		t.Fatalf("expected error for nil status, got nil")
	}

	// Test creation with nil subjectMessageID
	_, err = common.NewReceptionStatus(statusValue, nil, &diagnosticLabel)
	if err == nil {
		t.Fatalf("expected error for nil subjectMessageID, got nil")
	}

	// Test creation without DiagnosticLabel
	rs, err = common.NewReceptionStatus(statusValue, subjectMessageID, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if rs.DiagnosticLabel != nil {
		t.Errorf("expected DiagnosticLabel to be nil, got %v", rs.DiagnosticLabel)
	}
}
