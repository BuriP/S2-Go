package frbc

import (
	"errors"
	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCUsageForecastElement represents an element of FRBC usage forecast.
type FRBCUsageForecastElement struct {
	Duration            *common.Duration `json:"duration" description:"Indicator for how long the given usage_rate is valid."`
	UsageRateExpected   float64          `json:"usage_rate_expected" description:"The most likely value for the usage rate; the expected increase or decrease of the fill_level per second"`
	UsageRateLower68PPR *float64         `json:"usage_rate_lower_68PPR,omitempty" description:"The lower limit of the range with a 68% probability that the usage rate is within that range"`   // Used as pointer as in case the are not included.
	UsageRateLower95PPR *float64         `json:"usage_rate_lower_95PPR,omitempty" description:"The lower limit of the range with a 95% probability that the usage rate is within that range"`   // Used as pointer as in case the are not included.
	UsageRateLowerLimit *float64         `json:"usage_rate_lower_limit,omitempty" description:"The lower limit of the range with a 100% probability that the usage rate is within that range"`  // Used as pointer as in case the are not included.
	UsageRateUpper68PPR *float64         `json:"usage_rate_upper_68PPR,omitempty" description:"The upper limit of the range with a 68% probability that the usage rate is within that range"`   // Used as pointer as in case the are not included.
	UsageRateUpper95PPR *float64         `json:"usage_rate_upper_95PPR,omitempty" description:"The upper limit of the range with a 95% probability that the usage rate is within that range."`  // Used as pointer as in case the are not included.
	UsageRateUpperLimit *float64         `json:"usage_rate_upper_limit,omitempty" description:"The upper limit of the range with a 100% probability that the usage rate is within that range."` // Used as pointer as in case the are not included.
}

// NewFRBCUsageForecastElement creates a new instance if validation is correct, otherwise error.
func NewFRBCUsageForecastElement(duration *common.Duration, usageRateExp float64) (*FRBCUsageForecastElement, error) {
	if duration == nil {
		return nil, errors.New("a duration must be provided")
	}

	if usageRateExp == 0 {
		return nil, errors.New("a usage rate expected must be provided")
	}

	return &FRBCUsageForecastElement{
		Duration:          duration,
		UsageRateExpected: usageRateExp,
	}, nil
}
