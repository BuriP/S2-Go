package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewFRBCTimerStatus(t *testing.T) {
	actuatorID, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new actuator ID: %v", err)
	}
	timerID, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new timer ID: %v", err)
	}

	// Create Actuator and Timer
	actuator := &common.Actuator{ID: actuatorID}
	timer := &common.Timer{ID: timerID}
	finishedAt := time.Now()

	// Test valid FRBCTimerStatus creation
	timerStatus, err := frbc.NewFRBCTimerStatus(actuator, timer, finishedAt)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if timerStatus.ActuatorID == nil || timerStatus.ActuatorID.Value != actuatorID.Value {
		t.Errorf("expected ActuatorID %v, got %v", actuatorID.Value, timerStatus.ActuatorID.Value)
	}
	if timerStatus.TimerID == nil || timerStatus.TimerID.Value != timerID.Value {
		t.Errorf("expected TimerID %v, got %v", timerID.Value, timerStatus.TimerID.Value)
	}
	if timerStatus.MessageID == nil {
		t.Errorf("expected non-nil MessageID")
	}
	if timerStatus.MessageType != "FRBC.TimerStatus" {
		t.Errorf("expected MessageType FRBC.TimerStatus, got %v", timerStatus.MessageType)
	}
	if timerStatus.FinishedAt != finishedAt {
		t.Errorf("expected FinishedAt %v, got %v", finishedAt, timerStatus.FinishedAt)
	}

	// Test creation with nil actuator
	_, err = frbc.NewFRBCTimerStatus(nil, timer, finishedAt)
	if err == nil {
		t.Fatalf("expected error for nil actuator, got nil")
	}

	// Test creation with nil timer
	_, err = frbc.NewFRBCTimerStatus(actuator, nil, finishedAt)
	if err == nil {
		t.Fatalf("expected error for nil timer, got nil")
	}
}

func TestFRBCTimerStatusSerialization(t *testing.T) {
	actuatorID, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new actuator ID: %v", err)
	}
	timerID, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new timer ID: %v", err)
	}

	// Create Actuator and Timer
	actuator := &common.Actuator{ID: actuatorID}
	timer := &common.Timer{ID: timerID}
	finishedAt := time.Now()

	timerStatus, err := frbc.NewFRBCTimerStatus(actuator, timer, finishedAt)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(timerStatus)
	if err != nil {
		t.Fatalf("failed to marshal FRBCTimerStatus to JSON: %v", err)
	}

	// Deserialize from JSON
	var deserializedTimerStatus frbc.FRBCTimerStatus
	err = json.Unmarshal(jsonData, &deserializedTimerStatus)
	if err != nil {
		t.Fatalf("failed to unmarshal JSON to FRBCTimerStatus: %v", err)
	}

	// Compare original and deserialized values
	if deserializedTimerStatus.ActuatorID == nil || deserializedTimerStatus.ActuatorID.Value != actuatorID.Value {
		t.Errorf("expected ActuatorID %v, got %v", actuatorID.Value, deserializedTimerStatus.ActuatorID.Value)
	}
	if deserializedTimerStatus.TimerID == nil || deserializedTimerStatus.TimerID.Value != timerID.Value {
		t.Errorf("expected TimerID %v, got %v", timerID.Value, deserializedTimerStatus.TimerID.Value)
	}
	if deserializedTimerStatus.MessageID == nil {
		t.Errorf("expected non-nil MessageID")
	}
	if deserializedTimerStatus.MessageType != "FRBC.TimerStatus" {
		t.Errorf("expected MessageType FRBC.TimerStatus, got %v", deserializedTimerStatus.MessageType)
	}
	if deserializedTimerStatus.FinishedAt.Unix() != finishedAt.Unix() {
		t.Errorf("expected FinishedAt %v, got %v", finishedAt, deserializedTimerStatus.FinishedAt)
	}
}
