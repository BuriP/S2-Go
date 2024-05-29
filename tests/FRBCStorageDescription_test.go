package tests

import (
	"encoding/json"
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
)

func TestNewFRBCStorageDescription(t *testing.T) {
	fillLevelRange, err := common.NewNumberRange(0, 100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	diagnosticLabel := "Test Storage"
	fillLevelLabel := "Percentage"

	// Test valid FRBCStorageDescription creation
	storageDescription, err := frbc.NewFRBCStorageDescription(fillLevelRange, true, true, true, &diagnosticLabel, &fillLevelLabel)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if storageDescription.FillLevelRange == nil || storageDescription.FillLevelRange.StartOfRange != fillLevelRange.StartOfRange || storageDescription.FillLevelRange.EndOfRange != fillLevelRange.EndOfRange {
		t.Errorf("expected FillLevelRange %v, got %v", fillLevelRange, storageDescription.FillLevelRange)
	}
	if storageDescription.DiagnosticLabel == nil || *storageDescription.DiagnosticLabel != diagnosticLabel {
		t.Errorf("expected DiagnosticLabel %v, got %v", diagnosticLabel, storageDescription.DiagnosticLabel)
	}
	if storageDescription.FillLevelLabel == nil || *storageDescription.FillLevelLabel != fillLevelLabel {
		t.Errorf("expected FillLevelLabel %v, got %v", fillLevelLabel, storageDescription.FillLevelLabel)
	}
	if !storageDescription.ProvidesFillLevelTargetProfile {
		t.Errorf("expected ProvidesFillLevelTargetProfile to be true, got %v", storageDescription.ProvidesFillLevelTargetProfile)
	}
	if !storageDescription.ProvidesLeakageBehaviour {
		t.Errorf("expected ProvidesLeakageBehaviour to be true, got %v", storageDescription.ProvidesLeakageBehaviour)
	}
	if !storageDescription.ProvidesUsageForecast {
		t.Errorf("expected ProvidesUsageForecast to be true, got %v", storageDescription.ProvidesUsageForecast)
	}

	// Test creation with nil fillLevelRange
	_, err = frbc.NewFRBCStorageDescription(nil, true, true, true, &diagnosticLabel, &fillLevelLabel)
	if err != nil {
		t.Fatalf("expected error for nil fillLevelRange, got nil")
	}
}

func TestFRBCStorageDescriptionSerialization(t *testing.T) {
	fillLevelRange, err := common.NewNumberRange(0, 100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	diagnosticLabel := "Test Storage"
	fillLevelLabel := "Percentage"

	storageDescription, err := frbc.NewFRBCStorageDescription(fillLevelRange, true, true, true, &diagnosticLabel, &fillLevelLabel)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(storageDescription)
	if err != nil {
		t.Fatalf("failed to marshal FRBCStorageDescription to JSON: %v", err)
	}

	// Deserialize from JSON
	var deserializedStorageDescription frbc.FRBCStorageDescription
	err = json.Unmarshal(jsonData, &deserializedStorageDescription)
	if err != nil {
		t.Fatalf("failed to unmarshal JSON to FRBCStorageDescription: %v", err)
	}

	// Compare original and deserialized values
	if deserializedStorageDescription.FillLevelRange == nil || deserializedStorageDescription.FillLevelRange.StartOfRange != fillLevelRange.StartOfRange || deserializedStorageDescription.FillLevelRange.EndOfRange != fillLevelRange.EndOfRange {
		t.Errorf("expected FillLevelRange %v, got %v", fillLevelRange, deserializedStorageDescription.FillLevelRange)
	}
	if deserializedStorageDescription.DiagnosticLabel == nil || *deserializedStorageDescription.DiagnosticLabel != diagnosticLabel {
		t.Errorf("expected DiagnosticLabel %v, got %v", diagnosticLabel, deserializedStorageDescription.DiagnosticLabel)
	}
	if deserializedStorageDescription.FillLevelLabel == nil || *deserializedStorageDescription.FillLevelLabel != fillLevelLabel {
		t.Errorf("expected FillLevelLabel %v, got %v", fillLevelLabel, deserializedStorageDescription.FillLevelLabel)
	}
	if !deserializedStorageDescription.ProvidesFillLevelTargetProfile {
		t.Errorf("expected ProvidesFillLevelTargetProfile to be true, got %v", deserializedStorageDescription.ProvidesFillLevelTargetProfile)
	}
	if !deserializedStorageDescription.ProvidesLeakageBehaviour {
		t.Errorf("expected ProvidesLeakageBehaviour to be true, got %v", deserializedStorageDescription.ProvidesLeakageBehaviour)
	}
	if !deserializedStorageDescription.ProvidesUsageForecast {
		t.Errorf("expected ProvidesUsageForecast to be true, got %v", deserializedStorageDescription.ProvidesUsageForecast)
	}
}
