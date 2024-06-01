package common

import (
	"errors"
)

type PowerForecastElement struct {
	Duration    *Duration             `json:"duration" description:"Duration of the PowerForecastElement"`
	PowerValues []*PowerForecastValue `json:"power_values" description:"The values of power that are expected for the given period of time."`
}

// NewPowerForecastElement creates a new PowerForecastElement instance and returns an error if validation fails
func NewPowerForecastElement(duration *Duration, powerValues []*PowerForecastValue) (*PowerForecastElement, error) {
	if duration == nil {
		return nil, errors.New("duration is required")
	}

	if powerValues == nil || len(powerValues) == 0 {
		return nil, errors.New("power values must contain at least one element")
	}

	return &PowerForecastElement{
		Duration:    duration,
		PowerValues: powerValues,
	}, nil
}
