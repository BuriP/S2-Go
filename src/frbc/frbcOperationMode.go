package frbc

import (
	"errors"

	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCOperationMode represents an FRBC operation mode.
type FRBCOperationMode struct {
	AbnormalConditionOnly bool                        `json:"abnormal_condition_only" description:"Indicates if this FRBC.OperationMode may only be used during an abnormal condition"`
	DiagnosticLabel       *string                     `json:"diagnostic_label,omitempty" description:"Human readable name/description of the FRBC.OperationMode. This element is only intended for diagnostic purposes and not for HMI applications."`
	Elements              []*FRBCOperationModeElement `json:"elements" description:"List of FRBC.OperationModeElements, which describe the properties of this FRBC.OperationMode depending on the fill_level. The fill_level_ranges of the items in the Array must be contiguous."`
	ID                    *generated.ID               `json:"id" description:"ID of the FRBC.OperationMode. Must be unique in the scope of the FRBC.ActuatorDescription in which it is used."`
}

// NewFRBCOperationMode creates a new instance of FRBCOperationMode if validated, otherwise returns an error.
func NewFRBCOperationMode(abnormalConditionOnly bool, elements []*FRBCOperationModeElement, label *string) (*FRBCOperationMode, error) {
	if elements == nil || len(elements) == 0 {
		return nil, errors.New("elements must be provided")
	}

	// Ensure elements are in chronological order
	for i := 1; i < len(elements); i++ {
		if elements[i].FillLevelRange.StartOfRange != elements[i-1].FillLevelRange.EndOfRange {
			return nil, errors.New("elements must be in contiguous order")
		}
	}

	id, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &FRBCOperationMode{
		AbnormalConditionOnly: abnormalConditionOnly,
		DiagnosticLabel:       label,
		Elements:              elements,
		ID:                    id,
	}, nil
}

// GetAbnormalConditionOnly returns whether the operation mode is for abnormal conditions only.
func (m *FRBCOperationMode) GetAbnormalConditionOnly() bool {
	return m.AbnormalConditionOnly
}

// GetDiagnosticLabel returns the diagnostic label of the operation mode.
func (m *FRBCOperationMode) GetDiagnosticLabel() *string {
	return m.DiagnosticLabel
}

// GetElements returns the elements of the operation mode.
func (m *FRBCOperationMode) GetElements() []*FRBCOperationModeElement {
	return m.Elements
}

// GetID returns the ID of the operation mode.
func (m *FRBCOperationMode) GetID() *generated.ID {
	return m.ID
}
