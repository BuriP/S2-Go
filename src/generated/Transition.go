package generated

// Transition represents a transition entity. We use slices here to vary the size of the possible elements. 

// Maybe look into converting into an array instead of a slice.

type Transition struct {
	AbnormalConditionOnly bool        `json:"abnormal_condition_only" description:"Indicates if this Transition may only be used during an abnormal condition. Have a look at CLAUS"`
	BlockingTimers        []ID        `json:"blocking_timers" description:"List of IDs of Timers that block this Transition from initiating while at least one of these Timers is not yet finished" maxItems:"1000" minItems:"0"`
	From                  ID          `json:"from" alias:"from_" description:"ID of the OperationMode (exact type differs per ControlType) that should be switched from."`
	ID                    ID          `json:"id" description:"ID of the Transition. Must be unique in the scope of the OMBC.SystemDescription, FRBC.ActuatorDescription or DDBC.ActuatorDescription in which it is used."`
	StartTimers           []ID        `json:"start_timers" description:"List of IDs of Timers that will be (re)started when this transition is initiated" maxItems:"1000" minItems:"0"`
	To                    ID          `json:"to" description:"ID of the OperationMode (exact type differs per ControlType) that will be switched to."`
	TransitionCosts       *float64    `json:"transition_costs,omitempty" description:"Absolute costs for going through this Transition in the currency as described in the ResourceManagerDetails."` //pointer because its optional
	TransitionDuration    *Duration   `json:"transition_duration,omitempty" description:"Indicates the time between the initiation of this Transition, and the time at which the device behaves according to the Operation Mode which is defined in the 'to' data element. When no value is provided it is assumed the transition duration is negligible."` //pointer because its optional
}
