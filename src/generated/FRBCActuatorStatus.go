package generated

import(
	"time"
)

// FRBCActuatorStatus represents a FRBC actuator status model.
type FRBCActuatorStatus struct {
    ActiveOperationModeID   ID         `json:"active_operation_mode_id" description:"ID of the FRBC.OperationMode that is presently active."`
    ActuatorID              ID         `json:"actuator_id" description:"ID of the actuator this message refers to"`
    MessageID               ID         `json:"message_id" description:"ID of this message"`
    MessageType             string     `json:"message_type" description:"Type of the message"`
    OperationModeFactor     float64    `json:"operation_mode_factor" description:"The number indicates the factor with which the FRBC.OperationMode is configured"`
    PreviousOperationModeID *ID        `json:"previous_operation_mode_id,omitempty" description:"ID of the FRBC.OperationMode that was active before the present one"` //pointer because its optional
    TransitionTimestamp     *time.Time `json:"transition_timestamp,omitempty" description:"Time at which the transition from the previous FRBC.OperationMode to the active FRBC.OperationMode was initiated"`// pointer because its optional
}
