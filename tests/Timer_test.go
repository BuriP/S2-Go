package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
)

func TestNewTimer(t *testing.T) {
	duration := common.NewDuration(3600) // 1 hour

	// Test valid Timer creation without diagnostic label
	timer, err := common.NewTimer(duration)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if timer.Duration == nil || timer.Duration.Milliseconds != duration.Milliseconds {
		t.Errorf("expected duration %v, got %v", duration, timer.Duration)
	}
	if timer.ID == nil {
		t.Errorf("expected non-nil ID")
	}
	if timer.DiagnosticLabel != nil {
		t.Errorf("expected nil DiagnosticLabel, got %v", *timer.DiagnosticLabel)
	}

	// Test valid Timer creation with diagnostic label
	label := "Test Timer"
	timerWithLabel, err := common.NewTimer(duration, label)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if timerWithLabel.DiagnosticLabel == nil || *timerWithLabel.DiagnosticLabel != label {
		t.Errorf("expected DiagnosticLabel %v, got %v", label, *timerWithLabel.DiagnosticLabel)
	}

	// Test creation with nil duration
	_, err = common.NewTimer(nil)
	if err == nil {
		t.Fatalf("expected error for nil duration, got nil")
	}
}
