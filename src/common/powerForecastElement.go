package common

import(
	"github.com/BuriP/S2-Go/src/generated"
)

type PowerForecastElement struct {
    Duration     *Duration              `json:"duration" description:"Duration of the PowerForecastElement"`
    PowerValues  *[]generated.PowerForecastValue `json:"power_values" description:"The values of power that are expected for the given period of time."` // USed a slice here of type PowerForecastValue.
}

func NewPowerForecastElement(duration *Duration, powerval *[]generated.PowerForecastValue) *PowerForecastElement{
	return &PowerForecastElement{
		Duration: duration,
		PowerValues: powerval,
	}
}