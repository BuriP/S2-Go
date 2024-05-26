package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

// Handshake represents a handshake message.
type Handshake struct {
	MessageID                 *generated.ID                  `json:"message_id" description:"ID of this message"`
	MessageType               string                         `json:"message_type" description:"Handshake" const:"true"`
	Role                      generated.EnergyManagementRole `json:"role" description:"The role of the sender of this message"`
	SupportedProtocolVersions *[]string                      `json:"supported_protocol_versions,omitempty" description:"Protocol versions supported by the sender of this message. This field is mandatory for the RM, but optional for the CEM." minItems:"1"`
}

// NewHandshake creates a new Handshake instance.
func NewHandshake(role generated.EnergyManagementRole, supportedProtocols []string) (*Handshake, error) {
	if role == generated.RM && len(supportedProtocols) < 1 {
		return nil, errors.New("supportedProtocolVersions is mandatory for the RM")
	}
	if role != generated.CEM && role != generated.RM {
		return nil, errors.New("invalid role: must be CEM or RM")
	}

	var spv *[]string
	if role == generated.RM {
		spv = &supportedProtocols
	} else if role == generated.CEM && len(supportedProtocols) > 0 {
		spv = &supportedProtocols
	}

	// Generate a new ID
	id, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &Handshake{
		MessageID:                 id,
		MessageType:               "Handshake",
		Role:                      role,
		SupportedProtocolVersions: spv,
	}, nil
}

// NewVarHandshake creates multiple instances of Handshakes.
func NewVarHandshake(roles []generated.EnergyManagementRole, supportedProtocols [][]string) ([]*Handshake, error) {
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

// SetMessageID sets a new MessageID based on the pattern string given in the Handshake.
func (h *Handshake) SetMessageID(pattern string) *Handshake {
	validID := generated.TransformToValidID(pattern)
	newID := &generated.ID{Value: validID}
	h.MessageID = newID
	return h
}

// SetSupportedProtocolVersions sets the protocol versions in the Handshake.
func (h *Handshake) SetSupportedProtocolVersions(s []string) *Handshake {
	h.SupportedProtocolVersions = &s
	return h
}
