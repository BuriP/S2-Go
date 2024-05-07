package common

import (
	"github.com/BuriP/S2-Go/src/generated"
)


type ReceptionStatus struct {
    DiagnosticLabel    *string             `json:"diagnostic_label,omitempty" description:"Diagnostic label that can be used to provide additional information for debugging."` // USed a pointer due optional
    MessageType        string              `json:"message_type" description:"Type of the message"`
    Status             *generated.ReceptionStatusValues `json:"status" description:"Enumeration of status values"`
    SubjectMessageID   *generated.ID                  `json:"subject_message_id" description:"The message this ReceptionStatus refers to"`
}


//NewReceptionSatus creates new ReceptionStatusInstances with the subjectmessageid as the messages's ID.
func NewReceptionStatus(messageType string, status *generated.ReceptionStatusValues, subjectMessageID *generated.ID, label *string) *ReceptionStatus { // Maybe change so that it passes id directly
	return &ReceptionStatus{
		DiagnosticLabel: label,
		MessageType:      messageType,
		Status:           status,
		SubjectMessageID: subjectMessageID,
	}

}