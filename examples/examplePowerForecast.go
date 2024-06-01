//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func main() {
	// Create power value 1
	powerValue1, err := common.NewPowerForecastValue(generated.ELECTRIC_POWER_L1, 2.0)
	if err != nil {
		log.Fatalf("Error creating power value 1: %v", err)
	}

	// Create power value 2
	powerValue2, err := common.NewPowerForecastValue(generated.ELECTRIC_POWER_L2, 1.0)
	if err != nil {
		log.Fatalf("Error creating power value 2: %v", err)
	}

	// Combine power values
	powerValues := []*common.PowerForecastValue{powerValue1, powerValue2}

	// Create durations
	duration := common.NewDuration(10)
	duration2 := common.NewDuration(20)

	// Create power forecast element
	powerElements1, err := common.NewPowerForecastElement(duration, powerValues)
	if err != nil {
		log.Fatalf("Error creating power forecast element : %v", err)
	}
	powerElements2, err := common.NewPowerForecastElement(duration2, powerValues)
	if err != nil {
		log.Fatalf("Error creating power forecast element : %v", err)
	}

	powerElements := []*common.PowerForecastElement{powerElements1, powerElements2}

	// Create power forecast
	powerForecast, err := common.NewPowerForecast(powerElements, 10)
	if err != nil {
		log.Fatalf("Error creating power forecast: %v", err)
	}

	// Create Json message from PowerForecast
	powerForecastJson, err := json.MarshalIndent(powerForecast, "", " ")
	if err != nil {
		log.Fatalf("Error marshalling power forecast: %v", err)
	}

	fmt.Printf("PowerForecast: %s\n", powerForecastJson)

}

/* OUTPUT:
PowerForecast: {                                                                                                                                                                                                   "elements": [                                                                                                                                                                                                      {                                                                                                                                                                                                                  "duration": {                                                                                                                                                                                                      "milliseconds": 10                                                                                                                                                                                               },                                                                                                                                                                                                                "power_values": [                                                                                                                                                                                                  {                                                                                                                                                                                                                  "commodity_quantity": "ELECTRIC.POWER.L1",                                                                                                                                                                        "value_expected": 2                                                                                                                                                                                              },                                                                                                                                                                                                                {                                                                                                                                                                                                                  "commodity_quantity": "ELECTRIC.POWER.L2",                                                                                                                                                                        "value_expected": 1                                                                                                                                                                                              }                                                                                                                                                                                                                ]                                                                                                                                                                                                                },                                                                                                                                                                                                                {                                                                                                                                                                                                                  "duration": {                                                                                                                                                                                                      "milliseconds": 20                                                                                                                                                                                               },                                                                                                                                                                                                                "power_values": [                                                                                                                                                                                                  {                                                                                                                                                                                                                  "commodity_quantity": "ELECTRIC.POWER.L1",                                                                                                                                                                        "value_expected": 2                                                                                                                                                                                              },                                                                                                                                                                                                                {                                                                                                                                                                                                                  "commodity_quantity": "ELECTRIC.POWER.L2",                                                                                                                                                                        "value_expected": 1                                                                                                                                                                                              }                                                                                                                                                                                                                ]                                                                                                                                                                                                                }                                                                                                                                                                                                                ],                                                                                                                                                                                                                "message_id": {                                                                                                                                                                                                    "value": "5583e42268ef490ab713c6c6d8997199"                                                                                                                                                                      },                                                                                                                                                                                                                "message_type": "PowerForecast",                                                                                                                                                                                  "start_time": "1970-01-01T01:00:10+01:00"                                                                                                                                                                        }

*/
