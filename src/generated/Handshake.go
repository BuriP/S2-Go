// TO DO : Look if its has to be a slice or an array

package generated

// Handshake represents a handshake message.
type Handshake struct {
    MessageID                 ID            `json:"message_id" description:"ID of this message"`
    MessageType               string        `json:"message_type" description:"Handshake" const:"true"`
    Role                      EnergyManagementRole `json:"role" description:"The role of the sender of this message"`
    SupportedProtocolVersions *[]string      `json:"supported_protocol_versions,omitempty" description:"Protocol versions supported by the sender of this message. This field is mandatory for the RM, but optional for the CEM." minItems:"1"` // POinter due to it being optional
}
