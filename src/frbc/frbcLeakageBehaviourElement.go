package frbc

import (
	"errors"

	"github.com/BuriP/S2-Go/src/common"
)

// FRBCLeakageBehaviourElement represents an element of FRBC leakage behavior.
type FRBCLeakageBehaviourElement struct {
	FillLevelRange *common.NumberRange `json:"fill_level_range" description:"The fill level range for which this FRBC.LeakageBehaviourElement applies. The start of the range must be less than the end of the range"`
	LeakageRate    float64             `json:"leakage_rate" description:"Indicates how fast the momentary fill level will decrease per second due to leakage within the given range of the fill level"`
}

// NewFRBCLeakageBehaviourElement creates a new instance of FRBCLeakageBehaviourElement if validation is correct, otherwise returns an error.
func NewFRBCLeakageBehaviourElement(fillLevelRange *common.NumberRange, leakageRate float64) (*FRBCLeakageBehaviourElement, error) {
	if fillLevelRange == nil {
		return nil, errors.New("fill level range must be provided")
	}

	if leakageRate == 0 {
		return nil, errors.New("leakage rate should not be 0")
	}

	return &FRBCLeakageBehaviourElement{
		FillLevelRange: fillLevelRange,
		LeakageRate:    leakageRate,
	}, nil
}

// GetFillLevelRange returns the fill level range of the FRBCLeakageBehaviourElement.
func (e *FRBCLeakageBehaviourElement) GetFillLevelRange() *common.NumberRange {
	return e.FillLevelRange
}

// GetLeakageRate returns the leakage rate of the FRBCLeakageBehaviourElement.
func (e *FRBCLeakageBehaviourElement) GetLeakageRate() float64 {
	return e.LeakageRate
}
