// TO DO: check whether it has to be a n array or a slice. Notice max elemnts is 288 and min is 1.

package generated

import(
	"time"
)

// FRBCFillLevelTargetProfile represents a target profile for fill levels.
type FRBCFillLevelTargetProfile struct {
    Elements    []FRBCFillLevelTargetProfileElement `json:"elements" description:"List of different fill levels that have to be targeted within a given duration. There shall be at least one element. Elements must be placed in chronological order."` // Maybe it should be an array because max elements is 288 and min 1
    MessageID   ID                                   `json:"message_id" description:"ID of this message"`
    MessageType string                               `json:"message_type" const:"true"`
    StartTime   time.Time                             `json:"start_time" description:"Time at which the FRBC.FillLevelTargetProfile starts."`
}
