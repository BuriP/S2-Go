package generated

// FRBCLeakageBehaviourElement represents an element of FRBC leakage behavior.
type FRBCLeakageBehaviourElement struct {
    FillLevelRange NumberRange `json:"fill_level_range" description:"The fill level range for which this FRBC.LeakageBehaviourElement applies. The start of the range must be less than the end of the range"`
    LeakageRate    float64     `json:"leakage_rate" description:"Indicates how fast the momentary fill level will decrease per second due to leakage within the given range of the fill level"`
}
