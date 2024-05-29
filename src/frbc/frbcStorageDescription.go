package frbc

import (
	"errors"
	"github.com/BuriP/S2-Go/src/common"
)

// FRBCStorageDescription represents the description of FRBC storage.
type FRBCStorageDescription struct {
	DiagnosticLabel                *string             `json:"diagnostic_label,omitempty" description:"Human readable name/description of the storage (e.g. hot water buffer or battery). This element is only intended for diagnostic purposes and not for HMI applications."`                                                      //pointer because its optional.
	FillLevelLabel                 *string             `json:"fill_level_label,omitempty" description:"Human readable description of the (physical) units associated with the fill_level (e.g. degrees Celsius or percentage state of charge). This element is only intended for diagnostic purposes and not for HMI applications."` //pointer because its optional.
	FillLevelRange                 *common.NumberRange `json:"fill_level_range" description:"The range in which the fill_level should remain. It is expected of the CEM to keep the fill_level within this range. When the fill_level is not within this range, the Resource Manager can ignore instructions from the CEM (except during abnormal conditions)."`
	ProvidesFillLevelTargetProfile bool                `json:"provides_fill_level_target_profile" description:"Indicates whether the Storage could provide a target profile for the fill level through the FRBC.FillLevelTargetProfile."`
	ProvidesLeakageBehaviour       bool                `json:"provides_leakage_behaviour" description:"Indicates whether the Storage could provide details of power leakage behaviour through the FRBC.LeakageBehaviour."`
	ProvidesUsageForecast          bool                `json:"provides_usage_forecast" description:"Indicates whether the Storage could provide a UsageForecast through the FRBC.UsageForecast."`
}

// NewFRBCStorageDescription creates a new instance of FRBCStorageDescription.
func NewFRBCStorageDescription(fillLevelRange *common.NumberRange, providesFillLevelTargetProfile, providesLeakageBehaviour, providesUsageForecast bool, diagnosticLabel, fillLevelLabel *string) (*FRBCStorageDescription, error) {
	if fillLevelRange == nil {
		return nil, errors.New("fill_level_range is required")
	}

	return &FRBCStorageDescription{
		DiagnosticLabel:                diagnosticLabel,
		FillLevelLabel:                 fillLevelLabel,
		FillLevelRange:                 fillLevelRange,
		ProvidesFillLevelTargetProfile: providesFillLevelTargetProfile,
		ProvidesLeakageBehaviour:       providesLeakageBehaviour,
		ProvidesUsageForecast:          providesUsageForecast,
	}, nil
}

// GetDiagnosticLabel returns the diagnostic label of the storage description.
func (sd *FRBCStorageDescription) GetDiagnosticLabel() *string {
	return sd.DiagnosticLabel
}

// GetFillLevelLabel returns the fill level label of the storage description.
func (sd *FRBCStorageDescription) GetFillLevelLabel() *string {
	return sd.FillLevelLabel
}

// GetFillLevelRange returns the fill level range of the storage description.
func (sd *FRBCStorageDescription) GetFillLevelRange() *common.NumberRange {
	return sd.FillLevelRange
}

// GetProvidesFillLevelTargetProfile returns whether the storage provides a fill level target profile.
func (sd *FRBCStorageDescription) GetProvidesFillLevelTargetProfile() bool {
	return sd.ProvidesFillLevelTargetProfile
}

// GetProvidesLeakageBehaviour returns whether the storage provides details of power leakage behaviour.
func (sd *FRBCStorageDescription) GetProvidesLeakageBehaviour() bool {
	return sd.ProvidesLeakageBehaviour
}

// GetProvidesUsageForecast returns whether the storage provides a usage forecast.
func (sd *FRBCStorageDescription) GetProvidesUsageForecast() bool {
	return sd.ProvidesUsageForecast
}
