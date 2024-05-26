package common

import (
	"errors"
	"github.com/BuriP/S2-Go/src/generated"
)

type ResourceManagerDetails struct {
	AvailableControlTypes        []generated.ControlType       `json:"available_control_types" description:"The control types supported by this Resource Manager." min_items:"1" max_items:"5"`
	Currency                     generated.Currency            `json:"currency,omitempty" description:"Currency to be used for all information regarding costs. Mandatory if cost information is published."`
	FirmwareVersion              *string                       `json:"firmware_version,omitempty" description:"Version identifier of the firmware used in the device (provided by the manufacturer)"`
	InstructionProcessingDelay   *Duration                     `json:"instruction_processing_delay" description:"The average time the combination of Resource Manager and HBES/BACS/SASS or (Smart) device needs to process and execute an instruction"`
	Manufacturer                 *string                       `json:"manufacturer,omitempty" description:"Name of Manufacturer"`
	MessageID                    *generated.ID                 `json:"message_id" description:"ID of this message"`
	MessageType                  string                        `json:"message_type" description:"Type of this message" const:"ResourceManagerDetails"`
	Model                        *string                       `json:"model,omitempty" description:"Name of the model of the device (provided by the manufacturer)"`
	Name                         *string                       `json:"name,omitempty" description:"Human readable name given by user"`
	ProvidesForecast             bool                          `json:"provides_forecast" description:"Indicates whether the ResourceManager is able to provide PowerForecasts"`
	ProvidesPowerMeasurementType []generated.CommodityQuantity `json:"provides_power_measurement_types" description:"Array of all CommodityQuantities that this Resource Manager can provide measurements for." min_items:"1" max_items:"10"`
	ResourceID                   *generated.ID                 `json:"resource_id" description:"Identifier of the Resource Manager. Must be unique within the scope of the CEM."`
	Roles                        []*Role                       `json:"roles" description:"Each Resource Manager provides one or more energy Roles" min_items:"1" max_items:"3"`
	SerialNumber                 *string                       `json:"serial_number,omitempty" description:"Serial number of the device (provided by the manufacturer)"`
}

// NewResourceManagerDetails creates a new ResourceManagerDetails instance and returns an error if validation fails
func NewResourceManagerDetails(
	availableControlTypes []generated.ControlType,
	currency generated.Currency,
	firmwareVersion, manufacturer, model, name string,
	instructionProcessingDelay *Duration,
	providesForecast bool,
	providesPowerMeasurementType []generated.CommodityQuantity,
	resourceID *generated.ID,
	roles []*Role,

	serialNumber string,
) (*ResourceManagerDetails, error) {
	if len(availableControlTypes) < 1 || len(availableControlTypes) > 5 {
		return nil, errors.New("availableControlTypes must contain between 1 and 5 items")
	}

	if len(providesPowerMeasurementType) < 1 || len(providesPowerMeasurementType) > 10 {
		return nil, errors.New("providesPowerMeasurementType must contain between 1 and 10 items")
	}

	if len(roles) < 1 || len(roles) > 3 {
		return nil, errors.New("roles must contain between 1 and 3 items")
	}

	if instructionProcessingDelay == nil {
		return nil, errors.New("instructionProcessingDelay is required")
	}

	if resourceID == nil {
		return nil, errors.New("resourceID is required")
	}

	messageID, err := generated.NewID()
	if err != nil {
		return nil, err
	}

	return &ResourceManagerDetails{
		AvailableControlTypes:        availableControlTypes,
		Currency:                     currency,
		FirmwareVersion:              &firmwareVersion,
		InstructionProcessingDelay:   instructionProcessingDelay,
		Manufacturer:                 &manufacturer,
		MessageID:                    messageID,
		MessageType:                  "ResourceManagerDetails",
		Model:                        &model,
		Name:                         &name,
		ProvidesForecast:             providesForecast,
		ProvidesPowerMeasurementType: providesPowerMeasurementType,
		ResourceID:                   resourceID,
		Roles:                        roles,
		SerialNumber:                 &serialNumber,
	}, nil
}
