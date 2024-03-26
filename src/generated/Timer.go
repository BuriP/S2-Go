package generated

import()

// Timer represents a timer entity.
type Timer struct {
	DiagnosticLabel *string  `json:"diagnostic_label,omitempty" description:"Human readable name/description of the Timer"` //pointer because its optional.
	Duration        Duration `json:"duration" description:"The time it takes for the Timer to finish after it has been started"`
	ID              ID       `json:"id" description:"ID of the Timer. Must be unique in the scope of the OMBC.SystemDescription, FRBC.ActuatorDescription or DDBC.ActuatorDescription in which it is used."`
}
