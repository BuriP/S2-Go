package generated

import(

)

// SelectControlType represents a message to select a control type.

type SelectControlType struct {
	ControlType ControlType `json:"control_type" description:"The ControlType to activate"`
	MessageID   ID          `json:"message_id" description:"ID of this message"`
	MessageType string      `json:"message_type" description:"Type of this message"`
}

