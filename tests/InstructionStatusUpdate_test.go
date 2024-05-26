package tests

import (
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewInstructionStatusUpdate(t *testing.T) {
	statusTypes := []generated.InstructionStatus{
		generated.NEW,
		generated.ACCEPTED,
		generated.REJECTED,
		generated.REVOKED,
		generated.STARTED,
		generated.SUCCEDED,
		generated.ABORTED,
	}

	for _, statusType := range statusTypes {
		before := time.Now()
		isu, err := common.NewInstructionStatusUpdate(statusType)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if isu.MessageID == nil || isu.MessageID.Value == "" {
			t.Errorf("expected a non-empty MessageID, got %v", isu.MessageID)
		}
		if isu.InstructionID == nil || isu.InstructionID.Value == "" {
			t.Errorf("expected a non-empty InstructionID, got %v", isu.InstructionID)
		}
		if isu.MessageType != "Instruction Status Update" {
			t.Errorf("expected MessageType %v, got %v", "test-message", isu.MessageType)
		}
		if isu.StatusType != statusType {
			t.Errorf("expected StatusType %v, got %v", statusType, isu.StatusType)
		}
		if isu.Timestamp.Before(before) {
			t.Errorf("expected Timestamp to be after %v, got %v", before, isu.Timestamp)
		}
		if isu.Timestamp.After(time.Now()) {
			t.Errorf("expected Timestamp to be before now, got %v", isu.Timestamp)
		}
	}
}

func TestGetInstructionStatusUpdateMessageID(t *testing.T) {
	msgID, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new ID: %v", err)
	}

	isu := &common.InstructionStatusUpdate{
		MessageID: msgID,
	}

	if isu.GetMessageID() == nil || isu.GetMessageID().Value != msgID.Value {
		t.Errorf("expected %v, got %v", msgID.Value, isu.GetMessageID().Value)
	}
}

func TestGetInstructionID(t *testing.T) {
	insID, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new ID: %v", err)
	}

	isu := &common.InstructionStatusUpdate{
		InstructionID: insID,
	}

	if isu.GetInstructionID() == nil || isu.GetInstructionID().Value != insID.Value {
		t.Errorf("expected %v, got %v", insID.Value, isu.GetInstructionID().Value)
	}
}
