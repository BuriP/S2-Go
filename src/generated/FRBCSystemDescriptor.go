package generated

import(
	"time"
)

type FRBCSystemDescription struct {
	Actuators    []FRBCActuatorDescription `json:"actuators" description:"Details of all Actuators." min_items:"1" max_items:"10"`
	MessageID    ID                        `json:"message_id" description:"ID of this message"`
	MessageType  string                    `json:"message_type" description:"FRBC.SystemDescription"`
	Storage      FRBCStorageDescription    `json:"storage" description:"Details of the storage."`
	ValidFrom    time.Time                 `json:"valid_from" description:"Moment this FRBC.SystemDescription starts to be valid. If the system description is immediately valid, the DateTimeStamp should be now or in the past."`
}
