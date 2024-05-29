package tests

import (
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
)

func TestNewFRBCFillLevelTargetProfile(t *testing.T) {
	duration1 := common.NewDuration(3600) // 1 hour
	fillLevelRange1, err := common.NewNumberRange(20.0, 40.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	element1, err := frbc.NewFRBCFillLevelTargetProfileElement(duration1, fillLevelRange1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	duration2 := common.NewDuration(7200) // 2 hours
	fillLevelRange2, err := common.NewNumberRange(40.0, 60.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	element2, err := frbc.NewFRBCFillLevelTargetProfileElement(duration2, fillLevelRange2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	elements := []*frbc.FRBCFillLevelTargetProfileElement{element1, element2}
	startTime := time.Now()

	// Test valid FRBCFillLevelTargetProfile creation
	profile, err := frbc.NewFRBCFillLevelTargetProfile(elements, startTime)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(profile.GetElements()) != len(elements) {
		t.Errorf("expected %d elements, got %d", len(elements), len(profile.GetElements()))
	}
	if profile.GetMessageID() == nil || profile.GetMessageID().Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", profile.GetMessageID())
	}
	if profile.GetMessageType() != "FRBCFillLevelTargetProfile" {
		t.Errorf("expected message type FRBCFillLevelTargetProfile, got %v", profile.GetMessageType())
	}
	if !profile.GetStartTime().Equal(startTime) {
		t.Errorf("expected start time %v, got %v", startTime, profile.GetStartTime())
	}

	// Test creation with nil elements
	_, err = frbc.NewFRBCFillLevelTargetProfile(nil, startTime)
	if err == nil {
		t.Fatalf("expected error for nil elements, got nil")
	}

	// Test creation with empty elements
	_, err = frbc.NewFRBCFillLevelTargetProfile([]*frbc.FRBCFillLevelTargetProfileElement{}, startTime)
	if err == nil {
		t.Fatalf("expected error for empty elements, got nil")
	}

	// Test creation with elements not in chronological order
	durationOutOfOrder := common.NewDuration(1800) // 30 minutes
	elementOutOfOrder, err := frbc.NewFRBCFillLevelTargetProfileElement(durationOutOfOrder, fillLevelRange1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	elementsOutOfOrder := []*frbc.FRBCFillLevelTargetProfileElement{element1, elementOutOfOrder}
	_, err = frbc.NewFRBCFillLevelTargetProfile(elementsOutOfOrder, startTime)
	if err == nil {
		t.Fatalf("expected error for elements not in chronological order, got nil")
	}
}

func TestFRBCFillLevelTargetProfileMethods(t *testing.T) {
	duration1 := common.NewDuration(3600) // 1 hour
	fillLevelRange1, err := common.NewNumberRange(20.0, 40.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	element1, err := frbc.NewFRBCFillLevelTargetProfileElement(duration1, fillLevelRange1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	duration2 := common.NewDuration(7200) // 2 hours
	fillLevelRange2, err := common.NewNumberRange(40.0, 60.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	element2, err := frbc.NewFRBCFillLevelTargetProfileElement(duration2, fillLevelRange2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	elements := []*frbc.FRBCFillLevelTargetProfileElement{element1, element2}
	startTime := time.Now()

	profile, err := frbc.NewFRBCFillLevelTargetProfile(elements, startTime)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetElements method
	if len(profile.GetElements()) != len(elements) {
		t.Errorf("expected %d elements, got %d", len(elements), len(profile.GetElements()))
	}

	// Test GetMessageID method
	if profile.GetMessageID() == nil || profile.GetMessageID().Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", profile.GetMessageID())
	}

	// Test GetMessageType method
	if profile.GetMessageType() != "FRBCFillLevelTargetProfile" {
		t.Errorf("expected message type FRBCFillLevelTargetProfile, got %v", profile.GetMessageType())
	}

	// Test GetStartTime method
	if !profile.GetStartTime().Equal(startTime) {
		t.Errorf("expected start time %v, got %v", startTime, profile.GetStartTime())
	}
}
