package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
)

func TestNewNumberRange(t *testing.T) {
	t.Parallel()

	// Test valid range
	start := 1.0
	end := 5.0
	nr, err := common.NewNumberRange(start, end)
	if err != nil {
		t.Errorf("Unexpected error creating number range: %v", err)
	}
	if nr.StartOfRange != start || nr.EndOfRange != end {
		t.Errorf("Expected range [%f, %f], got [%f, %f]", start, end, nr.StartOfRange, nr.EndOfRange)
	}

	// Test invalid range (end < start)
	start = 5.0
	end = 1.0
	_, err = common.NewNumberRange(start, end)
	if err == nil {
		t.Error("Expected error for invalid range, got nil")
	} else if err.Error() != "invalid range: start range should be lower than end range" {
		t.Errorf("Expected error message 'invalid range: start range should be lower than end range', got '%v'", err)
	}
}

func TestValidateStartEndOrder(t *testing.T) {
	t.Parallel()

	// Test valid order
	nr := &common.NumberRange{
		StartOfRange: 1.0,
		EndOfRange:   5.0,
	}
	err := nr.ValidateStartEndOrder()
	if err != nil {
		t.Errorf("Unexpected error validating start and end order: %v", err)
	}

	// Test invalid order (start > end)
	nr = &common.NumberRange{
		StartOfRange: 5.0,
		EndOfRange:   1.0,
	}
	err = nr.ValidateStartEndOrder()
	if err == nil {
		t.Error("Expected error for invalid order, got nil")
	} else if err.Error() != "StartOFRange should not be higher than EndOfRange" {
		t.Errorf("Expected error message 'StartOFRange should not be higher than EndOfRange', got '%v'", err)
	}
}

func TestHash(t *testing.T) {
	t.Parallel()

	// Test hash calculation
	nr := &common.NumberRange{
		StartOfRange: 1.0,
		EndOfRange:   5.0,
	}
	expectedHash := "1.000000|5.000000"
	hash := nr.Hash()
	if hash != expectedHash {
		t.Errorf("Expected hash '%s', got '%s'", expectedHash, hash)
	}
}
