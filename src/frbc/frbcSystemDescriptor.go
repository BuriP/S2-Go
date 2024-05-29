package frbc

import (
	"errors"
	"time"

	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCSystemDescription represents the system description in FRBC.
type FRBCSystemDescription struct {
	Actuators   []*FRBCActuatorDescriptor `json:"actuators" description:"Details of all Actuators." min_items:"1" max_items:"10"`
	MessageID   *generated.ID             `json:"message_id" description:"ID of this message"`
	MessageType string                    `json:"message_type" description:"FRBC.SystemDescription"`
	Storage     *FRBCStorageDescription   `json:"storage" description:"Details of the storage."`
	ValidFrom   time.Time                 `json:"valid_from" description:"Moment this FRBC.SystemDescription starts to be valid. If the system description is immediately valid, the DateTimeStamp should be now or in the past."`
}

// NewFRBCSystemDescription creates a new instance of FRBCSystemDescription if validated, otherwise returns an error.
func NewFRBCSystemDescription(actuators []*FRBCActuatorDescriptor, storage *FRBCStorageDescription, validFrom time.Time) (*FRBCSystemDescription, error) {
	if len(actuators) == 0 {
		return nil, errors.New("at least one actuator must be provided")
	}

	if len(actuators) > 10 {
		return nil, errors.New("at most 10 actuators can be provided")
	}

	messageID, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &FRBCSystemDescription{
		Actuators:   actuators,
		MessageID:   messageID,
		MessageType: "FRBC.SystemDescription",
		Storage:     storage,
		ValidFrom:   validFrom,
	}, nil
}

// GetActuators returns the actuators of the system description.
func (sd *FRBCSystemDescription) GetActuators() []*FRBCActuatorDescriptor {
	return sd.Actuators
}

// GetMessageID returns the message ID of the system description.
func (sd *FRBCSystemDescription) GetMessageID() *generated.ID {
	return sd.MessageID
}

// GetMessageType returns the message type of the system description.
func (sd *FRBCSystemDescription) GetMessageType() string {
	return sd.MessageType
}

// GetStorage returns the storage description of the system description.
func (sd *FRBCSystemDescription) GetStorage() *FRBCStorageDescription {
	return sd.Storage
}

// GetValidFrom returns the valid from timestamp of the system description.
func (sd *FRBCSystemDescription) GetValidFrom() time.Time {
	return sd.ValidFrom
}
