package generated

import(
	
)

// HandshakeResponse represents the handshake response message.
type HandshakeResponse struct {
    MessageID               ID     `json:"message_id" description:"ID of this message"`
    MessageType             string `json:"message_type" description:"HandshakeResponse" const:"true"`
    SelectedProtocolVersion string `json:"selected_protocol_version" description:"The protocol version the CEM selected for this session"`
}
