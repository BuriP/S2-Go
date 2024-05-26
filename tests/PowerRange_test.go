package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewPowerRange(t *testing.T) {
	commodityQuantities := []generated.CommodityQuantity{
		generated.ELECTRIC_POWER_L1,
		generated.ELECTRIC_POWER_L2,
		generated.ELECTRIC_POWER_L3,
		generated.ELECTRIC_POWER_3_PHASE_SYMMETRIC,
		generated.NATURAL_GAS_FLOW_RATE,
		generated.HYDROGEN_FLOW_RATE,
		generated.HEAT_TEMPERATURE,
		generated.HEAT_FLOW_RATE,
		generated.HEAT_THERMAL_POWER,
		generated.OIL_FLOW_RATE,
	}

	startOfRange := 50.0
	endOfRange := 100.0

	for _, commodity := range commodityQuantities {
		// Test valid PowerRange creation
		pr, err := common.NewPowerRange(commodity, startOfRange, endOfRange)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if pr.CommodityQuantity != commodity {
			t.Errorf("expected CommodityQuantity %v, got %v", commodity, pr.CommodityQuantity)
		}
		if pr.StartOfRange != startOfRange {
			t.Errorf("expected StartOfRange %v, got %v", startOfRange, pr.StartOfRange)
		}
		if pr.EndOfRange != endOfRange {
			t.Errorf("expected EndOfRange %v, got %v", endOfRange, pr.EndOfRange)
		}
	}

	// Test creation with empty commodity quantity
	_, err := common.NewPowerRange("", startOfRange, endOfRange)
	if err == nil {
		t.Fatalf("expected error for empty commodity quantity, got nil")
	}

	// Test creation with start greater than end
	_, err = common.NewPowerRange(generated.ELECTRIC_POWER_L1, 200.0, 100.0)
	if err == nil {
		t.Fatalf("expected error for start greater than end, got nil")
	}
}
