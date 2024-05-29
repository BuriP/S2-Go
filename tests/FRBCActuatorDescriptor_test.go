package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewFRBCActuatorDescriptor(t *testing.T) {
	id, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	fillLevelRange, err := common.NewNumberRange(0.0, 100.0)
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
	element, err := frbc.NewFRBCOperationModeElement(fillLevelRange, fillRate, powerRanges, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	operationModes := []*frbc.FRBCOperationMode{
		{
			AbnormalConditionOnly: true,
			Elements:              []*frbc.FRBCOperationModeElement{element},
			ID:                    id,
		},
	}
	supportedCommodities := []generated.Commodity{generated.ELECTRICITY}

	// Test valid FRBCActuatorDescriptor creation
	descriptor, err := frbc.NewFRBCActuatorDescriptor(id, operationModes, supportedCommodities, nil, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if descriptor.GetID() == nil || descriptor.GetID().Value != id.Value {
		t.Errorf("expected ID %v, got %v", id.Value, descriptor.GetID().Value)
	}
	if len(descriptor.GetOperationModes()) != len(operationModes) {
		t.Errorf("expected %d operation modes, got %d", len(operationModes), len(descriptor.GetOperationModes()))
	}
	if len(descriptor.GetSupportedCommodities()) != len(supportedCommodities) {
		t.Errorf("expected %d supported commodities, got %d", len(supportedCommodities), len(descriptor.GetSupportedCommodities()))
	}

	// Test creation with nil ID
	_, err = frbc.NewFRBCActuatorDescriptor(nil, operationModes, supportedCommodities, nil, nil, nil)
	if err == nil {
		t.Fatalf("expected error for nil ID, got nil")
	}

	// Test creation with empty operation modes
	_, err = frbc.NewFRBCActuatorDescriptor(id, []*frbc.FRBCOperationMode{}, supportedCommodities, nil, nil, nil)
	if err == nil {
		t.Fatalf("expected error for empty operation modes, got nil")
	}

	// Test creation with empty supported commodities
	_, err = frbc.NewFRBCActuatorDescriptor(id, operationModes, []generated.Commodity{}, nil, nil, nil)
	if err == nil {
		t.Fatalf("expected error for empty supported commodities, got nil")
	}
}

func TestFRBCActuatorDescriptorMethods(t *testing.T) {
	id, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	fillLevelRange, err := common.NewNumberRange(0.0, 100.0)
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
	element, err := frbc.NewFRBCOperationModeElement(fillLevelRange, fillRate, powerRanges, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	operationModes := []*frbc.FRBCOperationMode{
		{
			AbnormalConditionOnly: true,
			Elements:              []*frbc.FRBCOperationModeElement{element},
			ID:                    id,
		},
	}
	supportedCommodities := []generated.Commodity{generated.ELECTRICITY}
	label := "Test Actuator"

	descriptor, err := frbc.NewFRBCActuatorDescriptor(id, operationModes, supportedCommodities, &label, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetDiagnosticLabel method
	if descriptor.GetDiagnosticLabel() == nil || *descriptor.GetDiagnosticLabel() != label {
		t.Errorf("expected diagnostic label %v, got %v", label, descriptor.GetDiagnosticLabel())
	}

	// Test GetID method
	if descriptor.GetID() == nil || descriptor.GetID().Value != id.Value {
		t.Errorf("expected ID %v, got %v", id.Value, descriptor.GetID().Value)
	}

	// Test GetOperationModes method
	if len(descriptor.GetOperationModes()) != len(operationModes) {
		t.Errorf("expected %d operation modes, got %d", len(operationModes), len(descriptor.GetOperationModes()))
	}

	// Test GetSupportedCommodities method
	if len(descriptor.GetSupportedCommodities()) != len(supportedCommodities) {
		t.Errorf("expected %d supported commodities, got %d", len(supportedCommodities), len(descriptor.GetSupportedCommodities()))
	}

	// Test GetTimers method
	if len(descriptor.GetTimers()) != 0 {
		t.Errorf("expected 0 timers, got %d", len(descriptor.GetTimers()))
	}

	// Test GetTransitions method
	if len(descriptor.GetTransitions()) != 0 {
		t.Errorf("expected 0 transitions, got %d", len(descriptor.GetTransitions()))
	}
}
