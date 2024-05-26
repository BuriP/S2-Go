package frbc

import (
	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCStorageStatus represents the status of FRBC storage.
type FRBCStorageStatus struct {
	MessageID        *generated.ID `json:"message_id" description:"ID of this message"`
	MessageType      string        `json:"message_type" description:"FRBC.StorageStatus"`
	PresentFillLevel float64       `json:"present_fill_level" description:"Present fill level of the Storage"`
}
