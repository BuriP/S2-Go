package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
)

func TestNewFRBCFillLevelTargetProfileElement(t *testing.T) {
	duration := common.NewDuration(3600) // 1 hour
	fillLevelRange, err := common.NewNumberRange(0, 100)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test valid FRBCFillLevelTargetProfileElement creation
	element, err := frbc.NewFRBCFillLevelTargetProfileElement(duration, fillLevelRange)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if element.Duration == nil || element.Duration.Milliseconds != duration.Milliseconds {
		t.Errorf("expected duration %v, got %v", duration.Milliseconds, element.Duration.Milliseconds)
	}
	if element.FillLevelRange == nil || element.FillLevelRange.StartOfRange != fillLevelRange.StartOfRange || element.FillLevelRange.EndOfRange != fillLevelRange.EndOfRange {
		t.Errorf("expected fill level range min %v max %v, got min %v max %v", fillLevelRange.StartOfRange, fillLevelRange.EndOfRange, element.FillLevelRange.StartOfRange, element.FillLevelRange.EndOfRange)
	}

	// Test creation with nil duration
	_, err = frbc.NewFRBCFillLevelTargetProfileElement(nil, fillLevelRange)
	if err == nil {
		t.Fatalf("expected error for nil duration, got nil")
	}

	// Test creation with nil fill level range
	_, err = frbc.NewFRBCFillLevelTargetProfileElement(duration, nil)
	if err == nil {
		t.Fatalf("expected error for nil fill level range, got nil")
	}
}

func TestFRBCFillLevelTargetProfileElementMethods(t *testing.T) {
	duration := common.NewDuration(3600) // 1 hour
	fillLevelRange, err := common.NewNumberRange(0, 100)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	element, err := frbc.NewFRBCFillLevelTargetProfileElement(duration, fillLevelRange)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetDuration method
	if element.GetDuration() == nil || element.GetDuration().Milliseconds != duration.Milliseconds {
		t.Errorf("expected duration %v, got %v", duration.Milliseconds, element.GetDuration().Milliseconds)
	}

	// Test GetFillLevelRange method
	if element.GetFillLevelRange() == nil || element.GetFillLevelRange().StartOfRange != fillLevelRange.StartOfRange || element.GetFillLevelRange().EndOfRange != fillLevelRange.EndOfRange {
		t.Errorf("expected fill level range min %v max %v, got min %v max %v", fillLevelRange.StartOfRange, fillLevelRange.EndOfRange, element.GetFillLevelRange().StartOfRange, element.GetFillLevelRange().EndOfRange)
	}
}
