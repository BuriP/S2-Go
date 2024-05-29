package frbc

import (
	"errors"
	"time"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCTimerStatus represents the status of FRBC timer.
type FRBCTimerStatus struct {
	ActuatorID  *generated.ID `json:"actuator_id" description:"The ID of the actuator the timer belongs to"`
	FinishedAt  time.Time     `json:"finished_at" description:"Indicates when the Timer will be finished. If the DateTimeStamp is in the future, the timer is not yet finished. If the DateTimeStamp is in the past, the timer is finished. If the timer was never started, the value can be an arbitrary DateTimeStamp in the past."`
	MessageID   *generated.ID `json:"message_id" description:"ID of this message"`
	MessageType string        `json:"message_type" description:"FRBC.TimerStatus"`
	TimerID     *generated.ID `json:"timer_id" description:"The ID of the timer this message refers to"`
}

// NewFRBCTimerStatus creates a new instance of FRBCTimerStatus.
func NewFRBCTimerStatus(actuator *generated.ID, timer *common.Timer, finishedAt time.Time) (*FRBCTimerStatus, error) {
	if actuator == nil || actuator.Value == "" {
		return nil, errors.New("actuator and actuator ID are required")
	}
	if timer == nil || timer.ID == nil {
		return nil, errors.New("timer and timer ID are required")
	}

	messageID, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &FRBCTimerStatus{
		ActuatorID:  actuator,
		FinishedAt:  finishedAt,
		MessageID:   messageID,
		MessageType: "FRBC.TimerStatus",
		TimerID:     timer.ID,
	}, nil
}
