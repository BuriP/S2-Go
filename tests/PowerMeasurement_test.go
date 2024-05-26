package tests

import (
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewPowerMeasurement(t *testing.T) {
	values := []common.PowerValue{
		{CommodityQuantity: generated.ELECTRIC_POWER_L1, Value: 100.0},
		{CommodityQuantity: generated.ELECTRIC_POWER_L2, Value: 200.0},
	}

	// Getting the time before the instance
	before := time.Now()

	// Test valid PowerMeasurement creation
	pm, err := common.NewPowerMeasurement(&values)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Getting the time after the instance
	after := time.Now()

	if pm.MessageID == nil || pm.MessageID.Value == "" {
		t.Errorf("expected a non-empty MessageID, got %v", pm.MessageID)
	}
	if pm.MessageType != "PowerMeasurement" {
		t.Errorf("expected MessageType %v, got %v", "PowerMeasurement", pm.MessageType)
	}
	if pm.MeasurementTimestamp.Before(before) || pm.MeasurementTimestamp.After(after) {
		t.Errorf("expected MeasurementTimeStamp to be between %v and %v, got %v", before, after, pm.MeasurementTimestamp)
	}
	if pm.Values == nil || len(*pm.Values) != len(values) {
		t.Errorf("expected %d values, got %d", len(values), len(*pm.Values))
	}

	// Test empty values list
	_, err = common.NewPowerMeasurement(&[]common.PowerValue{})
	if err == nil {
		t.Fatalf("expected error for empty values, got nil")
	}

	// Test nil values list
	_, err = common.NewPowerMeasurement(nil)
	if err == nil {
		t.Fatalf("expected error for nil values, got nil")
	}
}
