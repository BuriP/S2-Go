package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewPowerValue(t *testing.T) {
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

	validValue := 100.0

	for _, commodity := range commodityQuantities {
		// Test valid PowerValue creation
		pv, err := common.NewPowerValue(commodity, validValue)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if pv.CommodityQuantity != commodity {
			t.Errorf("expected CommodityQuantity %v, got %v", commodity, pv.CommodityQuantity)
		}
		if pv.Value != validValue {
			t.Errorf("expected Value %v, got %v", validValue, pv.Value)
		}
	}

	// Test creation with empty commodity quantity
	_, err := common.NewPowerValue("", validValue)
	if err == nil {
		t.Fatalf("expected error for empty commodity quantity, got nil")
	}

	// Test creation with negative value
	invalidValue := -50.0
	commodity := generated.ELECTRIC_POWER_L1
	_, err = common.NewPowerValue(commodity, invalidValue)
	if err == nil {
		t.Fatalf("expected error for negative value, got nil")
	}
}
