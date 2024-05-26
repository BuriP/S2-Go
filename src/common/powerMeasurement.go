package common

import (
	"errors"
	"time"

	"github.com/BuriP/S2-Go/src/generated"
)

type PowerMeasurement struct {
	MeasurementTimestamp time.Time     `json:"measurement_timestamp" description:"Timestamp when PowerValues were measured."`
	MessageID            *generated.ID `json:"message_id" description:"ID of this message"`
	MessageType          string        `json:"message_type" description:"Type of the message"`
	Values               *[]PowerValue `json:"values" description:"Array of measured PowerValues"`
}

// NewPowerMeasurement creates a new instance of PowerMeasurement and returns an error if any validation fai
func NewPowerMeasurement(values *[]PowerValue) (*PowerMeasurement, error) {
	if values == nil || len(*values) == 0 {
		return nil, errors.New("values must contain at least one element")
	}

	id, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &PowerMeasurement{
		MeasurementTimestamp: time.Now(),
		MessageID:            id,
		MessageType:          "PowerMeasurement",
		Values:               values,
	}, nil
}
