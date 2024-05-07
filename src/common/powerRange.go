package common

import (
	"github.com/BuriP/S2-Go/src/generated"
	
)


type PowerRange struct {
	CommodityQuantity generated.CommodityQuantity `json:"commodity_quantity" description:"The power quantity the values refer to"`
	EndOfRange        float64           `json:"end_of_range" description:"Power value that defines the end of the range."`
	StartOfRange      float64           `json:"start_of_range" description:"Power value that defines the start of the range."`
}

// NewPowerRange creates a new PowerRange instance
func NewPowerRange(commodity generated.CommodityQuantity, end float64, start float64) *PowerRange{
	return &PowerRange{
		EndOfRange: end,
		StartOfRange: start,
		CommodityQuantity: commodity,
	}

}