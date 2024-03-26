// TO DO: Look at the slice and arrays prob. Also Check for the max and mins of the 

package generated

type ResourceManagerDetails struct {
	AvailableControlTypes        []ControlType         `json:"available_control_types" description:"The control types supported by this Resource Manager." min_items:"1" max_items:"5"`
	Currency                     *Currency             `json:"currency,omitempty" description:"Currency to be used for all information regarding costs. Mandatory if cost information is published."`
	FirmwareVersion              *string               `json:"firmware_version,omitempty" description:"Version identifier of the firmware used in the device (provided by the manufacturer)"`
	InstructionProcessingDelay   Duration              `json:"instruction_processing_delay" description:"The average time the combination of Resource Manager and HBES/BACS/SASS or (Smart) device needs to process and execute an instruction"`
	Manufacturer                 *string               `json:"manufacturer,omitempty" description:"Name of Manufacturer"`
	MessageID                    ID                    `json:"message_id" description:"ID of this message"`
	MessageType                  string                `json:"message_type" description:"Type of this message" const:"ResourceManagerDetails"`
	Model                        *string               `json:"model,omitempty" description:"Name of the model of the device (provided by the manufacturer)"`
	Name                         *string               `json:"name,omitempty" description:"Human readable name given by user"`
	ProvidesForecast             bool                  `json:"provides_forecast" description:"Indicates whether the ResourceManager is able to provide PowerForecasts"`
	ProvidesPowerMeasurementType []CommodityQuantity   `json:"provides_power_measurement_types" description:"Array of all CommodityQuantities that this Resource Manager can provide measurements for." min_items:"1" max_items:"10"`
	ResourceID                   ID                    `json:"resource_id" description:"Identifier of the Resource Manager. Must be unique within the scope of the CEM."`
	Roles                        []Role                `json:"roles" description:"Each Resource Manager provides one or more energy Roles" min_items:"1" max_items:"3"`
	SerialNumber                 *string               `json:"serial_number,omitempty" description:"Serial number of the device (provided by the manufacturer)"`
}
