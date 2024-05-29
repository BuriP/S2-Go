package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
)

func TestNewFRBCLeakageBehaviourElement(t *testing.T) {
	fillLevelRange, err := common.NewNumberRange(0, 100)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	leakageRate := 0.5

	// Test valid FRBCLeakageBehaviourElement creation
	element, err := frbc.NewFRBCLeakageBehaviourElement(fillLevelRange, leakageRate)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if element.FillLevelRange == nil || element.FillLevelRange.StartOfRange != fillLevelRange.StartOfRange || element.FillLevelRange.EndOfRange != fillLevelRange.EndOfRange {
		t.Errorf("expected fill level range min %v max %v, got min %v max %v", fillLevelRange.StartOfRange, fillLevelRange.EndOfRange, element.FillLevelRange.StartOfRange, element.FillLevelRange.EndOfRange)
	}
	if element.LeakageRate != leakageRate {
		t.Errorf("expected leakage rate %v, got %v", leakageRate, element.LeakageRate)
	}

	// Test creation with nil fill level range
	_, err = frbc.NewFRBCLeakageBehaviourElement(nil, leakageRate)
	if err == nil {
		t.Fatalf("expected error for nil fill level range, got nil")
	}

	// Test creation with zero leakage rate
	_, err = frbc.NewFRBCLeakageBehaviourElement(fillLevelRange, 0)
	if err == nil {
		t.Fatalf("expected error for zero leakage rate, got nil")
	}
}

func TestFRBCLeakageBehaviourElementMethods(t *testing.T) {
	fillLevelRange, err := common.NewNumberRange(0, 100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	leakageRate := 0.5

	element, err := frbc.NewFRBCLeakageBehaviourElement(fillLevelRange, leakageRate)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetFillLevelRange method
	if element.GetFillLevelRange() == nil || element.GetFillLevelRange().StartOfRange != fillLevelRange.StartOfRange || element.GetFillLevelRange().EndOfRange != fillLevelRange.EndOfRange {
		t.Errorf("expected fill level range min %v max %v, got min %v max %v", fillLevelRange.StartOfRange, fillLevelRange.EndOfRange, element.GetFillLevelRange().StartOfRange, element.GetFillLevelRange().EndOfRange)
	}

	// Test GetLeakageRate method
	if element.GetLeakageRate() != leakageRate {
		t.Errorf("expected leakage rate %v, got %v", leakageRate, element.GetLeakageRate)
	}
}

