package frbc

import (
	"errors"
	"time"

	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCUsageForecast represents the usage forecast.
type FRBCUsageForecast struct {
	Elements    []FRBCUsageForecastElement `json:"elements" description:"Further elements that model the profile. There shall be at least one element. Elements must be placed in chronological order." maxItems:"288" minItems:"1"`
	MessageID   *generated.ID              `json:"message_id" description:"ID of this message"`
	MessageType string                     `json:"message_type" description:"FRBC.UsageForecast" const:"true"`
	StartTime   time.Time                  `json:"start_time" description:"Time at which the FRBC.UsageForecast starts."`
}

// NewFRBCUsageForecast creates a new instance of FRBCUsageForecast if validation is correct.
func NewFRBCUsageForecast(elements []FRBCUsageForecastElement, startTime time.Time) (*FRBCUsageForecast, error) {
	if len(elements) == 0 {
		return nil, errors.New("elements must contain at least one element")
	}

	// Ensure elements are in chronological order
	for i := 1; i < len(elements); i++ {
		if elements[i].Duration == nil || elements[i-1].Duration == nil {
			return nil, errors.New("each element must have a duration")
		}
		if elements[i].Duration.Milliseconds < elements[i-1].Duration.Milliseconds {
			return nil, errors.New("elements must be in chronological order")
		}
	}

	messageID, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &FRBCUsageForecast{
		Elements:    elements,
		MessageID:   messageID,
		MessageType: "FRBC.UsageForecast",
		StartTime:   startTime,
	}, nil
}
