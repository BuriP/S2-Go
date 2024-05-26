package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewResourceManagerDetails(t *testing.T) {
	controlTypes := []generated.ControlType{generated.NOT_CONTROLABLE, generated.NO_SELECTION}
	measurementTypes := []generated.CommodityQuantity{generated.ELECTRIC_POWER_L1}
	// Create roles using NewRole function
	role1, err := common.NewRole(generated.GAS, generated.ENERGY_CONSUMER)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	role2, err := common.NewRole(generated.ELECTRICITY, generated.ENERGY_PRODUCER)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	roles := []*common.Role{role1, role2}

	instructionProcessingDelay := common.NewDuration(5000)
	resourceID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test valid ResourceManagerDetails creation
	rmd, err := common.NewResourceManagerDetails(
		controlTypes,
		generated.USD,
		"1.0.0",
		"Manufacturer",
		"Model1",
		"ResourceManager1",
		instructionProcessingDelay,
		true,
		measurementTypes,
		resourceID,
		roles,
		"123456789",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if rmd.MessageType != "ResourceManagerDetails" {
		t.Errorf("expected MessageType ResourceManagerDetails, got %v", rmd.MessageType)
	}
	if rmd.ResourceID == nil || rmd.ResourceID.Value != resourceID.Value {
		t.Errorf("expected ResourceID %v, got %v", resourceID.Value, rmd.ResourceID.Value)
	}
	if rmd.InstructionProcessingDelay == nil || rmd.InstructionProcessingDelay.Milliseconds != instructionProcessingDelay.Milliseconds {
		t.Errorf("expected InstructionProcessingDelay %v, got %v", instructionProcessingDelay, rmd.InstructionProcessingDelay)
	}

	// Test creation with empty controlTypes
	_, err = common.NewResourceManagerDetails(
		[]generated.ControlType{},
		generated.USD,
		"1.0.0",
		"Manufacturer",
		"Model1",
		"ResourceManager1",
		instructionProcessingDelay,
		true,
		measurementTypes,
		resourceID,
		roles,
		"123456789",
	)
	if err == nil {
		t.Fatalf("expected error for empty controlTypes, got nil")
	}

	// Test creation with nil instructionProcessingDelay
	_, err = common.NewResourceManagerDetails(
		controlTypes,
		generated.USD,
		"1.0.0",
		"Manufacturer",
		"Model1",
		"ResourceManager1",
		nil,
		true,
		measurementTypes,
		resourceID,
		roles,
		"123456789",
	)
	if err == nil {
		t.Fatalf("expected error for nil instructionProcessingDelay, got nil")
	}

	// Test creation with nil resourceID
	_, err = common.NewResourceManagerDetails(
		controlTypes,
		generated.USD,
		"1.0.0",
		"Manufacturer",
		"Model1",
		"ResourceManager1",
		instructionProcessingDelay,
		true,
		measurementTypes,
		nil,
		roles,
		"123456789",
	)
	if err == nil {
		t.Fatalf("expected error for nil resourceID, got nil")
	}
}
