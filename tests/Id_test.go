package tests

import (
	"regexp"
	"testing"

	"github.com/BuriP/S2-Go/src/generated"
	"strings"
)

func TestTransformToValidID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Valid characters", "abc123-_:xyz", "abc123-_:xyz"},
		{"Invalid characters", "abc!@#123", "abc123"},
		{"Mixed characters", "a!b@c#1$2%3", "abc123"},
		{"Short string", "a", "aa"},
		{"Long string", strings.Repeat("a", 70), strings.Repeat("a", 64)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generated.TransformToValidID(tt.input)
			if got != tt.want {
				t.Errorf("TransformToValidID(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestNewID(t *testing.T) {
	t.Parallel()

	id, err := generated.NewID()
	if err != nil {
		t.Fatalf("error generating ID: %v", err)
	}

	// Validate the length of the generated ID
	if len(id.Value) < 2 || len(id.Value) > 64 {
		t.Errorf("generated ID length is not between 2 and 64: %s", id.Value)
	}

	// Validate the ID pattern
	pattern := regexp.MustCompile(`^[a-zA-Z0-9\-_:]{2,64}$`)
	if !pattern.MatchString(id.Value) {
		t.Errorf("generated ID does not match the expected pattern: %s", id.Value)
	}
}
