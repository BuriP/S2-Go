package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewPowerForecastValue(t *testing.T) {
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

	valueExpected := 100.0
	lower68 := 90.0
	lower95 := 80.0
	lowerLimit := 70.0
	upper68 := 110.0
	upper95 := 120.0
	upperLimit := 130.0

	for _, commodity := range commodityQuantities {
		// Test valid PowerForecastValue creation with no optional values
		pfv, err := common.NewPowerForecastValue(commodity, valueExpected)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if pfv.CommodityQuantity == "" || pfv.CommodityQuantity != commodity {
			t.Errorf("expected CommodityQuantity %v, got %v", commodity, pfv.CommodityQuantity)
		}
		if pfv.ValueExpected != valueExpected {
			t.Errorf("expected ValueExpected %v, got %v", valueExpected, pfv.ValueExpected)
		}

		// Test valid PowerForecastValue creation with optional values
		pfv, err = common.NewPowerForecastValue(commodity, valueExpected, &lower68, &lower95, &lowerLimit, &upper68, &upper95, &upperLimit)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if pfv.ValueLower68PPR == nil || *pfv.ValueLower68PPR != lower68 {
			t.Errorf("expected ValueLower68PPR %v, got %v", lower68, pfv.ValueLower68PPR)
		}
		if pfv.ValueLower95PPR == nil || *pfv.ValueLower95PPR != lower95 {
			t.Errorf("expected ValueLower95PPR %v, got %v", lower95, pfv.ValueLower95PPR)
		}
		if pfv.ValueLowerLimit == nil || *pfv.ValueLowerLimit != lowerLimit {
			t.Errorf("expected ValueLowerLimit %v, got %v", lowerLimit, pfv.ValueLowerLimit)
		}
		if pfv.ValueUpper68PPR == nil || *pfv.ValueUpper68PPR != upper68 {
			t.Errorf("expected ValueUpper68PPR %v, got %v", upper68, pfv.ValueUpper68PPR)
		}
		if pfv.ValueUpper95PPR == nil || *pfv.ValueUpper95PPR != upper95 {
			t.Errorf("expected ValueUpper95PPR %v, got %v", upper95, pfv.ValueUpper95PPR)
		}
		if pfv.ValueUpperLimit == nil || *pfv.ValueUpperLimit != upperLimit {
			t.Errorf("expected ValueUpperLimit %v, got %v", upperLimit, pfv.ValueUpperLimit)
		}
	}

	// Test creation with nil commodity quantity
	_, err := common.NewPowerForecastValue("", valueExpected)
	if err == nil {
		t.Fatalf("expected error for nil commodity quantity, got nil")
	}

}
