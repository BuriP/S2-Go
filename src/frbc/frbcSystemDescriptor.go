package frbc

import (
	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
	"time"
)

type FRBCSystemDescription struct {
	Actuators   *[]generated.FRBCActuatorDescription `json:"actuators" description:"Details of all Actuators." min_items:"1" max_items:"10"`
	MessageID   *generated.ID                        `json:"message_id" description:"ID of this message"`
	MessageType string                               `json:"message_type" description:"FRBC.SystemDescription"`
	Storage     *generated.FRBCStorageDescription    `json:"storage" description:"Details of the storage."`
	ValidFrom   time.Time                            `json:"valid_from" description:"Moment this FRBC.SystemDescription starts to be valid. If the system description is immediately valid, the DateTimeStamp should be now or in the past."`
}
