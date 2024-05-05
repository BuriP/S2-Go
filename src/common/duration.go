package common
import (
	"time"

)

type Duration struct {
	Milliseconds uint64 `json:"millisecons" description:"Duration in milliseconds"`
}

// NewDuration creates a new Duration // ToDO test for the 
func NewDuration(d uint64) *Duration{
	return &Duration{
		Milliseconds : d,
	}
}

// ModifyDuration changes the duration of a Duration instance.
func(d * Duration) ModifyDuration(newdur uint64) *Duration {
	d.Milliseconds = newdur
	return d //Returns the instance modified
}

// NewVarDuration creates n number of new Duration types
func NewVarDuration(n uint64, milliseconds []uint64) []*Duration {
	var durations []*Duration
	for i := uint64(0); i < n; i++ {
		durations = append(durations, NewDuration(milliseconds[i]))
	}
	return durations
}

// ToTimedelta converts Duration to time.Duration
func (d *Duration) ToTimedelta() time.Duration {
	return time.Duration(d.Milliseconds)
		}

// FromTimedelta creates a new Duration from time.Duration
func FromTimedelta(duration time.Duration) *Duration {
	return &Duration{
		Milliseconds: uint64(duration.Milliseconds()),
	}
}
	
// FromMilliseconds creates a new Duration from milliseconds
func FromMilliseconds(time uint64) *Duration{
	return &Duration{
		Milliseconds: time,
	}
}