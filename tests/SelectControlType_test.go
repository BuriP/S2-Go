package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewSelectControlType(t *testing.T) {
	// Test valid SelectControlType creation
	controlType := generated.POWER_ENVELOPE_BASED_CONTROL

	selectControlType, err := common.NewSelectControlType(controlType)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if selectControlType.MessageType != "Select Control Type" {
		t.Errorf("expected MessageType 'Select Control Type', got %v", selectControlType.MessageType)
	}
	if selectControlType.MessageID == nil {
		t.Errorf("expected MessageID to be non-nil")
	}
	if selectControlType.ControlType != controlType {
		t.Errorf("expected ControlType %v, got %v", controlType, selectControlType.ControlType)
	}

	// Additional tests for other ControlType constants
	controlTypes := []generated.ControlType{
		generated.POWER_PROFILE_BASED_CONTROL,
		generated.OPERATION_MODE_BASED_CONTROL,
		generated.FILL_RATE_BASED_CONTROL,
		generated.DEMAND_DRIVEN_BASED_CONTROL,
		generated.NOT_CONTROLABLE,
		generated.NO_SELECTION,
	}

	for _, ct := range controlTypes {
		selectControlType, err := common.NewSelectControlType(ct)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if selectControlType.ControlType != ct {
			t.Errorf("expected ControlType %v, got %v", ct, selectControlType.ControlType)
		}
	}
}
