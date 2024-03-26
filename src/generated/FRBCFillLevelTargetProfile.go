package generated

// FRBCFillLevelTargetProfileElement represents an element of FRBC fill level target profile.

type FRBCFillLevelTargetProfileElement struct {
    Duration       Duration   `json:"duration" description:"The duration of the element"`
    FillLevelRange NumberRange `json:"fill_level_range" description:"The target range in which the fill_level must be for the time period during which the element is active"`
}
