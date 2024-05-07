package common

import(
	"github.com/BuriP/S2-Go/src/generated"
)

type PowerForecastValue struct {
	CommodityQuantity   *generated.CommodityQuantity `json:"commodity_quantity" description:"The power quantity the value refers to"`
	ValueExpected       float64           `json:"value_expected" description:"The expected power value."`
	ValueLower68PPR     *float64          `json:"value_lower_68PPR,omitempty" description:"The lower boundary of the range with 68% certainty the power value is in it"`
	ValueLower95PPR     *float64          `json:"value_lower_95PPR,omitempty" description:"The lower boundary of the range with 95% certainty the power value is in it"`
	ValueLowerLimit     *float64          `json:"value_lower_limit,omitempty" description:"The lower boundary of the range with 100% certainty the power value is in it"`
	ValueUpper68PPR     *float64          `json:"value_upper_68PPR,omitempty" description:"The upper boundary of the range with 68% certainty the power value is in it"`
	ValueUpper95PPR     *float64          `json:"value_upper_95PPR,omitempty" description:"The upper boundary of the range with 95% certainty the power value is in it"`
	ValueUpperLimit     *float64          `json:"value_upper_limit,omitempty" description:"The upper boundary of the range with 100% certainty the power value is in it"`
}

func NewPowerForecastValue (commodity *generated.CommodityQuantity, valueExp float64,lower68 *float64, lower95 *float64, lowerlim *float64, upper68 *float64, upper95 *float64, upperlim *float64 ) *PowerForecastValue{
	return &PowerForecastValue{
		CommodityQuantity: commodity,
		ValueExpected: valueExp,
		ValueLower68PPR: lower68,
		ValueLower95PPR: lower95,
		ValueLowerLimit: lowerlim,
		ValueUpper68PPR: upper68,
		ValueUpper95PPR: upper95,
		ValueUpperLimit: upperlim,
	}
}