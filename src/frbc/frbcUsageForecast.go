package frbc

import (
	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
	"time"
)

// FRBCUsageForecast represents the usage forecast.
type FRBCUsageForecast struct {
	Elements    []generated.FRBCUsageForecastElement `json:"elements" description:"Further elements that model the profile. There shall be at least one element. Elements must be placed in chronological order." maxItems:"288" minItems:"1"` // Using a slice here.
	MessageID   *generated.ID                        `json:"message_id" description:"ID of this message"`
	MessageType string                               `json:"message_type" description:"FRBC.UsageForecast" const:"true"`
	StartTime   time.Time                            `json:"start_time" description:"Time at which the FRBC.UsageForecast starts."`
}
