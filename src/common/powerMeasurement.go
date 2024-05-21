package common

import (
	"github.com/BuriP/S2-Go/src/generated"
	"time"
)

type PowerMeasurement struct {
	MeasurementTimestamp time.Time               `json:"measurement_timestamp" description:"Timestamp when PowerValues were measured."`
	MessageID            *generated.ID           `json:"message_id" description:"ID of this message"`
	MessageType          string                  `json:"message_type" description:"Type of the message"`
	Values               *[]generated.PowerValue `json:"values" description:"Array of measured PowerValues"` // USed a slice here With type PowerValue
}

// NewPowerMeasurement creates anew instance of PowerMeasurement. It has unique ID
func NewPowerMeasurement(messageid string, values *[]generated.PowerValue, messagetype string) (*PowerMeasurement, error) {
	id, err := generated.NewID()
	if err != nil {
		return nil, err
	}
	return &PowerMeasurement{
		MeasurementTimestamp: time.Now(),
		MessageID:            id,
		MessageType:          messagetype,
		Values:               values,
	}, nil
}
