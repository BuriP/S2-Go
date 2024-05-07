package common

import(
	"github.com/BuriP/S2-Go/src/generated"
	"github.com/google/uuid"
)

type SelectControlType struct {
	ControlType generated.ControlType `json:"control_type" description:"The ControlType to activate"`
	MessageID   *generated.ID          `json:"message_id" description:"ID of this message"`
	MessageType string      `json:"message_type" description:"Type of this message"`
}

// NewSelectControlType creates a new instance of SelectControlType.
func NewSelectControlType(controlType generated.ControlType, messageID uuid.UUID, messageType string) (*SelectControlType, error) {
	id, err := generated.NewID(messageID.String())
	if err != nil {
		return nil, err
	}
	return &SelectControlType{
		ControlType: controlType,
		MessageID:   id,
		MessageType: messageType,
	}, nil
}