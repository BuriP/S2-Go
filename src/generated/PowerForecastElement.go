package generated

// PowerForecastElement represents an element of a power forecast.
type PowerForecastElement struct {
    Duration     *Duration              `json:"duration" description:"Duration of the PowerForecastElement"`
    PowerValues  *[]PowerForecastValue `json:"power_values" description:"The values of power that are expected for the given period of time."` // USed a slice here of type PowerForecastValue.
}
