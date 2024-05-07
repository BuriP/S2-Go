package generated

// Role represents a role type of the Resource Manager for a given commodity.
type Role struct {
    Commodity Commodity `json:"commodity" description:"Commodity the role refers to."`
    Role      *RoleType  `json:"role" description:"Role type of the Resource Manager for the given commodity"`
}
