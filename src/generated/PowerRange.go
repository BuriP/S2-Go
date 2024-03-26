package generated

// PowerRange represents a range of power values.
type PowerRange struct {
	CommodityQuantity CommodityQuantity `json:"commodity_quantity" description:"The power quantity the values refer to"`
	EndOfRange        float64           `json:"end_of_range" description:"Power value that defines the end of the range."`
	StartOfRange      float64           `json:"start_of_range" description:"Power value that defines the start of the range."`
}

