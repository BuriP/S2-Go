// TO DO : Look at wether it should be an array or a slice.

package generated

import (
	"time"
)

// FRBCUsageForecast represents the usage forecast.
type FRBCUsageForecast struct {
    Elements    []FRBCUsageForecastElement `json:"elements" description:"Further elements that model the profile. There shall be at least one element. Elements must be placed in chronological order." maxItems:"288" minItems:"1"` // Using a slice here. 
    MessageID   ID                         `json:"message_id" description:"ID of this message"`
    MessageType string                     `json:"message_type" description:"FRBC.UsageForecast" const:"true"`
    StartTime   time.Time                  `json:"start_time" description:"Time at which the FRBC.UsageForecast starts."`
}
