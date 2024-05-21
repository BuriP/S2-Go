package common

import (
	"github.com/BuriP/S2-Go/src/generated"
)

type RevokeObject struct {
	MessageID   *generated.ID             `json:"message_id" description:"ID of this message"`
	MessageType string                    `json:"message_type" description:"Type of this message"`
	ObjectID    *generated.ID             `json:"object_id" description:"ID of object that needs to be revoked"`
	ObjectType  generated.RevokableObject `json:"object_type" description:"Type of object that needs to be revoked"`
}

// NewRevokeObject crewates a new instance of the RevokeObject and an error
func NewRevokeObject(messageID string, objectID *generated.ID, messageType string, objectType generated.RevokableObject) (*RevokeObject, error) {
	messageIDValue, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &RevokeObject{
		MessageID:   messageIDValue,
		MessageType: messageType,
		ObjectID:    objectID,
		ObjectType:  objectType,
	}, nil
}

