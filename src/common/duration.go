package common
import (
	"time"
	"S2-Go/src/generated" // Importing the generated package
)

// ToTimedelta converts Duration to time.Duration
func ToTimedelta(d time.Duration) *generated.Duration {
	return &generated.Duration{
		Milliseconds : float64(d.Milliseconds()),
		}
}
// FromTimedelta creates a new Duration from time.Duration
func FromTimedelta(duration time.Duration) *generated.Duration {
	return &generated.Duration{
		Milliseconds: float64(duration.Milliseconds()),
	}
}

// FromMilliseconds creates a new Duration from milliseconds
func FromMilliseconds(milliseconds float64) *generated.Duration {
	return &generated.Duration{
		Milliseconds: milliseconds,
	}
}