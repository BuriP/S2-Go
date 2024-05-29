package frbc

import (
	"errors"

	"github.com/BuriP/S2-Go/src/common"
)

// FRBCOperationModeElement represents an element of FRBC operation mode.
type FRBCOperationModeElement struct {
	FillLevelRange *common.NumberRange  `json:"fill_level_range" description:"The range of the fill level for which this FRBC.OperationModeElement applies. The start of the NumberRange shall be smaller than the end of the NumberRange."`
	FillRate       *common.NumberRange  `json:"fill_rate" description:"Indicates the change in fill_level per second. The lower_boundary of the NumberRange is associated with an operation_mode_factor of 0, the upper_boundary is associated with an operation_mode_factor of 1."`
	PowerRanges    []*common.PowerRange `json:"power_ranges" description:"The power produced or consumed by this operation mode. The start of each PowerRange is associated with an operation_mode_factor of 0, the end is associated with an operation_mode_factor of 1. In the array there must be at least one PowerRange, and at most one PowerRange per CommodityQuantity." maxItems:"10" minItems:"1"`
	RunningCosts   *common.NumberRange  `json:"running_costs,omitempty" description:"Additional costs per second (e.g. wear, services) associated with this operation mode in the currency defined by the ResourceManagerDetails, excluding the commodity cost. The range is expressing uncertainty and is not linked to the operation_mode_factor."`
}

// NewFRBCOperationModeElement creates a new instance of FRBCOperationModeElement if validation is correct, otherwise returns an error.
func NewFRBCOperationModeElement(fillLevelRange, fillRate *common.NumberRange, powerRanges []*common.PowerRange, runningCosts *common.NumberRange) (*FRBCOperationModeElement, error) {
	if fillLevelRange == nil {
		return nil, errors.New("fill level range must be provided")
	}
	if fillRate == nil {
		return nil, errors.New("fill rate must be provided")
	}
	if len(powerRanges) == 0 {
		return nil, errors.New("at least one power range must be provided")
	}

	if len(powerRanges) >= 10 {
		tenPowerRanges := powerRanges[:10]
		return &FRBCOperationModeElement{
			FillLevelRange: fillLevelRange,
			FillRate:       fillRate,
			PowerRanges:    tenPowerRanges,
			RunningCosts:   runningCosts,
		}, nil

	}

	return &FRBCOperationModeElement{
		FillLevelRange: fillLevelRange,
		FillRate:       fillRate,
		PowerRanges:    powerRanges,
		RunningCosts:   runningCosts,
	}, nil
}

// GetFillLevelRange returns the fill level range of the FRBCOperationModeElement.
func (e *FRBCOperationModeElement) GetFillLevelRange() *common.NumberRange {
	return e.FillLevelRange
}

// GetFillRate returns the fill rate of the FRBCOperationModeElement.
func (e *FRBCOperationModeElement) GetFillRate() *common.NumberRange {
	return e.FillRate
}

// GetPowerRanges returns the power ranges of the FRBCOperationModeElement.
func (e *FRBCOperationModeElement) GetPowerRanges() []*common.PowerRange {
	return e.PowerRanges
}

// GetRunningCosts returns the running costs of the FRBCOperationModeElement, if any.
func (e *FRBCOperationModeElement) GetRunningCosts() *common.NumberRange {
	return e.RunningCosts
}
