package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

// Timer represents a timer with a duration and optional diagnostic label.
type Timer struct {
	DiagnosticLabel *string       `json:"diagnostic_label,omitempty" description:"Human readable name/description of the Timer"`
	Duration        *Duration     `json:"duration" description:"The time it takes for the Timer to finish after it has been started"`
	ID              *generated.ID `json:"id" description:"ID of the Timer. Must be unique in the scope of the OMBC.SystemDescription, FRBC.ActuatorDescription or DDBC.ActuatorDescription in which it is used."`
}

// NewTimer creates a new Timer instance.
func NewTimer(duration *Duration, diagnosticLabel ...string) (*Timer, error) {
	if duration == nil {
		return nil, errors.New("duration is required")
	}

	id, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	var label *string
	if len(diagnosticLabel) > 0 {
		label = &diagnosticLabel[0]
	}

	return &Timer{
		ID:              id,
		Duration:        duration,
		DiagnosticLabel: label,
	}, nil
}
