package common

import (
	"github.com/BuriP/S2-Go/src/generated"

	"github.com/google/uuid"
)

type SessionRequest struct {
    DiagnosticLabel *string           `json:"diagnostic_label,omitempty" description:"Optional field for a human readable description for debugging purposes"`
    MessageID       *generated.ID                `json:"message_id" description:"ID of this message"`
    MessageType     string            `json:"message_type" const:"true"`
    Request         *generated.SessionRequestType `json:"request" description:"The type of request"`
}


// NewSessionRequest creates a new instance of SessionRequest.
func NewSessionRequest(diagnosticLabel *string, messageID uuid.UUID, messageType string, request *generated.SessionRequestType) *SessionRequest { //is it best to handle the error inside?
	id, _ := generated.NewID(messageID.String())
	return &SessionRequest{
		DiagnosticLabel: diagnosticLabel,
		MessageID:       id,
		MessageType:     messageType,
		Request:         request,
	}
}