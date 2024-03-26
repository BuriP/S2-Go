package generated

import (
	
)

// PowerForecastValue represents a forecasted power value along with associated metadata.
type PowerForecastValue struct {
	CommodityQuantity   CommodityQuantity `json:"commodity_quantity" description:"The power quantity the value refers to"`
	ValueExpected       float64           `json:"value_expected" description:"The expected power value."`
	ValueLower68PPR     *float64          `json:"value_lower_68PPR,omitempty" description:"The lower boundary of the range with 68% certainty the power value is in it"`
	ValueLower95PPR     *float64          `json:"value_lower_95PPR,omitempty" description:"The lower boundary of the range with 95% certainty the power value is in it"`
	ValueLowerLimit     *float64          `json:"value_lower_limit,omitempty" description:"The lower boundary of the range with 100% certainty the power value is in it"`
	ValueUpper68PPR     *float64          `json:"value_upper_68PPR,omitempty" description:"The upper boundary of the range with 68% certainty the power value is in it"`
	ValueUpper95PPR     *float64          `json:"value_upper_95PPR,omitempty" description:"The upper boundary of the range with 95% certainty the power value is in it"`
	ValueUpperLimit     *float64          `json:"value_upper_limit,omitempty" description:"The upper boundary of the range with 100% certainty the power value is in it"`
}
