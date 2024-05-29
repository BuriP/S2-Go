package tests

import (
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
)

func TestNewFRBCUsageForecast(t *testing.T) {
	duration1 := common.NewDuration(3600) // 1 hour
	duration2 := common.NewDuration(7200) // 2 hours

	element1, err := frbc.NewFRBCUsageForecastElement(duration1, 50.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	element2, err := frbc.NewFRBCUsageForecastElement(duration2, 75.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	elements := []frbc.FRBCUsageForecastElement{*element1, *element2}
	startTime := time.Now()

	// Test valid FRBCUsageForecast creation
	forecast, err := frbc.NewFRBCUsageForecast(elements, startTime)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if forecast.MessageID == nil {
		t.Errorf("expected non-nil MessageID")
	}
	if forecast.MessageType != "FRBC.UsageForecast" {
		t.Errorf("expected MessageType FRBC.UsageForecast, got %v", forecast.MessageType)
	}
	if forecast.StartTime != startTime {
		t.Errorf("expected StartTime %v, got %v", startTime, forecast.StartTime)
	}
	if len(forecast.Elements) != len(elements) {
		t.Errorf("expected %d elements, got %d", len(elements), len(forecast.Elements))
	}

	// Test creation with empty elements
	_, err = frbc.NewFRBCUsageForecast([]frbc.FRBCUsageForecastElement{}, startTime)
	if err == nil {
		t.Fatalf("expected error for empty elements, got nil")
	}

	// Test creation with elements not in chronological order
	durationOutOfOrder := common.NewDuration(1800) // 30 minutes
	elementOutOfOrder, err := frbc.NewFRBCUsageForecastElement(durationOutOfOrder, 45.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	elementsOutOfOrder := []frbc.FRBCUsageForecastElement{*element1, *elementOutOfOrder}
	_, err = frbc.NewFRBCUsageForecast(elementsOutOfOrder, startTime)
	if err == nil {
		t.Fatalf("expected error for elements not in chronological order, got nil")
	}
}
