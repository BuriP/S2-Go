package frbc

import (
	"github.com/BuriP/S2-Go/src/common"

	"github.com/BuriP/S2-Go/src/generated"
)

type FRBCOperationMode struct {
	AbnormalConditionOnly bool                       `json:"abnormal_condition_only" description:"Indicates if this FRBC.OperationMode may only be used during an abnormal condition"`
	DiagnosticLabel       *string                    `json:"diagnostic_label,omitempty" description:"Human readable name/description of the FRBC.OperationMode. This element is only intended for diagnostic purposes and not for HMI applications."`
	Elements              []FRBCOperationModeElement `json:"elements" description:"List of FRBC.OperationModeElements, which describe the properties of this FRBC.OperationMode depending on the fill_level. The fill_level_ranges of the items in the Array must be contiguous."`
	ID                    *generated.ID              `json:"id" description:"ID of the FRBC.OperationMode. Must be unique in the scope of the FRBC.ActuatorDescription in which it is used."`
}

type FRBCActuatorDescription struct {
	DiagnosticLabel      *string             `json:"diagnostic_label,omitempty" description:"Human readable name/description for the actuator. This element is only intended for diagnostic purposes and not for HMI applications."`
	ID                   ID                  `json:"id" description:"ID of the Actuator. Must be unique in the scope of the Resource Manager, for at least the duration of the session between Resource Manager and CEM."`
	OperationModes       []FRBCOperationMode `json:"operation_modes" description:"Provided FRBC.OperationModes associated with this actuator" min_items:"1" max_items:"100"`
	SupportedCommodities []Commodity         `json:"supported_commodities" description:"List of all supported Commodities." min_items:"1" max_items:"4"`
	Timers               []Timer             `json:"timers,omitempty" description:"List of Timers associated with this actuator" max_items:"1000"`
	Transitions          []Transition        `json:"transitions,omitempty" description:"Possible transitions between FRBC.OperationModes associated with this actuator." max_items:"1000"`
}
