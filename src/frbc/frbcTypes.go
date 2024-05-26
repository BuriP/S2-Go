package frbc

import (
	"github.com/BuriP/S2-Go/src/common"

	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCFillLevelTargetProfileElement represents an element of FRBC fill level target profile.
type FRBCFillLevelTargetProfileElement struct {
	Duration       *common.Duration    `json:"duration" description:"The duration of the element"`
	FillLevelRange *common.NumberRange `json:"fill_level_range" description:"The target range in which the fill_level must be for the time period during which the element is active"`
}

// FRBCUsageForecastElement represents an element of FRBC usage forecast.
type FRBCUsageForecastElement struct {
	Duration            *common.Duration `json:"duration" description:"Indicator for how long the given usage_rate is valid."`
	UsageRateExpected   float64          `json:"usage_rate_expected" description:"The most likely value for the usage rate; the expected increase or decrease of the fill_level per second"`
	UsageRateLower68PPR *float64         `json:"usage_rate_lower_68PPR,omitempty" description:"The lower limit of the range with a 68% probability that the usage rate is within that range"`   // Used as pointer as in case the are not included.
	UsageRateLower95PPR *float64         `json:"usage_rate_lower_95PPR,omitempty" description:"The lower limit of the range with a 95% probability that the usage rate is within that range"`   // Used as pointer as in case the are not included.
	UsageRateLowerLimit *float64         `json:"usage_rate_lower_limit,omitempty" description:"The lower limit of the range with a 100% probability that the usage rate is within that range"`  // Used as pointer as in case the are not included.
	UsageRateUpper68PPR *float64         `json:"usage_rate_upper_68PPR,omitempty" description:"The upper limit of the range with a 68% probability that the usage rate is within that range"`   // Used as pointer as in case the are not included.
	UsageRateUpper95PPR *float64         `json:"usage_rate_upper_95PPR,omitempty" description:"The upper limit of the range with a 95% probability that the usage rate is within that range."`  // Used as pointer as in case the are not included.
	UsageRateUpperLimit *float64         `json:"usage_rate_upper_limit,omitempty" description:"The upper limit of the range with a 100% probability that the usage rate is within that range."` // Used as pointer as in case the are not included.
}

// FRBCLeakageBehaviourElement represents an element of FRBC leakage behavior.
type FRBCLeakageBehaviourElement struct {
	FillLevelRange *common.NumberRange `json:"fill_level_range" description:"The fill level range for which this FRBC.LeakageBehaviourElement applies. The start of the range must be less than the end of the range"`
	LeakageRate    float64             `json:"leakage_rate" description:"Indicates how fast the momentary fill level will decrease per second due to leakage within the given range of the fill level"`
}

type FRBCOperationMode struct {
	AbnormalConditionOnly bool                       `json:"abnormal_condition_only" description:"Indicates if this FRBC.OperationMode may only be used during an abnormal condition"`
	DiagnosticLabel       *string                    `json:"diagnostic_label,omitempty" description:"Human readable name/description of the FRBC.OperationMode. This element is only intended for diagnostic purposes and not for HMI applications."`
	Elements              []FRBCOperationModeElement `json:"elements" description:"List of FRBC.OperationModeElements, which describe the properties of this FRBC.OperationMode depending on the fill_level. The fill_level_ranges of the items in the Array must be contiguous."`
	ID                    *generated.ID              `json:"id" description:"ID of the FRBC.OperationMode. Must be unique in the scope of the FRBC.ActuatorDescription in which it is used."`
}

// FRBCOperationModeElement represents an element of FRBC operation mode.
type FRBCOperationModeElement struct {
	FillLevelRange *NumberRange  `json:"fill_level_range" description:"The range of the fill level for which this FRBC.OperationModeElement applies. The start of the NumberRange shall be smaller than the end of the NumberRange."`
	FillRate       *NumberRange  `json:"fill_rate" description:"Indicates the change in fill_level per second. The lower_boundary of the NumberRange is associated with an operation_mode_factor of 0, the upper_boundary is associated with an operation_mode_factor of 1."`
	PowerRanges    []*PowerRange `json:"power_ranges" description:"The power produced or consumed by this operation mode. The start of each PowerRange is associated with an operation_mode_factor of 0, the end is associated with an operation_mode_factor of 1. In the array there must be at least one PowerRange, and at most one PowerRange per CommodityQuantity." maxItems:"10" minItems:"1"` // using slice here
	RunningCosts   *NumberRange  `json:"running_costs,omitempty" description:"Additional costs per second (e.g. wear, services) associated with this operation mode in the currency defined by the ResourceManagerDetails, excluding the commodity cost. The range is expressing uncertainty and is not linked to the operation_mode_factor."`
}

// FRBCStorageDescription represents the description of FRBC storage.
type FRBCStorageDescription struct {
	DiagnosticLabel                *string             `json:"diagnostic_label,omitempty" description:"Human readable name/description of the storage (e.g. hot water buffer or battery). This element is only intended for diagnostic purposes and not for HMI applications."`                                                      //pointer because its optional.
	FillLevelLabel                 *string             `json:"fill_level_label,omitempty" description:"Human readable description of the (physical) units associated with the fill_level (e.g. degrees Celsius or percentage state of charge). This element is only intended for diagnostic purposes and not for HMI applications."` //pointer because its optional.
	FillLevelRange                 *common.NumberRange `json:"fill_level_range" description:"The range in which the fill_level should remain. It is expected of the CEM to keep the fill_level within this range. When the fill_level is not within this range, the Resource Manager can ignore instructions from the CEM (except during abnormal conditions)."`
	ProvidesFillLevelTargetProfile bool                `json:"provides_fill_level_target_profile" description:"Indicates whether the Storage could provide a target profile for the fill level through the FRBC.FillLevelTargetProfile."`
	ProvidesLeakageBehaviour       bool                `json:"provides_leakage_behaviour" description:"Indicates whether the Storage could provide details of power leakage behaviour through the FRBC.LeakageBehaviour."`
	ProvidesUsageForecast          bool                `json:"provides_usage_forecast" description:"Indicates whether the Storage could provide a UsageForecast through the FRBC.UsageForecast."`
}

type FRBCActuatorDescription struct {
	DiagnosticLabel      *string             `json:"diagnostic_label,omitempty" description:"Human readable name/description for the actuator. This element is only intended for diagnostic purposes and not for HMI applications."`
	ID                   ID                  `json:"id" description:"ID of the Actuator. Must be unique in the scope of the Resource Manager, for at least the duration of the session between Resource Manager and CEM."`
	OperationModes       []FRBCOperationMode `json:"operation_modes" description:"Provided FRBC.OperationModes associated with this actuator" min_items:"1" max_items:"100"`
	SupportedCommodities []Commodity         `json:"supported_commodities" description:"List of all supported Commodities." min_items:"1" max_items:"4"`
	Timers               []Timer             `json:"timers,omitempty" description:"List of Timers associated with this actuator" max_items:"1000"`
	Transitions          []Transition        `json:"transitions,omitempty" description:"Possible transitions between FRBC.OperationModes associated with this actuator." max_items:"1000"`
}
