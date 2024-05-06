package generated

import (
	"regexp"
	"fmt"

)
// ID represents an identifier expressed as a UUID.
type ID struct {
	Value string `json:"value" description:"An identifier expressed as a UUID"`
}

// NewID creates a new instance of ID with the provided value.
func NewID(value string) (*ID, error) {
	// Validate the provided value against the regex pattern.
	if !isValidID(value) {
		return nil, fmt.Errorf("invalid ID format")
	}
	return &ID{Value: value}, nil
}

// isValidID checks if the provided value matches the regex pattern.
func isValidID(value string) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9\-_:]{2,64}$`)
	return pattern.MatchString(value)
}