package tests

import (
	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
	"testing"
)

func TestNewFRBCUsageForecastElement(t *testing.T) {
	duration := common.NewDuration(3600) // 1 hour
	usageRateExpected := 50.0

	// Test valid FRBCUsageForecastElement creation
	element, err := frbc.NewFRBCUsageForecastElement(duration, usageRateExpected)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if element.Duration == nil || element.Duration.Milliseconds != duration.Milliseconds {
		t.Errorf("expected duration %v, got %v", duration.Milliseconds, element.Duration.Milliseconds)
	}
	if element.UsageRateExpected != usageRateExpected {
		t.Errorf("expected usage rate expected %v, got %v", usageRateExpected, element.UsageRateExpected)
	}

	// Test creation with nil duration
	_, err = frbc.NewFRBCUsageForecastElement(nil, usageRateExpected)
	if err == nil {
		t.Fatalf("expected error for nil duration, got nil")
	}

	// Test creation with zero usage rate expected
	_, err = frbc.NewFRBCUsageForecastElement(duration, 0)
	if err == nil {
		t.Fatalf("expected error for zero usage rate expected, got nil")
	}
}

func TestOptionalFieldsFRBCUsageForecastElement(t *testing.T) {
	duration := common.NewDuration(3600) // 1 hour
	usageRateExpected := 50.0
	usageRateLower68PPR := 45.0
	usageRateUpper68PPR := 55.0

	element, err := frbc.NewFRBCUsageForecastElement(duration, usageRateExpected)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	element.UsageRateLower68PPR = &usageRateLower68PPR
	element.UsageRateUpper68PPR = &usageRateUpper68PPR

	if element.UsageRateLower68PPR == nil || *element.UsageRateLower68PPR != usageRateLower68PPR {
		t.Errorf("expected usage rate lower 68PPR %v, got %v", usageRateLower68PPR, element.UsageRateLower68PPR)
	}
	if element.UsageRateUpper68PPR == nil || *element.UsageRateUpper68PPR != usageRateUpper68PPR {
		t.Errorf("expected usage rate upper 68PPR %v, got %v", usageRateUpper68PPR, element.UsageRateUpper68PPR)
	}
}
