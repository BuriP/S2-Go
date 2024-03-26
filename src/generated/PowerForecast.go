package generated

import(
	"time"
)

type PowerForecast struct {
	Elements    []PowerForecastElement `json:"elements" description:"Elements of which this forecast consists. Contains at least one element. Elements must be placed in chronological order."`
	MessageID   ID                     `json:"message_id" description:"ID of this message"`
	MessageType string                 `json:"message_type" description:"Type of this message" const:"PowerForecast"`
	StartTime   time.Time              `json:"start_time" description:"Start time of time period that is covered by the profile."`
}
