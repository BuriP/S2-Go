package frbc

import (
	"errors"
	"time"

	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCLeakageBehaviour represents the leakage behaviour of the buffer.
type FRBCLeakageBehaviour struct {
	Elements    []*FRBCLeakageBehaviourElement `json:"elements" description:"List of elements that model the leakage behaviour of the buffer. The fill_level_ranges of the elements must be contiguous."`
	MessageID   *generated.ID                  `json:"message_id" description:"ID of this message"`
	MessageType string                         `json:"message_type" description:"FRBC.LeakageBehaviour"`
	ValidFrom   time.Time                      `json:"valid_from" description:"Moment this FRBC.LeakageBehaviour starts to be valid. If the FRBC.LeakageBehaviour is immediately valid, the DateTimeStamp should be now or in the past."`
}

// NewFRBCLeakageBehaviour creates a new instance of FRBCLeakageBehaviour if validated, otherwise returns an error.
func NewFRBCLeakageBehaviour(elements []*FRBCLeakageBehaviourElement, validFrom time.Time) (*FRBCLeakageBehaviour, error) {
	if len(elements) == 0 {
		return nil, errors.New("at least one element must be provided")
	}

	messageID, err := generated.NewID()
	if err != nil {
		return nil, err
	}
	// Ensure elements are in chronological order
	for i := 1; i < len(elements); i++ {
		if elements[i].FillLevelRange == nil || elements[i-1].FillLevelRange == nil {
			return nil, errors.New("each element must have a duration")
		}
		if elements[i].FillLevelRange.StartOfRange < elements[i-1].FillLevelRange.StartOfRange {
			return nil, errors.New("elements must be in chronological order")
		}
	}

	return &FRBCLeakageBehaviour{
		Elements:    elements,
		MessageID:   messageID,
		MessageType: "FRBC.LeakageBehaviour",
		ValidFrom:   validFrom,
	}, nil
}

// GetElements returns the elements of the leakage behaviour.
func (lb *FRBCLeakageBehaviour) GetElements() []*FRBCLeakageBehaviourElement {
	return lb.Elements
}

// GetMessageID returns the message ID of the leakage behaviour.
func (lb *FRBCLeakageBehaviour) GetMessageID() *generated.ID {
	return lb.MessageID
}

// GetMessageType returns the message type of the leakage behaviour.
func (lb *FRBCLeakageBehaviour) GetMessageType() string {
	return lb.MessageType
}

// GetValidFrom returns the valid from timestamp of the leakage behaviour.
func (lb *FRBCLeakageBehaviour) GetValidFrom() time.Time {
	return lb.ValidFrom
}
