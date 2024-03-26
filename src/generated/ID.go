// Code may need changes still!.

package generated

import (
	"regexp"
)

type ID string

// Validate method checks if the ID has the required format.
func (id ID) Validate() bool {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9\-_:]{2,64}$`) // Defining the expression for the pattern.
	return pattern.MatchString(string(id))
}

