package generated

// PowerValue represents a power value as a float.

type PowerValue struct {
	CommodityQuantity CommodityQuantity `json:"commodity_quantity" description:"The power quantity the value refers to"`
	Value             float64           `json:"value" description:"Power value expressed in the unit associated with the CommodityQuantity"`
}
