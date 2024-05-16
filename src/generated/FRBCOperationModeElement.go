// Look whether PowerRanges should be a slice or an array.

package generated

import(

)


// FRBCOperationModeElement represents an element of FRBC operation mode.
type FRBCOperationModeElement struct {
    FillLevelRange *NumberRange `json:"fill_level_range" description:"The range of the fill level for which this FRBC.OperationModeElement applies. The start of the NumberRange shall be smaller than the end of the NumberRange."`
    FillRate       *NumberRange `json:"fill_rate" description:"Indicates the change in fill_level per second. The lower_boundary of the NumberRange is associated with an operation_mode_factor of 0, the upper_boundary is associated with an operation_mode_factor of 1."`
    PowerRanges    []*PowerRange `json:"power_ranges" description:"The power produced or consumed by this operation mode. The start of each PowerRange is associated with an operation_mode_factor of 0, the end is associated with an operation_mode_factor of 1. In the array there must be at least one PowerRange, and at most one PowerRange per CommodityQuantity." maxItems:"10" minItems:"1"`// using slice here
    RunningCosts   *NumberRange `json:"running_costs,omitempty" description:"Additional costs per second (e.g. wear, services) associated with this operation mode in the currency defined by the ResourceManagerDetails, excluding the commodity cost. The range is expressing uncertainty and is not linked to the operation_mode_factor."`
}
