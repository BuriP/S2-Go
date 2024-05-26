package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

type SessionRequest struct {
	DiagnosticLabel *string                      `json:"diagnostic_label,omitempty" description:"Optional field for a human readable description for debugging purposes"`
	MessageID       *generated.ID                `json:"message_id" description:"ID of this message"`
	MessageType     string                       `json:"message_type" const:"true"`
	Request         generated.SessionRequestType `json:"request" description:"The type of request"`
}

// NewSessionRequest creates a new instance of SessionRequest.
func NewSessionRequest(request generated.SessionRequestType, diagnosticLabel ...string) (*SessionRequest, error) {
	if request == "" {
		return nil, errors.New("request is required")
	}

	id, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	var label *string
	if len(diagnosticLabel) > 0 {
		label = &diagnosticLabel[0]
	}

	return &SessionRequest{
		DiagnosticLabel: label,
		MessageID:       id,
		MessageType:     "Session Request",
		Request:         request,
	}, nil
}
