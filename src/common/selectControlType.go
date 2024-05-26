package common

import (
	"github.com/BuriP/S2-Go/src/generated"
)

type SelectControlType struct {
	ControlType generated.ControlType `json:"control_type" description:"The ControlType to activate"`
	MessageID   *generated.ID         `json:"message_id" description:"ID of this message"`
	MessageType string                `json:"message_type" description:"Type of this message"`
}

// NewSelectControlType creates a new instance of SelectControlType.
func NewSelectControlType(controlType generated.ControlType) (*SelectControlType, error) {
	id, err := generated.NewID()
	if err != nil {
		return nil, err
	}
	return &SelectControlType{
		ControlType: controlType,
		MessageID:   id,
		MessageType: "Select Control Type",
	}, nil
}
