package frbc

import (
	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

type FRBCActuatorDescription struct {
	DiagnosticLabel      *string               `json:"diagnostic_label,omitempty" description:"Human readable name/description for the actuator. This element is only intended for diagnostic purposes and not for HMI applications."`
	ID                   *generated.ID         `json:"id" description:"ID of the Actuator. Must be unique in the scope of the Resource Manager, for at least the duration of the session between Resource Manager and CEM."`
	OperationModes       []*FRBCOperationMode  `json:"operation_modes" description:"Provided FRBC.OperationModes associated with this actuator" min_items:"1" max_items:"100"`
	SupportedCommodities []generated.Commodity `json:"supported_commodities" description:"List of all supported Commodities." min_items:"1" max_items:"4"`
	Timers               []common.Timer        `json:"timers,omitempty" description:"List of Timers associated with this actuator" max_items:"1000"`
	Transitions          []common.Transition   `json:"transitions,omitempty" description:"Possible transitions between FRBC.OperationModes associated with this actuator." max_items:"1000"`
}
