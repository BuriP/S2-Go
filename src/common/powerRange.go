package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

type PowerRange struct {
	CommodityQuantity generated.CommodityQuantity `json:"commodity_quantity" description:"The power quantity the values refer to"`
	EndOfRange        float64                     `json:"end_of_range" description:"Power value that defines the end of the range."`
	StartOfRange      float64                     `json:"start_of_range" description:"Power value that defines the start of the range."`
}

// NewPowerRange creates a new PowerRange instance and returns an error if validation fails
func NewPowerRange(commodity generated.CommodityQuantity, start float64, end float64) (*PowerRange, error) {
	if commodity == "" {
		return nil, errors.New("commodity quantity is required")
	}

	if start > end {
		return nil, errors.New("start of range must be less than or equal to end of range")
	}

	return &PowerRange{
		CommodityQuantity: commodity,
		EndOfRange:        end,
		StartOfRange:      start,
	}, nil
}
