package generated

import(
	
)

// FRBCStorageStatus represents the status of FRBC storage.
type FRBCStorageStatus struct {
    MessageID         ID     `json:"message_id" description:"ID of this message"`
    MessageType       string `json:"message_type" description:"FRBC.StorageStatus"`
    PresentFillLevel  float64 `json:"present_fill_level" description:"Present fill level of the Storage"`
}
