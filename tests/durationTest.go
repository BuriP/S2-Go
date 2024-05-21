package tests

import (
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common"
)

func TestDuration(t *testing.T) {
	// Test NewDuration function
	d := common.NewDuration(100)
	if d.Milliseconds != 100 {
		t.Errorf("Expected milliseconds to be 100, got %d", d.Milliseconds)
	}

	// Test ModifyDuration function
	d.ModifyDuration(200)
	if d.Milliseconds != 200 {
		t.Errorf("Expected milliseconds to be 200 after ModifyDuration, got %d", d.Milliseconds)
	}

	// Test ToTimedelta function
	expectedDuration := time.Duration(200) * time.Millisecond
	if d.ToTimedelta() != expectedDuration {
		t.Errorf("Expected ToTimedelta to return %v, got %v", expectedDuration, d.ToTimedelta())
	}

	// Test FromTimedelta function
	fromDuration := common.FromTimedelta(expectedDuration)
	if fromDuration.Milliseconds != 200 {
		t.Errorf("Expected FromTimedelta to return milliseconds as 200, got %d", fromDuration.Milliseconds)
	}

	// Test FromMilliseconds function
	fromMilliseconds := common.FromMilliseconds(300)
	if fromMilliseconds.Milliseconds != 300 {
		t.Errorf("Expected FromMilliseconds to return milliseconds as 300, got %d", fromMilliseconds.Milliseconds)
	}
}
