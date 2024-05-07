package common

import (
	"github.com/BuriP/S2-Go/src/generated"
)

type PowerValue struct {
	CommodityQuantity generated.CommodityQuantity `json:"commodity_quantity" description:"The power quantity the value refers to"`
	Value             float64           `json:"value" description:"Power value expressed in the unit associated with the CommodityQuantity"`
}

// NewPowerValue creates a new instance of the PowerValue type. 
func NewPowerValue(commodity generated.CommodityQuantity, value float64) *PowerValue{
	return &PowerValue{
		CommodityQuantity: commodity,
		Value: value,
	}
}