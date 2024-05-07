package common

import(

	"github.com/BuriP/S2-Go/src/generated"
)

type Role struct {
    Commodity generated.Commodity `json:"commodity" description:"Commodity the role refers to."`
    Role      generated.RoleType  `json:"role" description:"Role type of the Resource Manager for the given commodity"`
}

// NewRole creates a new Role instance.
func NewRole(commodity generated.Commodity, role generated.RoleType) *Role {
	return &Role{
		Commodity: commodity,
		Role:      role,
	}
}