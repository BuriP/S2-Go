package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestNewRole(t *testing.T) {
	commodity := generated.GAS
	roleType := generated.ENERGY_PRODUCER

	// Test valid Role creation
	role, err := common.NewRole(commodity, roleType)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if role.Commodity != commodity {
		t.Errorf("expected Commodity %v, got %v", commodity, role.Commodity)
	}
	if role.Role != roleType {
		t.Errorf("expected Role %v, got %v", roleType, role.Role)
	}

	// Test creation with empty commodity
	_, err = common.NewRole("", roleType)
	if err == nil {
		t.Fatalf("expected error for empty commodity, got nil")
	}

	// Test creation with empty role
	_, err = common.NewRole(commodity, "")
	if err == nil {
		t.Fatalf("expected error for empty role, got nil")
	}
}
