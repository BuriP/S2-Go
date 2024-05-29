package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewFRBCOperationModeElement(t *testing.T) {
	fillLevelRange, err := common.NewNumberRange(0, 100)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	fillRate, err := common.NewNumberRange(-10, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	powerRange, err := common.NewPowerRange(generated.ELECTRIC_POWER_L1, 0, 100)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	powerRanges := []*common.PowerRange{powerRange}
	runningCosts, err := common.NewNumberRange(1, 5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test valid FRBCOperationModeElement creation
	element, err := frbc.NewFRBCOperationModeElement(fillLevelRange, fillRate, powerRanges, runningCosts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if element.FillLevelRange == nil || element.FillLevelRange.StartOfRange != fillLevelRange.StartOfRange || element.FillLevelRange.EndOfRange != fillLevelRange.EndOfRange {
		t.Errorf("expected fill level range start %v end %v, got start %v end %v", fillLevelRange.StartOfRange, fillLevelRange.EndOfRange, element.FillLevelRange.StartOfRange, element.FillLevelRange.EndOfRange)
	}
	if element.FillRate == nil || element.FillRate.StartOfRange != fillRate.StartOfRange || element.FillRate.EndOfRange != fillRate.EndOfRange {
		t.Errorf("expected fill rate start %v end %v, got start %v end %v", fillRate.StartOfRange, fillRate.EndOfRange, element.FillRate.StartOfRange, element.FillRate.EndOfRange)
	}
	if len(element.PowerRanges) != len(powerRanges) || element.PowerRanges[0].StartOfRange != powerRange.StartOfRange || element.PowerRanges[0].EndOfRange != powerRange.EndOfRange {
		t.Errorf("expected power range start %v end %v, got start %v end %v", powerRange.StartOfRange, powerRange.EndOfRange, element.PowerRanges[0].StartOfRange, element.PowerRanges[0].EndOfRange)
	}
	if element.RunningCosts == nil || element.RunningCosts.StartOfRange != runningCosts.StartOfRange || element.RunningCosts.EndOfRange != runningCosts.EndOfRange {
		t.Errorf("expected running costs start %v end %v, got start %v end %v", runningCosts.StartOfRange, runningCosts.EndOfRange, element.RunningCosts.StartOfRange, element.RunningCosts.EndOfRange)
	}

	// Test creation with nil fill level range
	_, err = frbc.NewFRBCOperationModeElement(nil, fillRate, powerRanges, runningCosts)
	if err == nil {
		t.Fatalf("expected error for nil fill level range, got nil")
	}

	// Test creation with nil fill rate
	_, err = frbc.NewFRBCOperationModeElement(fillLevelRange, nil, powerRanges, runningCosts)
	if err == nil {
		t.Fatalf("expected error for nil fill rate, got nil")
	}

	// Test creation with no power ranges
	_, err = frbc.NewFRBCOperationModeElement(fillLevelRange, fillRate, []*common.PowerRange{}, runningCosts)
	if err == nil {
		t.Fatalf("expected error for no power ranges, got nil")
	}
}

func TestFRBCOperationModeElementMethods(t *testing.T) {
	fillLevelRange, err := common.NewNumberRange(0, 100)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	fillRate, err := common.NewNumberRange(-10, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	powerRange, err := common.NewPowerRange(generated.ELECTRIC_POWER_L2, 0, 100)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	powerRanges := []*common.PowerRange{powerRange}
	runningCosts, err := common.NewNumberRange(1, 5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	element, err := frbc.NewFRBCOperationModeElement(fillLevelRange, fillRate, powerRanges, runningCosts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetFillLevelRange method
	if element.GetFillLevelRange() == nil || element.GetFillLevelRange().StartOfRange != fillLevelRange.StartOfRange || element.GetFillLevelRange().EndOfRange != fillLevelRange.EndOfRange {
		t.Errorf("expected fill level range start %v end %v, got start %v end %v", fillLevelRange.StartOfRange, fillLevelRange.EndOfRange, element.GetFillLevelRange().StartOfRange, element.GetFillLevelRange().EndOfRange)
	}

	// Test GetFillRate method
	if element.GetFillRate() == nil || element.GetFillRate().StartOfRange != fillRate.StartOfRange || element.GetFillRate().EndOfRange != fillRate.EndOfRange {
		t.Errorf("expected fill rate start %v end %v, got start %v end %v", fillRate.StartOfRange, fillRate.EndOfRange, element.GetFillRate().StartOfRange, element.GetFillRate().EndOfRange)
	}

	// Test GetPowerRanges method
	if len(element.GetPowerRanges()) != len(powerRanges) || element.GetPowerRanges()[0].StartOfRange != powerRange.StartOfRange || element.GetPowerRanges()[0].EndOfRange != powerRange.EndOfRange {
		t.Errorf("expected power range start %v end %v, got start %v end %v", powerRange.StartOfRange, powerRange.EndOfRange, element.GetPowerRanges()[0].StartOfRange, element.GetPowerRanges()[0].EndOfRange)
	}

	// Test GetRunningCosts method
	if element.GetRunningCosts() == nil || element.GetRunningCosts().StartOfRange != runningCosts.StartOfRange || element.GetRunningCosts().EndOfRange != runningCosts.EndOfRange {
		t.Errorf("expected running costs start %v end %v, got start %v end %v", runningCosts.StartOfRange, runningCosts.EndOfRange, element.GetRunningCosts().StartOfRange, element.GetRunningCosts().EndOfRange)
	}
}
