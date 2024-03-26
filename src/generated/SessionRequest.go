package generated

// SessionRequest represents a request for a session.
type SessionRequest struct {
    DiagnosticLabel *string           `json:"diagnostic_label,omitempty" description:"Optional field for a human readable description for debugging purposes"`
    MessageID       ID                `json:"message_id" description:"ID of this message"`
    MessageType     string            `json:"message_type" const:"true"`
    Request         SessionRequestType `json:"request" description:"The type of request"`
}
