package tests

import (
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/frbc"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewFRBCInstruction(t *testing.T) {
	actuatorID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	operationModeID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	executionTime := time.Now()

	// Test valid FRBCInstruction creation
	instruction, err := frbc.NewFRBCInstruction(true, actuatorID, operationModeID, executionTime, 0.5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if instruction.GetAbnormalCondition() != true {
		t.Errorf("expected abnormal condition to be true, got %v", instruction.GetAbnormalCondition())
	}
	if instruction.GetActuatorID() == nil || instruction.GetActuatorID().Value != actuatorID.Value {
		t.Errorf("expected actuator ID %v, got %v", actuatorID.Value, instruction.GetActuatorID().Value)
	}
	if !instruction.GetExecutionTime().Equal(executionTime) {
		t.Errorf("expected execution time %v, got %v", executionTime, instruction.GetExecutionTime())
	}
	if instruction.GetID() == nil || instruction.GetID().Value == "" {
		t.Errorf("expected a non-empty instruction ID, got %v", instruction.GetID())
	}
	if instruction.GetMessageID() == nil || instruction.GetMessageID().Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", instruction.GetMessageID())
	}
	if instruction.GetMessageType() != "FRBC.Instruction" {
		t.Errorf("expected message type FRBC.Instruction, got %v", instruction.GetMessageType())
	}
	if instruction.GetOperationMode() == nil || instruction.GetOperationMode().Value != operationModeID.Value {
		t.Errorf("expected operation mode ID %v, got %v", operationModeID.Value, instruction.GetOperationMode().Value)
	}
	if instruction.GetOperationModeFactor() != 0.5 {
		t.Errorf("expected operation mode factor 0.5, got %v", instruction.GetOperationModeFactor())
	}

	// Test creation with nil actuator ID
	_, err = frbc.NewFRBCInstruction(true, nil, operationModeID, executionTime, 0.5)
	if err == nil {
		t.Fatalf("expected error for nil actuator ID, got nil")
	}

	// Test creation with nil operation mode ID
	_, err = frbc.NewFRBCInstruction(true, actuatorID, nil, executionTime, 0.5)
	if err == nil {
		t.Fatalf("expected error for nil operation mode ID, got nil")
	}

	// Test creation with invalid operation mode factor
	_, err = frbc.NewFRBCInstruction(true, actuatorID, operationModeID, executionTime, -0.1)
	if err == nil {
		t.Fatalf("expected error for invalid operation mode factor, got nil")
	}
	_, err = frbc.NewFRBCInstruction(true, actuatorID, operationModeID, executionTime, 1.1)
	if err == nil {
		t.Fatalf("expected error for invalid operation mode factor, got nil")
	}
}

func TestFRBCInstructionMethods(t *testing.T) {
	actuatorID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	operationModeID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	executionTime := time.Now()

	instruction, err := frbc.NewFRBCInstruction(true, actuatorID, operationModeID, executionTime, 0.5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetAbnormalCondition method
	if instruction.GetAbnormalCondition() != true {
		t.Errorf("expected abnormal condition to be true, got %v", instruction.GetAbnormalCondition())
	}

	// Test GetActuatorID method
	if instruction.GetActuatorID() == nil || instruction.GetActuatorID().Value != actuatorID.Value {
		t.Errorf("expected actuator ID %v, got %v", actuatorID.Value, instruction.GetActuatorID().Value)
	}

	// Test GetExecutionTime method
	if !instruction.GetExecutionTime().Equal(executionTime) {
		t.Errorf("expected execution time %v, got %v", executionTime, instruction.GetExecutionTime())
	}

	// Test GetID method
	if instruction.GetID() == nil || instruction.GetID().Value == "" {
		t.Errorf("expected a non-empty instruction ID, got %v", instruction.GetID())
	}

	// Test GetMessageID method
	if instruction.GetMessageID() == nil || instruction.GetMessageID().Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", instruction.GetMessageID())
	}

	// Test GetMessageType method
	if instruction.GetMessageType() != "FRBC.Instruction" {
		t.Errorf("expected message type FRBC.Instruction, got %v", instruction.GetMessageType())
	}

	// Test GetOperationMode method
	if instruction.GetOperationMode() == nil || instruction.GetOperationMode().Value != operationModeID.Value {
		t.Errorf("expected operation mode ID %v, got %v", operationModeID.Value, instruction.GetOperationMode().Value)
	}

	// Test GetOperationModeFactor method
	if instruction.GetOperationModeFactor() != 0.5 {
		t.Errorf("expected operation mode factor 0.5, got %v", instruction.GetOperationModeFactor())
	}
}
