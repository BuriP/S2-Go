package frbc

import (
	"errors"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCActuatorDescriptor represents the description of an FRBC actuator.
type FRBCActuatorDescriptor struct {
	DiagnosticLabel      *string               `json:"diagnostic_label,omitempty" description:"Human readable name/description for the actuator. This element is only intended for diagnostic purposes and not for HMI applications."`
	ID                   *generated.ID         `json:"id" description:"ID of the Actuator. Must be unique in the scope of the Resource Manager, for at least the duration of the session between Resource Manager and CEM."`
	OperationModes       []*FRBCOperationMode  `json:"operation_modes" description:"Provided FRBC.OperationModes associated with this actuator" min_items:"1" max_items:"100"`
	SupportedCommodities []generated.Commodity `json:"supported_commodities" description:"List of all supported Commodities." min_items:"1" max_items:"4"`
	Timers               []*common.Timer       `json:"timers,omitempty" description:"List of Timers associated with this actuator" max_items:"1000"`
	Transitions          []*common.Transition  `json:"transitions,omitempty" description:"Possible transitions between FRBC.OperationModes associated with this actuator." max_items:"1000"`
}

// NewFRBCActuatorDescriptor creates a new instance of FRBCActuatorDescriptor if validated, otherwise returns an error.
func NewFRBCActuatorDescriptor(id *generated.ID, operationModes []*FRBCOperationMode, supportedCommodities []generated.Commodity, diagnosticLabel *string, timers []*common.Timer, transitions []*common.Transition) (*FRBCActuatorDescriptor, error) {
	if id == nil {
		return nil, errors.New("id must be provided")
	}
	if len(operationModes) == 0 {
		return nil, errors.New("at least one operation mode must be provided")
	}
	if len(supportedCommodities) == 0 || len(supportedCommodities) > 4 {
		return nil, errors.New("at least one supported commodity must be provided and maximum 4")
	}

	return &FRBCActuatorDescriptor{
		DiagnosticLabel:      diagnosticLabel,
		ID:                   id,
		OperationModes:       operationModes,
		SupportedCommodities: supportedCommodities,
		Timers:               timers,
		Transitions:          transitions,
	}, nil
}

// GetDiagnosticLabel returns the diagnostic label of the actuator descriptor.
func (d *FRBCActuatorDescriptor) GetDiagnosticLabel() *string {
	return d.DiagnosticLabel
}

// GetID returns the ID of the actuator descriptor.
func (d *FRBCActuatorDescriptor) GetID() *generated.ID {
	return d.ID
}

// GetOperationModes returns the operation modes of the actuator descriptor.
func (d *FRBCActuatorDescriptor) GetOperationModes() []*FRBCOperationMode {
	return d.OperationModes
}

// GetSupportedCommodities returns the supported commodities of the actuator descriptor.
func (d *FRBCActuatorDescriptor) GetSupportedCommodities() []generated.Commodity {
	return d.SupportedCommodities
}

// GetTimers returns the timers of the actuator descriptor.
func (d *FRBCActuatorDescriptor) GetTimers() []*common.Timer {
	return d.Timers
}

// GetTransitions returns the transitions of the actuator descriptor.
func (d *FRBCActuatorDescriptor) GetTransitions() []*common.Transition {
	return d.Transitions
}
