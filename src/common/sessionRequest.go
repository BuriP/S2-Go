package common

import (
	"github.com/BuriP/S2-Go/src/generated"
)

type SessionRequest struct {
	DiagnosticLabel *string                       `json:"diagnostic_label,omitempty" description:"Optional field for a human readable description for debugging purposes"`
	MessageID       *generated.ID                 `json:"message_id" description:"ID of this message"`
	MessageType     string                        `json:"message_type" const:"true"`
	Request         *generated.SessionRequestType `json:"request" description:"The type of request"`
}

// NewSessionRequest creates a new instance of SessionRequest.
func NewSessionRequest(diagnosticLabel *string, request *generated.SessionRequestType) (*SessionRequest, error) { //is it best to handle the error inside?
	id, err := generated.NewID()
	if err != nil {
		return nil, err
	}
	return &SessionRequest{
		DiagnosticLabel: diagnosticLabel,
		MessageID:       id,
		MessageType:     "Session Request.",
		Request:         request,
	}, nil
}

