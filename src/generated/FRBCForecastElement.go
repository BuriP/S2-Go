package generated

import ()

// FRBCUsageForecastElement represents an element of FRBC usage forecast.
type FRBCUsageForecastElement struct {
	Duration               Duration `json:"duration" description:"Indicator for how long the given usage_rate is valid."`
	UsageRateExpected      float64  `json:"usage_rate_expected" description:"The most likely value for the usage rate; the expected increase or decrease of the fill_level per second"`
	UsageRateLower68PPR    *float64 `json:"usage_rate_lower_68PPR,omitempty" description:"The lower limit of the range with a 68% probability that the usage rate is within that range"` // Used as pointer as in case the are not included.
	UsageRateLower95PPR    *float64 `json:"usage_rate_lower_95PPR,omitempty" description:"The lower limit of the range with a 95% probability that the usage rate is within that range"` // Used as pointer as in case the are not included.
	UsageRateLowerLimit    *float64 `json:"usage_rate_lower_limit,omitempty" description:"The lower limit of the range with a 100% probability that the usage rate is within that range"` // Used as pointer as in case the are not included.
	UsageRateUpper68PPR    *float64 `json:"usage_rate_upper_68PPR,omitempty" description:"The upper limit of the range with a 68% probability that the usage rate is within that range"` // Used as pointer as in case the are not included.
	UsageRateUpper95PPR    *float64 `json:"usage_rate_upper_95PPR,omitempty" description:"The upper limit of the range with a 95% probability that the usage rate is within that range."` // Used as pointer as in case the are not included.
	UsageRateUpperLimit    *float64 `json:"usage_rate_upper_limit,omitempty" description:"The upper limit of the range with a 100% probability that the usage rate is within that range."` // Used as pointer as in case the are not included.
}