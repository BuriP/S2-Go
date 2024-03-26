package generated

import (

)

// ReceptionStatus represents the reception status of a message.
type ReceptionStatus struct {
    DiagnosticLabel    *string             `json:"diagnostic_label,omitempty" description:"Diagnostic label that can be used to provide additional information for debugging."` // USed a pointer due optional
    MessageType        string              `json:"message_type" description:"Type of the message"`
    Status             ReceptionStatusValues `json:"status" description:"Enumeration of status values"`
    SubjectMessageID   ID                  `json:"subject_message_id" description:"The message this ReceptionStatus refers to"`
}
