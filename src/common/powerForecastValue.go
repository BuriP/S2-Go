package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

type PowerForecastValue struct {
	CommodityQuantity generated.CommodityQuantity `json:"commodity_quantity" description:"The power quantity the value refers to"`
	ValueExpected     float64                     `json:"value_expected" description:"The expected power value."`
	ValueLower68PPR   *float64                    `json:"value_lower_68PPR,omitempty" description:"The lower boundary of the range with 68% certainty the power value is in it"`
	ValueLower95PPR   *float64                    `json:"value_lower_95PPR,omitempty" description:"The lower boundary of the range with 95% certainty the power value is in it"`
	ValueLowerLimit   *float64                    `json:"value_lower_limit,omitempty" description:"The lower boundary of the range with 100% certainty the power value is in it"`
	ValueUpper68PPR   *float64                    `json:"value_upper_68PPR,omitempty" description:"The upper boundary of the range with 68% certainty the power value is in it"`
	ValueUpper95PPR   *float64                    `json:"value_upper_95PPR,omitempty" description:"The upper boundary of the range with 95% certainty the power value is in it"`
	ValueUpperLimit   *float64                    `json:"value_upper_limit,omitempty" description:"The upper boundary of the range with 100% certainty the power value is in it"`
}

// NewPowerForecast creates a new instance of the PowerForecast value and an error.
func NewPowerForecastValue(commodity generated.CommodityQuantity, valueExp float64, optionalValues ...*float64) (*PowerForecastValue, error) {
	if commodity == "" {
		return nil, errors.New("Commodity is required.")
	}

	if valueExp < 0 {
		return nil, errors.New("Value Expected cannot be negative")
	}

	pfv := &PowerForecastValue{
		CommodityQuantity: commodity,
		ValueExpected:     valueExp,
	}

	if len(optionalValues) > 0 {
		pfv.ValueLower68PPR = optionalValues[0]
	}
	if len(optionalValues) > 1 {
		pfv.ValueLower95PPR = optionalValues[1]
	}
	if len(optionalValues) > 2 {
		pfv.ValueLowerLimit = optionalValues[2]
	}
	if len(optionalValues) > 3 {
		pfv.ValueUpper68PPR = optionalValues[3]
	}
	if len(optionalValues) > 4 {
		pfv.ValueUpper95PPR = optionalValues[4]
	}
	if len(optionalValues) > 5 {
		pfv.ValueUpperLimit = optionalValues[5]
	}

	return pfv, nil
}

