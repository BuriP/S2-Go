package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestCommodityHasQuantity(t *testing.T) {
	// Test cases for each commodity with valid quantities
	testCases := []struct {
		commodity generated.Commodity
		quantity  generated.CommodityQuantity
		expected  bool
	}{
		{generated.HEAT, generated.HEAT_THERMAL_POWER, true},
		{generated.HEAT, generated.HEAT_TEMPERATURE, true},
		{generated.HEAT, generated.HEAT_FLOW_RATE, true},
		{generated.ELECTRICITY, generated.ELECTRIC_POWER_3_PHASE_SYMMETRIC, true},
		{generated.ELECTRICITY, generated.ELECTRIC_POWER_L1, true},
		{generated.ELECTRICITY, generated.ELECTRIC_POWER_L2, true},
		{generated.ELECTRICITY, generated.ELECTRIC_POWER_L3, true},
		{generated.GAS, generated.NATURAL_GAS_FLOW_RATE, true},
		{generated.OIL, generated.OIL_FLOW_RATE, true},
	}

	for _, tc := range testCases {
		result, err := common.CommodityHasQuantity(tc.commodity, tc.quantity)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if result != tc.expected {
			t.Errorf("expected %v, got %v for commodity %v and quantity %v", tc.expected, result, tc.commodity, tc.quantity)
		}
	}

	// Test unsupported commodity
	unsupportedCommodity := generated.Commodity("UNSUPPORTED_COMMODITY")
	_, err := common.CommodityHasQuantity(unsupportedCommodity, generated.CommodityQuantity("ANY_QUANTITY"))
	if err == nil {
		t.Fatalf("expected error for unsupported commodity, got nil")
	}
}

func TestCommodityHasQuantityWithInvalidQuantity(t *testing.T) {
	// Test cases for each commodity with invalid quantities
	testCases := []struct {
		commodity generated.Commodity
		quantity  generated.CommodityQuantity
		expected  bool
	}{
		{generated.HEAT, generated.CommodityQuantity("INVALID_QUANTITY"), false},
		{generated.ELECTRICITY, generated.CommodityQuantity("INVALID_QUANTITY"), false},
		{generated.GAS, generated.CommodityQuantity("INVALID_QUANTITY"), false},
		{generated.OIL, generated.CommodityQuantity("INVALID_QUANTITY"), false},
	}

	for _, tc := range testCases {
		result, err := common.CommodityHasQuantity(tc.commodity, tc.quantity)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if result != tc.expected {
			t.Errorf("expected %v, got %v for commodity %v and quantity %v", tc.expected, result, tc.commodity, tc.quantity)
		}
	}
}
