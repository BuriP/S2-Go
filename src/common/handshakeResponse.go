package common

import (
	"github.com/BuriP/S2-Go/src/generated"
)


// HandshakeResponse represents the handshake response message.
type HandshakeResponse struct {
    MessageID               *generated.ID     `json:"message_id" description:"ID of this message"`
    MessageType             string `json:"message_type" description:"HandshakeResponse" const:"true"`
    SelectedProtocolVersion *[]string `json:"selected_protocol_version" description:"The protocol version the CEM selected for this session"`
}

// NewHandshakeResponse creeates a new instance of the HandshakeResponse.
func NewHandshakeResponse(id *generated.ID, t string, handshake *Handshake) *HandshakeResponse {
	return &HandshakeResponse{ // returns a pointer with the new handshake response
		MessageID : id,
		MessageType: t,
		SelectedProtocolVersion: handshake.SupportedProtocolVersions,
	}
}

// GetSelectedProtocolVersionsResponse retuns a slice of the protocol versions defined in the HandshakeResponse
func (h *HandshakeResponse) GetSelectedProtocolVersionsResponse() *[]string {
	return h.SelectedProtocolVersion
}


