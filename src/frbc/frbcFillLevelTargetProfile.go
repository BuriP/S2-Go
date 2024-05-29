package frbc

import (
	"errors"
	"time"

	"github.com/BuriP/S2-Go/src/generated"
)

// FRBCFillLevelTargetProfile represents a target profile for fill levels.
type FRBCFillLevelTargetProfile struct {
	Elements    []*FRBCFillLevelTargetProfileElement `json:"elements" description:"List of different fill levels that have to be targeted within a given duration. There shall be at least one element. Elements must be placed in chronological order."`
	MessageID   *generated.ID                        `json:"message_id" description:"ID of this message"`
	MessageType string                               `json:"message_type" description:"FRBCFillLevelTargetProfile" const:"true"`
	StartTime   time.Time                            `json:"start_time" description:"Time at which the FRBC.FillLevelTargetProfile starts."`
}

// NewFRBCFillLevelTargetProfile creates a new instance of FRBCFillLevelTargetProfile if validated, otherwise returns an error.
func NewFRBCFillLevelTargetProfile(elements []*FRBCFillLevelTargetProfileElement, startTime time.Time) (*FRBCFillLevelTargetProfile, error) {
	if len(elements) == 0 {
		return nil, errors.New("at least one element must be provided")
	}

	// Ensure elements are in chronological order
	for i := 1; i < len(elements); i++ {
		if elements[i].Duration.Milliseconds < elements[i-1].Duration.Milliseconds {
			return nil, errors.New("elements must be in chronological order")
		}
	}

	messageID, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &FRBCFillLevelTargetProfile{
		Elements:    elements,
		MessageID:   messageID,
		MessageType: "FRBCFillLevelTargetProfile",
		StartTime:   startTime,
	}, nil
}

// GetElements returns the elements of the fill level target profile.
func (p *FRBCFillLevelTargetProfile) GetElements() []*FRBCFillLevelTargetProfileElement {
	return p.Elements
}

// GetMessageID returns the message ID of the fill level target profile.
func (p *FRBCFillLevelTargetProfile) GetMessageID() *generated.ID {
	return p.MessageID
}

// GetMessageType returns the message type of the fill level target profile.
func (p *FRBCFillLevelTargetProfile) GetMessageType() string {
	return p.MessageType
}

// GetStartTime returns the start time of the fill level target profile.
func (p *FRBCFillLevelTargetProfile) GetStartTime() time.Time {
	return p.StartTime
}
