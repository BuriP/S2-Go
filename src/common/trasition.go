package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

// Transition represents a transition between operation modes.
type Transition struct {
	AbnormalConditionOnly bool            `json:"abnormal_condition_only" description:"Indicates if this Transition may only be used during an abnormal condition."`
	BlockingTimers        []*generated.ID `json:"blocking_timers" description:"List of IDs of Timers that block this Transition from initiating while at least one of these Timers is not yet finished" maxItems:"1000" minItems:"0"`
	From                  *generated.ID   `json:"from" alias:"from_" description:"ID of the OperationMode (exact type differs per ControlType) that should be switched from."`
	ID                    *generated.ID   `json:"id" description:"ID of the Transition. Must be unique in the scope of the OMBC.SystemDescription, FRBC.ActuatorDescription or DDBC.ActuatorDescription in which it is used."`
	StartTimers           []*generated.ID `json:"start_timers" description:"List of IDs of Timers that will be (re)started when this transition is initiated" maxItems:"1000" minItems:"0"`
	To                    *generated.ID   `json:"to" description:"ID of the OperationMode (exact type differs per ControlType) that will be switched to."`
	TransitionCosts       *float64        `json:"transition_costs,omitempty" description:"Absolute costs for going through this Transition in the currency as described in the ResourceManagerDetails."`
	TransitionDuration    *Duration       `json:"transition_duration,omitempty" description:"Indicates the time between the initiation of this Transition, and the time at which the device behaves according to the Operation Mode which is defined in the 'to' data element."`
}

// NewTransition creates a new instance of Transition.
func NewTransition(id *generated.ID, from *generated.ID, to *generated.ID, startTimers []*generated.ID, blockingTimers []*generated.ID, abnormalConditionOnly bool, transitionCosts *float64, transitionDuration *Duration) (*Transition, error) {
	if id == nil || from == nil || to == nil {
		return nil, errors.New("id, from, and to are required")
	}

	return &Transition{
		ID:                    id,
		From:                  from,
		To:                    to,
		StartTimers:           startTimers,
		BlockingTimers:        blockingTimers,
		AbnormalConditionOnly: abnormalConditionOnly,
		TransitionCosts:       transitionCosts,
		TransitionDuration:    transitionDuration,
	}, nil
}

// uuidListToIDList converts a list of UUID strings to a list of generated.ID pointers.// Check whether this is still needed.
func uuidListToIDList(uuids []string) []*generated.ID {
	idList := make([]*generated.ID, len(uuids))
	for i, u := range uuids {
		idList[i] = &generated.ID{Value: u}
	}
	return idList
}
