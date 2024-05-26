package common

import (
	"errors"
	"time"
)

type Duration struct {
	Milliseconds uint64 `json:"milliseconds" description:"Duration in milliseconds"`
}

// NewDuration creates a new Duration
func NewDuration(d uint64) *Duration {
	return &Duration{
		Milliseconds: d,
	}
}

// ModifyDuration changes the duration of a Duration instance.
func (d *Duration) ModifyDuration(newDur uint64) *Duration {
	d.Milliseconds = newDur
	return d
}

// NewVarDuration creates n number of new Duration types
func NewVarDuration(n uint64, milliseconds []uint64) ([]*Duration, error) {
	if len(milliseconds) < int(n) {
		return nil, errors.New("not enough milliseconds provided")
	}

	durations := make([]*Duration, n)
	for i := uint64(0); i < n; i++ {
		durations[i] = NewDuration(milliseconds[i])
	}
	return durations, nil
}

// ToTimedelta converts Duration to time.Duration
func (d *Duration) ToTimedelta() time.Duration {
	return time.Duration(d.Milliseconds) * time.Millisecond
}

// FromTimedelta creates a new Duration from time.Duration
func FromTimedelta(duration time.Duration) *Duration {
	return &Duration{
		Milliseconds: uint64(duration / time.Millisecond),
	}
}

// FromMilliseconds creates a new Duration from milliseconds
func FromMilliseconds(time uint64) *Duration {
	return &Duration{
		Milliseconds: time,
	}
}
