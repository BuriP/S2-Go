package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
	"time"
)

type PowerForecast struct {
	Elements    *[]PowerForecastElement `json:"elements" description:"Elements of which this forecast consists. Contains at least one element. Elements must be placed in chronological order."`
	MessageID   *generated.ID           `json:"message_id" description:"ID of this message"`
	MessageType string                  `json:"message_type" description:"Type of this message" const:"PowerForecast"`
	StartTime   time.Time               `json:"start_time" description:"Start time of time period that is covered by the profile."`
}

// NewPowerForecast creates a new instance of PowerForecast and an error
func NewPowerForecast(elements *[]PowerForecastElement, startTime int64) (*PowerForecast, error) {
	if elements == nil || len(*elements) == 0 {
		return nil, errors.New("Elements must contain at least one element")
	}

	// Ensure elements are in chronological order and have valid durations
	for i := 0; i < len(*elements); i++ {
		if (*elements)[i].Duration == nil {
			return nil, errors.New("each element must have a duration")
		}
		if i > 0 && (*elements)[i].Duration.Milliseconds < (*elements)[i-1].Duration.Milliseconds {
			return nil, errors.New("elements must be in chronological order")
		}
	}

	id, err := generated.NewID()
	if err != nil {
		return nil, err
	}
	startTimeConverted := time.Unix(startTime, 0)
	return &PowerForecast{
		Elements:    elements,
		MessageID:   id,
		MessageType: "PowerForecast",
		StartTime:   startTimeConverted,
	}, nil
}
