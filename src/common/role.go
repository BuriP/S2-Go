package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

type Role struct {
	Commodity generated.Commodity `json:"commodity" description:"Commodity the role refers to."`
	Role      generated.RoleType  `json:"role" description:"Role type of the Resource Manager for the given commodity"`
}

// NewRole creates a new Role instance and returns an error if validation fails
func NewRole(commodity generated.Commodity, role generated.RoleType) (*Role, error) {
	if commodity == "" {
		return nil, errors.New("commodity is required")
	}
	if role == "" {
		return nil, errors.New("role is required")
	}

	return &Role{
		Commodity: commodity,
		Role:      role,
	}, nil
}
