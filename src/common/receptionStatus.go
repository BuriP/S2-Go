package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

type ReceptionStatus struct {
	DiagnosticLabel  *string                         `json:"diagnostic_label,omitempty" description:"Diagnostic label that can be used to provide additional information for debugging."` // USed a pointer due optional
	MessageType      string                          `json:"message_type" description:"Type of the message"`
	Status           generated.ReceptionStatusValues `json:"status" description:"Enumeration of status values"`
	SubjectMessageID *generated.ID                   `json:"subject_message_id" description:"The message this ReceptionStatus refers to"`
}

// NewReceptionSatus creates new ReceptionStatusInstances with the subjectmessageid as the messages's ID.
func NewReceptionStatus(status generated.ReceptionStatusValues, subjectMessageID *generated.ID, label *string) (*ReceptionStatus, error) { // Maybe change so that it passes id directly
	if status == "" {
		return nil, errors.New("status is required")
	}

	if subjectMessageID == nil {
		return nil, errors.New("subjectMessageID is required")
	}
	return &ReceptionStatus{
		DiagnosticLabel:  label,
		MessageType:      "ReceptionStatus",
		Status:           status,
		SubjectMessageID: subjectMessageID,
	}, nil

}
