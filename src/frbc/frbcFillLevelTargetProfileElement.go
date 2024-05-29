package frbc

import (
	"errors"

	"github.com/BuriP/S2-Go/src/common"
)

// FRBCFillLevelTargetProfileElement represents an element of FRBC fill level target profile.
type FRBCFillLevelTargetProfileElement struct {
	Duration       *common.Duration    `json:"duration" description:"The duration of the element"`
	FillLevelRange *common.NumberRange `json:"fill_level_range" description:"The target range in which the fill_level must be for the time period during which the element is active"`
}

// NewFRBCFillLevelTargetProfileElement creates a new instance of FRBCFillLevelTargetProfileElement if validated, otherwise returns an error.
func NewFRBCFillLevelTargetProfileElement(duration *common.Duration, fillLevelRange *common.NumberRange) (*FRBCFillLevelTargetProfileElement, error) {
	if duration == nil {
		return nil, errors.New("duration must be provided")
	}
	if fillLevelRange == nil {
		return nil, errors.New("fill level range must be provided")
	}

	return &FRBCFillLevelTargetProfileElement{
		Duration:       duration,
		FillLevelRange: fillLevelRange,
	}, nil
}

// GetDuration returns the duration of the FRBCFillLevelTargetProfileElement.
func (e *FRBCFillLevelTargetProfileElement) GetDuration() *common.Duration {
	return e.Duration
}

// GetFillLevelRange returns the fill level range of the FRBCFillLevelTargetProfileElement.
func (e *FRBCFillLevelTargetProfileElement) GetFillLevelRange() *common.NumberRange {
	return e.FillLevelRange
}
