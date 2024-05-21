package generated

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// ID represents an identifier expressed as a UUID.
type ID struct {
	Value string `json:"value" description:"An identifier expressed as a UUID"`
}

// TransformToValidID transforms a string into a valid ID type
func TransformToValidID(value string) string {
	// Only keep valid characters
	re := regexp.MustCompile(`[^a-zA-Z0-9\-_:]`)
	transformedValue := re.ReplaceAllString(value, "")

	// Ensure the length is within the valid range
	if len(transformedValue) < 2 {
		transformedValue = transformedValue + strings.Repeat("a", 2-len(transformedValue))
	} else if len(transformedValue) > 64 {
		transformedValue = transformedValue[:64]
	}

	return transformedValue
}

// NewID creates a new ID following the specific pattern.
func NewID() (*ID, error) {
	// Generating a new random UUID
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// Convert the UUID to string and remove hyphens
	idString := strings.ReplaceAll(id.String(), "-", "")

	// Transform the string into a valid ID
	validID := TransformToValidID(idString)

	// Ensure the transformed ID meets the length requirements
	if len(validID) < 2 || len(validID) > 64 {
		return nil, fmt.Errorf("invalid ID length after transformation: %s", validID)
	}

	return &ID{Value: validID}, nil
}
