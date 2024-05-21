package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

// Role constants
var (
	CEM = generated.EnergyManagementRole("CEM")
	RM  = generated.EnergyManagementRole("RM")
)

// Handshake represents a handshake message.
type Handshake struct {
	MessageID                 *generated.ID                   `json:"message_id" description:"ID of this message"`
	MessageType               string                          `json:"message_type" description:"Handshake" const:"true"`
	Role                      *generated.EnergyManagementRole `json:"role" description:"The role of the sender of this message"`
	SupportedProtocolVersions *[]string                       `json:"supported_protocol_versions,omitempty" description:"Protocol versions supported by the sender of this message. This field is mandatory for the RM, but optional for the CEM." minItems:"1"` // Pointer due to it being optional
}

// NewHandshake creates a new Handshake instance.
func NewHandshake(role string, supportedProtocols []string) (*Handshake, error) {
	// Transform the role string to the corresponding EnergyManagementRole constant
	var r *generated.EnergyManagementRole
	switch role {
	case "CEM":
		r = &CEM
	case "RM":
		r = &RM
		if len(supportedProtocols) < 1 {
			return nil, errors.New("supportedProtocolVersions is mandatory for the RM")
		}
	default:
		return nil, errors.New("invalid role: must be CEM or RM")
	}

	// Generate a new ID
	id, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &Handshake{
		MessageID:                 id,
		MessageType:               "Handshake",
		Role:                      r,
		SupportedProtocolVersions: &supportedProtocols,
	}, nil
}

// NewVarHandshake creates n number of instances of Handshakes.
func NewVarHandshake(roles []string, supportedProtocols [][]string) ([]*Handshake, error) {
	var handshakes []*Handshake
	for i := 0; i < len(roles); i++ {
		handshake, err := NewHandshake(roles[i], supportedProtocols[i])
		if err != nil {
			return nil, err
		}
		handshakes = append(handshakes, handshake)
	}
	return handshakes, nil
}

// SetMessageID sets a new MessageId based on the pattern string given in the Handshake
func (h *Handshake) SetMessageID(pattern string) *Handshake {
	validID := generated.TransformToValidID(pattern)
	newID := &generated.ID{Value: validID}
	h.MessageID = newID
	return h
}

// SetSupportedProtocolVersions sets the protocol versions in the Handshake
func (h *Handshake) SetSupportedProtocolVersions(s []string) *Handshake {
	h.SupportedProtocolVersions = &s
	return h // Returns the modified handshake
}
