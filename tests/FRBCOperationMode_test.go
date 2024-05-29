package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewFRBCOperationMode(t *testing.T) {
	fillLevelRange1, err := common.NewNumberRange(0.0, 50.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	fillLevelRange2, err := common.NewNumberRange(50.0, 100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	fillRate, err := common.NewNumberRange(-10.0, 10.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	powerRange, err := common.NewPowerRange(generated.ELECTRIC_POWER_L1, 0.0, 100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	powerRanges := []*common.PowerRange{powerRange}
	element1, err := frbc.NewFRBCOperationModeElement(fillLevelRange1, fillRate, powerRanges, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	element2, err := frbc.NewFRBCOperationModeElement(fillLevelRange2, fillRate, powerRanges, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	elements := []*frbc.FRBCOperationModeElement{element1, element2}
	label := "Test Operation Mode"

	// Test valid FRBCOperationMode creation
	mode, err := frbc.NewFRBCOperationMode(true, elements, &label)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if mode.GetAbnormalConditionOnly() != true {
		t.Errorf("expected abnormal condition only to be true, got %v", mode.GetAbnormalConditionOnly())
	}
	if mode.GetDiagnosticLabel() == nil || *mode.GetDiagnosticLabel() != label {
		t.Errorf("expected diagnostic label %v, got %v", label, mode.GetDiagnosticLabel())
	}
	if len(mode.GetElements()) != len(elements) {
		t.Errorf("expected %d elements, got %d", len(elements), len(mode.GetElements()))
	}
	if mode.GetID() == nil {
		t.Errorf("expected non-nil ID")
	}

	// Test creation with nil elements
	_, err = frbc.NewFRBCOperationMode(true, nil, &label)
	if err == nil {
		t.Fatalf("expected error for nil elements, got nil")
	}

	// Test creation with empty elements
	_, err = frbc.NewFRBCOperationMode(true, []*frbc.FRBCOperationModeElement{}, &label)
	if err == nil {
		t.Fatalf("expected error for empty elements, got nil")
	}

	// Test creation with non-contiguous elements
	nonContiguousRange, err := common.NewNumberRange(200.0, 300.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	element3, err := frbc.NewFRBCOperationModeElement(nonContiguousRange, fillRate, powerRanges, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	_, err = frbc.NewFRBCOperationMode(true, []*frbc.FRBCOperationModeElement{element1, element3}, &label)
	if err == nil {
		t.Fatalf("expected error for non-contiguous elements, got nil")
	}
}

func TestFRBCOperationModeMethods(t *testing.T) {
	fillLevelRange1, err := common.NewNumberRange(0.0, 50.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	fillLevelRange2, err := common.NewNumberRange(50.0, 100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	fillRate, err := common.NewNumberRange(-10.0, 10.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	powerRange, err := common.NewPowerRange(generated.ELECTRIC_POWER_L1, 0.0, 100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	powerRanges := []*common.PowerRange{powerRange}
	element1, err := frbc.NewFRBCOperationModeElement(fillLevelRange1, fillRate, powerRanges, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	element2, err := frbc.NewFRBCOperationModeElement(fillLevelRange2, fillRate, powerRanges, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	elements := []*frbc.FRBCOperationModeElement{element1, element2}
	label := "Test Operation Mode"

	mode, err := frbc.NewFRBCOperationMode(true, elements, &label)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetAbnormalConditionOnly method
	if mode.GetAbnormalConditionOnly() != true {
		t.Errorf("expected abnormal condition only to be true, got %v", mode.GetAbnormalConditionOnly())
	}

	// Test GetDiagnosticLabel method
	if mode.GetDiagnosticLabel() == nil || *mode.GetDiagnosticLabel() != label {
		t.Errorf("expected diagnostic label %v, got %v", label, mode.GetDiagnosticLabel())
	}

	// Test GetElements method
	if len(mode.GetElements()) != len(elements) {
		t.Errorf("expected %d elements, got %d", len(elements), len(mode.GetElements()))
	}

	// Test GetID method
	if mode.GetID() == nil {
		t.Errorf("expected non-nil ID")
	}
}
