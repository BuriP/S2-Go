package tests

import (
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common"
)

func TestNewDuration(t *testing.T) {
	// Test NewDuration function
	d := common.NewDuration(100)
	if d.Milliseconds != 100 {
		t.Errorf("Expected milliseconds to be 100, got %d", d.Milliseconds)
	}
}

// Test ModifyDuration function
func TestModifyDuration(t *testing.T) {
	d := common.NewDuration(200)
	d.ModifyDuration(300)
	if d.Milliseconds != 300 {
		t.Errorf("Expected milliseconds to be 200 after ModifyDuration, got %d", d.Milliseconds)
	}
}

// Test ToTimedelta function
func TestToTimedelta(t *testing.T) {
	d := common.NewDuration(1000)
	td := d.ToTimedelta()
	if td != 1000*time.Millisecond {
		t.Errorf("Expected ToTimedelta to return %v, got %v", 1000*time.Millisecond, td)
	}
}

func TestFromTimedelta(t *testing.T) {
	td := 1000 * time.Millisecond
	d := common.FromTimedelta(td)
	if d.Milliseconds != 1000 {
		t.Errorf("expected 1000, got %d", d.Milliseconds)
	}
}

func TestFromMilliseconds(t *testing.T) {
	d := common.FromMilliseconds(1500)
	if d.Milliseconds != 1500 {
		t.Errorf("expected 1500, got %d", d.Milliseconds)
	}
}
