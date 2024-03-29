package common

import (
	"fmt"
	"errors"
	"../generated"
)

// NumberRange represents a number range
type NumberRange struct {
	StartOfRange float64
	EndOfRange   float64
}

// NewNumberRange creates a new NumberRange instance
func NewNumberRange(start, end float64) *NumberRange {
	return &NumberRange{
		StartOfRange: start,
		EndOfRange:   end,
	}
}

// ValidateStartEndOrder validates the order of start and end values
func (nr *NumberRange) ValidateStartEndOrder() error {
	if nr.StartOfRange > nr.EndOfRange {
		return errors.New("start_of_range should not be higher than end_of_range")
	}
	return nil
}

// Hash calculates the hash of the NumberRange instance
func (nr *NumberRange) Hash() string {
	return fmt.Sprintf("%f|%f", nr.StartOfRange, nr.EndOfRange)
}

// Equals checks if two NumberRange instances are equal
func (nr *NumberRange) Equals(other *NumberRange) bool {
	if other == nil {
		return false
	}
	return nr.StartOfRange == other.StartOfRange && nr.EndOfRange == other.EndOfRange
}


