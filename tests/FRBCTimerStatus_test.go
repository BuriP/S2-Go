package tests

import (
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewFRBCTimerStatus(t *testing.T) {
	actuatorID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	timerDuration := common.NewDuration(3600) // 1 hour
	timer, err := common.NewTimer(timerDuration)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	finishedAt := time.Now()

	// Test valid FRBCTimerStatus creation
	timerStatus, err := frbc.NewFRBCTimerStatus(actuatorID, timer, finishedAt)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if timerStatus.ActuatorID == nil || timerStatus.ActuatorID.Value != actuatorID.Value {
		t.Errorf("expected actuator ID %v, got %v", actuatorID.Value, timerStatus.ActuatorID.Value)
	}
	if timerStatus.TimerID == nil || timerStatus.TimerID.Value != timer.ID.Value {
		t.Errorf("expected timer ID %v, got %v", timer.ID.Value, timerStatus.TimerID.Value)
	}
	if !timerStatus.FinishedAt.Equal(finishedAt) {
		t.Errorf("expected finished at time %v, got %v", finishedAt, timerStatus.FinishedAt)
	}
	if timerStatus.MessageID == nil || timerStatus.MessageID.Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", timerStatus.MessageID)
	}
	if timerStatus.MessageType != "FRBC.TimerStatus" {
		t.Errorf("expected message type FRBC.TimerStatus, got %v", timerStatus.MessageType)
	}

	// Test creation with nil actuator ID
	_, err = frbc.NewFRBCTimerStatus(nil, timer, finishedAt)
	if err == nil {
		t.Fatalf("expected error for nil actuator ID, got nil")
	}

	// Test creation with nil timer
	_, err = frbc.NewFRBCTimerStatus(actuatorID, nil, finishedAt)
	if err == nil {
		t.Fatalf("expected error for nil timer, got nil")
	}

	// Test creation with empty actuator ID
	emptyActuatorID := &generated.ID{Value: ""}
	_, err = frbc.NewFRBCTimerStatus(emptyActuatorID, timer, finishedAt)
	if err == nil {
		t.Fatalf("expected error for empty actuator ID, got nil")
	}

	// Test creation with nil timer ID
	timer.ID = nil
	_, err = frbc.NewFRBCTimerStatus(actuatorID, timer, finishedAt)
	if err == nil {
		t.Fatalf("expected error for nil timer ID, got nil")
	}
}

func TestFRBCTimerStatusMethods(t *testing.T) {
	actuatorID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	timerDuration := common.NewDuration(3600) // 1 hour
	timer, err := common.NewTimer(timerDuration)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	finishedAt := time.Now()

	timerStatus, err := frbc.NewFRBCTimerStatus(actuatorID, timer, finishedAt)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetActuatorID method
	if timerStatus.ActuatorID == nil || timerStatus.ActuatorID.Value != actuatorID.Value {
		t.Errorf("expected actuator ID %v, got %v", actuatorID.Value, timerStatus.ActuatorID.Value)
	}

	// Test GetTimerID method
	if timerStatus.TimerID == nil || timerStatus.TimerID.Value != timer.ID.Value {
		t.Errorf("expected timer ID %v, got %v", timer.ID.Value, timerStatus.TimerID.Value)
	}

	// Test GetFinishedAt method
	if !timerStatus.FinishedAt.Equal(finishedAt) {
		t.Errorf("expected finished at time %v, got %v", finishedAt, timerStatus.FinishedAt)
	}

	// Test GetMessageID method
	if timerStatus.MessageID == nil || timerStatus.MessageID.Value == "" {
		t.Errorf("expected a non-empty message ID, got %v", timerStatus.MessageID)
	}

	// Test GetMessageType method
	if timerStatus.MessageType != "FRBC.TimerStatus" {
		t.Errorf("expected message type FRBC.TimerStatus, got %v", timerStatus.MessageType)
	}
}
