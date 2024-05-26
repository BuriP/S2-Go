package tests

import (
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewPowerForecast(t *testing.T) {
	// Create PowerForecastValues using the constructor function
	pfv1, err := common.NewPowerForecastValue(generated.ELECTRIC_POWER_L1, 100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	pfv2, err := common.NewPowerForecastValue(generated.ELECTRIC_POWER_L1, 200.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Create PowerForecastElements using the constructor function
	duration1 := common.NewDuration(3600) // 1 hour
	element1, err := common.NewPowerForecastElement(duration1, &[]common.PowerForecastValue{*pfv1})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	duration2 := common.NewDuration(7200) // 2 hours
	element2, err := common.NewPowerForecastElement(duration2, &[]common.PowerForecastValue{*pfv2})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	elements := []common.PowerForecastElement{*element1, *element2}
	startTime := time.Now().Unix()

	// Test valid PowerForecast creation
	pf, err := common.NewPowerForecast(&elements, startTime)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if pf.MessageType != "PowerForecast" {
		t.Errorf("expected MessageType PowerForecast, got %v", pf.MessageType)
	}
	if pf.Elements == nil || len(*pf.Elements) != len(elements) {
		t.Errorf("expected %d elements, got %d", len(elements), len(*pf.Elements))
	}

	// Test creation with nil elements
	_, err = common.NewPowerForecast(nil, startTime)
	if err == nil {
		t.Fatalf("expected error for nil elements, got nil")
	}

	// Test creation with empty elements
	emptyElements := []common.PowerForecastElement{}
	_, err = common.NewPowerForecast(&emptyElements, startTime)
	if err == nil {
		t.Fatalf("expected error for empty elements, got nil")
	}

	// Test creation with elements not in chronological order
	durationOutOfOrder := common.NewDuration(1800) // 30 minutes
	elementOutOfOrder, err := common.NewPowerForecastElement(durationOutOfOrder, &[]common.PowerForecastValue{*pfv1})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	elementsOutOfOrder := []common.PowerForecastElement{*element1, *elementOutOfOrder}
	_, err = common.NewPowerForecast(&elementsOutOfOrder, startTime)
	if err == nil {
		t.Fatalf("expected error for elements not in chronological order, got nil")
	}

	// Test creation with an element missing duration
	elementNoDuration := common.PowerForecastElement{
		Duration:    nil,
		PowerValues: &[]common.PowerForecastValue{*pfv1},
	}
	elementsNoDuration := []common.PowerForecastElement{elementNoDuration}
	_, err = common.NewPowerForecast(&elementsNoDuration, startTime)
	if err == nil {
		t.Fatalf("expected error for element with nil duration, got nil")
	}
}
