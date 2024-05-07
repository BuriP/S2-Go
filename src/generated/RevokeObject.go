//TODO: Remove from dir

package generated

// RevokeObject represents an object revocation.

type RevokeObject struct {
	MessageID   ID                 `json:"message_id" description:"ID of this message"`
	MessageType string             `json:"message_type" description:"Type of this message"`
	ObjectID    ID                 `json:"object_id" description:"ID of object that needs to be revoked"`
	ObjectType  *RevokableObject  `json:"object_type" description:"Type of object that needs to be revoked"`
}

