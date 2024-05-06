package common

import (

	"github.com/BuriP/S2-Go/src/generated"
	"github.com/google/uuid"
)

// Handshake represents a handshake message.
type Handshake struct {
    MessageID                 *generated.ID `json:"message_id" description:"ID of this message"`
    MessageType               string        `json:"message_type" description:"Handshake" const:"true"`
    Role                      *generated.EnergyManagementRole `json:"role" description:"The role of the sender of this message"`
    SupportedProtocolVersions *[]string      `json:"supported_protocol_versions,omitempty" description:"Protocol versions supported by the sender of this message. This field is mandatory for the RM, but optional for the CEM." minItems:"1"` // POinter due to it being optional
}

// NewHandshake creates a new Handshake instance.
func NewHandshake(id *generated.ID, t string, r *generated.EnergyManagementRole, protocol *Handshake) *Handshake {
	return &Handshake{
		MessageID : id,
		MessageType: t,
		Role: r, // requires &generated.CEM for example
		SupportedProtocolVersions : protocol.SupportedProtocolVersions, 
	}
}

// NewVarHandshake creates new n number of instances of Handshakes.
func NewVarHandshake(ids []*generated.ID, types []string, roles []*generated.EnergyManagementRole, n uint64) []*Handshake {
	var handshakes []*Handshake
	for i := uint64(0); i < n; i++ {
		handshake:= &Handshake{
			MessageID : ids[i],
			MessageType: types[i],
			Role: roles[i],
		}
		handshakes = append(handshakes, handshake)
	}
	return handshakes // returns slice of handshake instances
}

// SetMessageID sets a new MessageId in the Handshake
func (h *Handshake) SetMessageID() *Handshake{
	// Generating new UUID
	newUUID := uuid.New().String()
  
	// Converting the UUID into an ID type
	id := generated.IDFromString(newUUID)

	// Asigning the new UUID to the MessageID filed
	 h.MessageID = id

	return h // Returns back the  modified Handshake instance
}

// SetSupportedProtocolVersions set the protocol versions in the Handshake
func (h *Handshake) SetSupportedProtocolVersions(s *[]string) *Handshake{
	h.SupportedProtocolVersions = s
	return h // Returns the modified handshake
}




