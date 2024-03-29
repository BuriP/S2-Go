package common
import (
	"time"
	"../generated" // Importing the generated package
)

// ToTimedelta converts Duration to time.Duration
func (d *generated.Duration) ToTimedelta() time.Duration {
	return time.Duration(d.Milliseconds) * time.Millisecond
}

// FromTimedelta creates a new Duration from time.Duration
func FromTimedelta(duration time.Duration) *generated.Duration {
	return &generated.Duration{
		Milliseconds: uint64(duration.Milliseconds()),
	}
}

// FromMilliseconds creates a new Duration from milliseconds
func FromMilliseconds(milliseconds uint64) *generated.Duration {
	return &generated.Duration{
		Milliseconds: milliseconds,
	}
}