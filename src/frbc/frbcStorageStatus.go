package frbc

import (
	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCStorageStatus represents the status of FRBC storage.
type FRBCStorageStatus struct {
	MessageID        *generated.ID `json:"message_id" description:"ID of this message"`
	MessageType      string        `json:"message_type" description:"FRBC.StorageStatus"`
	PresentFillLevel float64       `json:"present_fill_level" description:"Present fill level of the Storage"`
}

// NewFRBCStorageStatus creates a new instance of FRBCStorageStatus.
func NewFRBCStorageStatus(presentFillLevel float64) (*FRBCStorageStatus, error) {
	messageID, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &FRBCStorageStatus{
		MessageID:        messageID,
		MessageType:      "FRBC.StorageStatus",
		PresentFillLevel: presentFillLevel,
	}, nil
}

// GetMessageID returns the message ID of the storage status.
func (s *FRBCStorageStatus) GetMessageID() *generated.ID {
	return s.MessageID
}

// GetMessageType returns the message type of the storage status.
func (s *FRBCStorageStatus) GetMessageType() string {
	return s.MessageType
}

// GetPresentFillLevel returns the present fill level of the storage.
func (s *FRBCStorageStatus) GetPresentFillLevel() float64 {
	return s.PresentFillLevel
}
