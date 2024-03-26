package generated

import (
	"time"
)

type FRBCLeakageBehaviour struct {
	Elements   []FRBCLeakageBehaviourElement `json:"elements" description:"List of elements that model the leakage behaviour of the buffer. The fill_level_ranges of the elements must be contiguous."`
	MessageID  ID                            `json:"message_id" description:"ID of this message"`
	MessageType string                        `json:"message_type" const:"true" description:"FRBC.LeakageBehaviour"`
	ValidFrom  time.Time                     `json:"valid_from" description:"Moment this FRBC.LeakageBehaviour starts to be valid. If the FRBC.LeakageBehaviour is immediately valid, the DateTimeStamp should be now or in the past."`
}
