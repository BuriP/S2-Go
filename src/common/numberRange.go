package common

import (
	"fmt"
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)


type NumberRange struct {
	EndOfRange   float64 `json:"end_of_range" description:"Number that defines the end of the range"`
	StartOfRange float64 `json:"start_of_range" description:"Number that defines the start of the range"`
}



// NewNumberRange creates a new NumberRange instance
func NewNumberRange(start float64, end float64) (*NumberRange, error) {
	if end < start{
		return nil, fmt.Errorf("invalid range: start range should be lower than end range")
	}
	return &NumberRange{
		StartOfRange: start,
		EndOfRange:   end,
	}, nil
}

// ValidateStartEndOrder validates the order of start and end values
func (nr *NumberRange) ValidateStartEndOrder() error {
	if nr.StartOfRange > nr.EndOfRange {
		return errors.New("StartOFRange should not be higher than EndOfRange")
	}
	return nil
}

// Hash calculates the hash of the NumberRange instance
func (nr *NumberRange) Hash() string {
	return fmt.Sprintf("%f|%f", nr.StartOfRange, nr.EndOfRange)
}

// Equals checks if two NumberRange instances are equal
func (nr *NumberRange) Equals(other *generated.NumberRange) bool {
	if other == nil {
		return false
	}
	return nr.StartOfRange == other.StartOfRange && nr.EndOfRange == other.EndOfRange
}


