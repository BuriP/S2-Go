package common

import(
	"github.com/BuriP/S2-Go/src/generated"
)

// CommodityHasQuantity return a bool wheter commodity has the respective quantity
func CommodityHasQuantity(commodity generated.Commodity, quantity generated.CommodityQuantity) bool {
	switch commodity {
	case generated.HEAT:
		return quantity == generated.HEAT_THERMAL_POWER || quantity == generated.HEAT_TEMPERATURE || quantity == generated.HEAT_FLOW_RATE
	case generated.ELECTRICITY:
		return quantity == generated.ELECTRIC_POWER_3_PHASE_SYMMETRIC || quantity == generated.ELECTRIC_POWER_L1 ||
			quantity == generated.ELECTRIC_POWER_L2 || quantity == generated.ELECTRIC_POWER_L3
	case generated.GAS:
		return quantity == generated.NATURAL_GAS_FLOW_RATE
	case generated.OIL:
		return quantity == generated.OIL_FLOW_RATE
	default:
		panic("Unsupported commodity: " + string(commodity))
	}
}