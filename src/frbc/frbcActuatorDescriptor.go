package frbc

import (
	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
	"github.com/google/uuid"
)

type FRBCActuatorDescription struct {
	DiagnosticLabel      *string               `json:"diagnostic_label,omitempty" description:"Human readable name/description for the actuator. This element is only intended for diagnostic purposes and not for HMI applications."`
	ID                   *generated.ID         `json:"id" description:"ID of the Actuator. Must be unique in the scope of the Resource Manager, for at least the duration of the session between Resource Manager and CEM."`
	OperationModes       []*FRBCOperationMode  `json:"operation_modes" description:"Provided FRBC.OperationModes associated with this actuator" min_items:"1" max_items:"100"`
	SupportedCommodities []generated.Commodity `json:"supported_commodities" description:"List of all supported Commodities." min_items:"1" max_items:"4"`
	Timers               []*common.Timer       `json:"timers,omitempty" description:"List of Timers associated with this actuator" max_items:"1000"`
	Transitions          []*common.Transition  `json:"transitions,omitempty" description:"Possible transitions between FRBC.OperationModes associated with this actuator." max_items:"1000"`
}

type FRBCOperationMode struct {
	AbnormalConditionOnly bool                        `json:"abnormal_condition_only" description:"Indicates if this FRBC.OperationMode may only be used during an abnormal condition"`
	DiagnosticLabel       *string                     `json:"diagnostic_label,omitempty" description:"Human readable name/description of the FRBC.OperationMode. This element is only intended for diagnostic purposes and not for HMI applications."`
	Elements              []*FRBCOperationModeElement `json:"elements" description:"List of FRBC.OperationModeElements, which describe the properties of this FRBC.OperationMode depending on the fill_level. The fill_level_ranges of the items in the Array must be contiguous."`
	ID                    *generated.ID               `json:"id" description:"ID of the FRBC.OperationMode. Must be unique in the scope of the FRBC.ActuatorDescription in which it is used."`
}

type FRBCOperationModeElement struct {
	FillLevelRange *common.NumberRange  `json:"fill_level_range" description:"The range of the fill level for which this FRBC.OperationModeElement applies. The start of the NumberRange shall be smaller than the end of the NumberRange."`
	FillRate       *common.NumberRange  `json:"fill_rate" description:"Indicates the change in fill_level per second. The lower_boundary of the NumberRange is associated with an operation_mode_factor of 0, the upper_boundary is associated with an operation_mode_factor of 1."`
	PowerRanges    []*common.PowerRange `json:"power_ranges" description:"The power produced or consumed by this operation mode. The start of each PowerRange is associated with an operation_mode_factor of 0, the end is associated with an operation_mode_factor of 1. In the array there must be at least one PowerRange, and at most one PowerRange per CommodityQuantity." maxItems:"10" minItems:"1"` // using slice here
	RunningCosts   *common.NumberRange  `json:"running_costs,omitempty" description:"Additional costs per second (e.g. wear, services) associated with this operation mode in the currency defined by the ResourceManagerDetails, excluding the commodity cost. The range is expressing uncertainty and is not linked to the operation_mode_factor."`
}

// NewActuatorDescriptor creates a new instace of the FRBCActuatorDescriptor.
func NewActuatorDescriptor(label *string, modes []*FRBCOperationMode, commodities []generated.Commodity, timers []*common.Timer, transitions []*common.Transition) *FRBCActuatorDescription {
	id, _ := generated.NewID()
	return &FRBCActuatorDescription{
		DiagnosticLabel:      label,
		ID:                   id,
		OperationModes:       modes,
		SupportedCommodities: commodities,
		Timers:               timers,
		Transitions:          transitions,
	}
}
