package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

// CommodityHasQuantity returns a bool indicating whether the commodity has the respective quantity.
// If the commodity is unsupported, it returns an error.
func CommodityHasQuantity(commodity generated.Commodity, quantity generated.CommodityQuantity) (bool, error) {
	switch commodity {
	case generated.HEAT:
		return quantity == generated.HEAT_THERMAL_POWER || quantity == generated.HEAT_TEMPERATURE || quantity == generated.HEAT_FLOW_RATE, nil
	case generated.ELECTRICITY:
		return quantity == generated.ELECTRIC_POWER_3_PHASE_SYMMETRIC || quantity == generated.ELECTRIC_POWER_L1 ||
			quantity == generated.ELECTRIC_POWER_L2 || quantity == generated.ELECTRIC_POWER_L3, nil
	case generated.GAS:
		return quantity == generated.NATURAL_GAS_FLOW_RATE, nil
	case generated.OIL:
		return quantity == generated.OIL_FLOW_RATE, nil
	default:
		return false, errors.New("unsupported commodity: " + string(commodity))
	}
}

