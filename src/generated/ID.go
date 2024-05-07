package generated

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
	
)

// ID represents an identifier expressed as a UUID.
type ID struct {
	Value uuid.UUID `json:"value" description:"An identifier expressed as a UUID"`
}

// NewID creates a new instance of ID with the provided value.
func NewID(value string) (*ID, error) {
	validID, err := TransformToValidID(value)
	if err!= nil {
		return nil, err
	}

	if !isValidID(validID.String()){
		return nil, fmt.Errorf("invalid Id Format")
	}

	return &ID{Value: validID}, nil
}

// isValidID checks if the provided value matches the regex pattern.
func isValidID(value string) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9\-_:]{2,64}$`)
	return pattern.MatchString(value)
}

// TransformToValidID transforms a string into a vlid ID type
func TransformToValidID(value string) (uuid.UUID, error) {
	// Here's a simple transformation that removes invalid characters:
	re := regexp.MustCompile(`[^a-zA-Z0-9\-_:]{2,64}$`)
	transformedValue := re.ReplaceAllString(value, "")
	// Ensure the length is within the valid range (UUID length).
	transformedValue = transformedValue[:min(len(transformedValue), 64)]

	// Attempt to parse the transformed value as a UUID.
	u, err := uuid.Parse(transformedValue)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID format: %v", err)
	}
	return u, nil
}