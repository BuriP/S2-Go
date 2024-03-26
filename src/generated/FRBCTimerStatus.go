package generated

import (
	"time"
)

// FRBCTimerStatus represents the status of FRBC timer.
type FRBCTimerStatus struct {
    ActuatorID   ID       `json:"actuator_id" description:"The ID of the actuator the timer belongs to"`
    FinishedAt   time.Time `json:"finished_at" description:"Indicates when the Timer will be finished. If the DateTimeStamp is in the future, the timer is not yet finished. If the DateTimeStamp is in the past, the timer is finished. If the timer was never started, the value can be an arbitrary DateTimeStamp in the past."`
    MessageID    ID       `json:"message_id" description:"ID of this message"`
    MessageType  string   `json:"message_type" description:"FRBC.TimerStatus"`
    TimerID      ID       `json:"timer_id" description:"The ID of the timer this message refers to"`
}
