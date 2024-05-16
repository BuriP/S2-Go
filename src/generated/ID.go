package generated

import (
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// ID represents an identifier expressed as a UUID.
type ID struct {
	Value string `json:"value" description:"An identifier expressed as a UUID"`
}

// TransformToValidID transforms a string into a vlid ID type
func TransformToValidID(value string) string {
	// Here's a simple transformation that removes invalid characters:
	re := regexp.MustCompile(`[^a-zA-Z0-9\-_:]{2,64}$`)
	transformedValue := re.ReplaceAllString(value, "")
	return transformedValue
}

// CreateNewID creates a new ID following the specific pattern.
func NewID() (*ID, error) {
	// Generating a new random UUID
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// Convert the UUID to string
	idString := strings.ReplaceAll(id.String(), "-", "")

	// Transform the string into a valid ID
	validID := TransformToValidID(idString)
	return &ID{Value: validID}, nil
}
