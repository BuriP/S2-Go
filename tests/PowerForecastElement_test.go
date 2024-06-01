package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewPowerForecastElement(t *testing.T) {
	duration := common.NewDuration(3600)

	// Create PowerForecastValues using the constructor function
	pfv1, err := common.NewPowerForecastValue(generated.ELECTRIC_POWER_L1, 100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	pfv2, err := common.NewPowerForecastValue(generated.ELECTRIC_POWER_L2, 200.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	powerValues := []*common.PowerForecastValue{pfv1, pfv2}

	// Test valid PowerForecastElement creation
	pfe, err := common.NewPowerForecastElement(duration, powerValues)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if pfe.Duration == nil || pfe.Duration.Milliseconds != duration.Milliseconds {
		t.Errorf("expected Duration %v, got %v", duration, pfe.Duration)
	}
	if pfe.PowerValues == nil || len(pfe.PowerValues) != len(powerValues) {
		t.Errorf("expected %d power values, got %d", len(powerValues), len(pfe.PowerValues))
	}

	// Test creation with nil duration
	_, err = common.NewPowerForecastElement(nil, powerValues)
	if err == nil {
		t.Fatalf("expected error for nil duration, got nil")
	}

	// Test creation with empty power values
	_, err = common.NewPowerForecastElement(duration, []*common.PowerForecastValue{})
	if err == nil {
		t.Fatalf("expected error for empty power values, got nil")
	}

	// Test creation with nil power values
	_, err = common.NewPowerForecastElement(duration, nil)
	if err == nil {
		t.Fatalf("expected error for nil power values, got nil")
	}
}
