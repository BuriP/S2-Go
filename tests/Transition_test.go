package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewTransition(t *testing.T) {
	// Generate IDs for testing
	id, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new ID: %v", err)
	}
	from, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new ID: %v", err)
	}
	to, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new ID: %v", err)
	}
	startTimer1, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new ID: %v", err)
	}
	startTimer2, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new ID: %v", err)
	}
	blockingTimer1, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new ID: %v", err)
	}
	blockingTimer2, err := generated.NewID()
	if err != nil {
		t.Fatalf("failed to generate new ID: %v", err)
	}

	startTimers := []*generated.ID{startTimer1, startTimer2}
	blockingTimers := []*generated.ID{blockingTimer1, blockingTimer2}
	abnormalConditionOnly := true
	transitionCosts := 100.0
	transitionDuration := common.NewDuration(3600) // 1 hour

	// Test valid Transition creation
	transition, err := common.NewTransition(id, from, to, startTimers, blockingTimers, abnormalConditionOnly, &transitionCosts, transitionDuration)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if transition.ID == nil || transition.ID.Value != id.Value {
		t.Errorf("expected ID %v, got %v", id.Value, transition.ID.Value)
	}
	if transition.From == nil || transition.From.Value != from.Value {
		t.Errorf("expected From %v, got %v", from.Value, transition.From.Value)
	}
	if transition.To == nil || transition.To.Value != to.Value {
		t.Errorf("expected To %v, got %v", to.Value, transition.To.Value)
	}
	if len(transition.StartTimers) != len(startTimers) {
		t.Errorf("expected %d start timers, got %d", len(startTimers), len(transition.StartTimers))
	}
	if len(transition.BlockingTimers) != len(blockingTimers) {
		t.Errorf("expected %d blocking timers, got %d", len(blockingTimers), len(transition.BlockingTimers))
	}
	if transition.AbnormalConditionOnly != abnormalConditionOnly {
		t.Errorf("expected AbnormalConditionOnly %v, got %v", abnormalConditionOnly, transition.AbnormalConditionOnly)
	}
	if transition.TransitionCosts == nil || *transition.TransitionCosts != transitionCosts {
		t.Errorf("expected TransitionCosts %v, got %v", transitionCosts, *transition.TransitionCosts)
	}
	if transition.TransitionDuration == nil || transition.TransitionDuration.Milliseconds != transitionDuration.Milliseconds {
		t.Errorf("expected TransitionDuration %v, got %v", transitionDuration.Milliseconds, transition.TransitionDuration.Milliseconds)
	}

	// Test creation with missing required fields
	_, err = common.NewTransition(nil, from, to, startTimers, blockingTimers, abnormalConditionOnly, &transitionCosts, transitionDuration)
	if err == nil {
		t.Fatalf("expected error for missing id, got nil")
	}

	_, err = common.NewTransition(id, nil, to, startTimers, blockingTimers, abnormalConditionOnly, &transitionCosts, transitionDuration)
	if err == nil {
		t.Fatalf("expected error for missing from, got nil")
	}

	_, err = common.NewTransition(id, from, nil, startTimers, blockingTimers, abnormalConditionOnly, &transitionCosts, transitionDuration)
	if err == nil {
		t.Fatalf("expected error for missing to, got nil")
	}
}
