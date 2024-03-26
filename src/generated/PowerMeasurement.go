package generated

import (
	"time"
)

// PowerMeasurement represents a measurement of power values.
type PowerMeasurement struct {
    MeasurementTimestamp time.Time  `json:"measurement_timestamp" description:"Timestamp when PowerValues were measured."`
    MessageID            ID         `json:"message_id" description:"ID of this message"`
    MessageType          string     `json:"message_type" description:"Type of the message"`
    Values               []PowerValue `json:"values" description:"Array of measured PowerValues"` // USed a slice here With type PowerValue
}
