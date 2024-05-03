package generated

import(

)

// Duration represents a duration in milliseconds. 
type Duration struct {
    Milliseconds float64 `json:"milliseconds" description:"Duration in milliseconds"` // Used uint64 to delimit minimum at 0
}
