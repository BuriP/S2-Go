package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

type PowerValue struct {
	CommodityQuantity generated.CommodityQuantity `json:"commodity_quantity" description:"The power quantity the value refers to"`
	Value             float64                     `json:"value" description:"Power value expressed in the unit associated with the CommodityQuantity"`
}

// NewPowerValue creates a new instance of the PowerValue type and returns an error if validation fails.
func NewPowerValue(commodity generated.CommodityQuantity, value float64) (*PowerValue, error) {
	if commodity == "" {
		return nil, errors.New("commodity quantity is required")
	}

	if value < 0 {
		return nil, errors.New("value cannot be negative")
	}

	return &PowerValue{
		CommodityQuantity: commodity,
		Value:             value,
	}, nil
}
