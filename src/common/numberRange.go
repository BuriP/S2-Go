package common

import (
	"fmt"
	"errors"
	"S2-Go/src/generated"
)



// NewNumberRange creates a new NumberRange instance
func NewNumberRange(start, end float64) *generated.NumberRange {
	return &generated.NumberRange{
		StartOfRange: start,
		EndOfRange:   end,
	}
}

// ValidateStartEndOrder validates the order of start and end values
func (nr *generated.NumberRange) ValidateStartEndOrder() error {
	if nr.StartOfRange > nr.EndOfRange {
		return errors.New("start_of_range should not be higher than end_of_range")
	}
	return nil
}

// Hash calculates the hash of the NumberRange instance
func (nr *generated.NumberRange) Hash() string {
	return fmt.Sprintf("%f|%f", nr.StartOfRange, nr.EndOfRange)
}

// Equals checks if two NumberRange instances are equal
func (nr *generated.NumberRange) Equals(other *generated.NumberRange) bool {
	if other == nil {
		return false
	}
	return nr.StartOfRange == other.StartOfRange && nr.EndOfRange == other.EndOfRange
}


