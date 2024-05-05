package generated

import (
	"regexp"

	"github.com/google/uuid"
)

type ID string

// Validate method checks if the ID has the required format.
func (id ID) Validate() bool {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9\-_:]{2,64}$`) // Defining the expression for the pattern.
	return pattern.MatchString(string(id))
}

// IDFromString converts a string into an ID.
func IDFromString(s string) ID {
	return ID(s)
}
