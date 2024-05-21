package common

import (
	"github.com/BuriP/S2-Go/src/generated"
	"time"
)

type PowerForecast struct {
	Elements    *[]generated.PowerForecastElement `json:"elements" description:"Elements of which this forecast consists. Contains at least one element. Elements must be placed in chronological order."`
	MessageID   *generated.ID                     `json:"message_id" description:"ID of this message"`
	MessageType string                            `json:"message_type" description:"Type of this message" const:"PowerForecast"`
	StartTime   time.Time                         `json:"start_time" description:"Start time of time period that is covered by the profile."`
}

// NewPowerForecast creates a new instance of PowerForecast and an error
func NewPowerForecast(elements *[]generated.PowerForecastElement, message string, messagetype string, starttime int64) (*PowerForecast, error) {
	id, err := generated.NewID()
	if err != nil {
		return nil, err
	}
	startTimeConverted := time.Unix(starttime, 0)
	return &PowerForecast{
		Elements:    elements,
		MessageID:   id,
		MessageType: messagetype,
		StartTime:   startTimeConverted,
	}, nil
}

