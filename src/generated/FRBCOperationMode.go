// TO DO : Look at the slice and array issue

package generated

type FRBCOperationMode struct {
	AbnormalConditionOnly bool                      `json:"abnormal_condition_only" description:"Indicates if this FRBC.OperationMode may only be used during an abnormal condition"`
	DiagnosticLabel       *string                   `json:"diagnostic_label,omitempty" description:"Human readable name/description of the FRBC.OperationMode. This element is only intended for diagnostic purposes and not for HMI applications."`
	Elements              []FRBCOperationModeElement `json:"elements" description:"List of FRBC.OperationModeElements, which describe the properties of this FRBC.OperationMode depending on the fill_level. The fill_level_ranges of the items in the Array must be contiguous."`
	ID                    ID                        `json:"id" description:"ID of the FRBC.OperationMode. Must be unique in the scope of the FRBC.ActuatorDescription in which it is used."`
}
