package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

// HandshakeResponse represents the handshake response message.
type HandshakeResponse struct {
	MessageID               *generated.ID `json:"message_id" description:"ID of this message"`
	MessageType             string        `json:"message_type" description:"HandshakeResponse" const:"true"`
	SelectedProtocolVersion *[]string     `json:"selected_protocol_version" description:"The protocol version the CEM selected for this session"`
}

// NewHandshakeResponse creates a new instance of the HandshakeResponse.
func NewHandshakeResponse(handshake *Handshake) (*HandshakeResponse, error) {
	if handshake == nil {
		return nil, errors.New("handshake is required")
	}
	if handshake.SupportedProtocolVersions == nil {
		return nil, errors.New("supported protocol versions are required in the handshake")
	}

	newID, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &HandshakeResponse{
		MessageID:               newID,
		MessageType:             "HandshakeResponse",
		SelectedProtocolVersion: handshake.SupportedProtocolVersions,
	}, nil
}

// GetSelectedProtocolVersions returns a slice of the protocol versions defined in the HandshakeResponse.
func (h *HandshakeResponse) GetSelectedProtocolVersions() *[]string {
	return h.SelectedProtocolVersion
}
