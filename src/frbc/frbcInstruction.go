package frbc

import (
	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"

	"time"
)

// FRBCInstruction represents an instruction in FRBC.
type FRBCInstruction struct {
	AbnormalCondition   bool          `json:"abnormal_condition" description:"Indicates if this is an instruction during an abnormal condition"`
	ActuatorID          *generated.ID `json:"actuator_id" description:"ID of the actuator this instruction belongs to"`
	ExecutionTime       time.Time     `json:"execution_time" description:"Time when instruction should be executed"`
	ID                  *generated.ID `json:"id" description:"ID of the instruction. Must be unique in the scope of the Resource Manager, for at least the duration of the session between Resource Manager and CEM"`
	MessageID           *generated.ID `json:"message_id" description:"ID of this message"`
	MessageType         string        `json:"message_type" description:"FRBC.Instruction"`
	OperationMode       *generated.ID `json:"operation_mode" description:"ID of the FRBC.OperationMode that should be activated"`
	OperationModeFactor float64       `json:"operation_mode_factor" description:"The number indicates the factor with which the FRBC.OperationMode should be configured. The factor should be greater than or equal to 0 and less than or equal to 1"`
}
