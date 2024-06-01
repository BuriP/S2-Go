// go:build ignore
//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/frbc"
)

func main() {

	// Create the Number Ranges
	numberRange1, err := common.NewNumberRange(10.0, 20.0)
	if err != nil {
		log.Fatalf("Error creating number range 1: %v", err)
	}
	numberRange2, err := common.NewNumberRange(30.0, 40.0)
	if err != nil {
		log.Fatalf("Error creating number range 2: %v", err)
	}

	// Create the FRBCLeakageBehaviourElements
	frbcleakagebehaviourelement1, err := frbc.NewFRBCLeakageBehaviourElement(numberRange1, 10.0)
	if err != nil {
		log.Fatalf("error creating leakage behaviour element 1: %v", err)
	}

	frbcleakagebehaviourelement2, err := frbc.NewFRBCLeakageBehaviourElement(numberRange2, 20.0)
	if err != nil {
		log.Fatalf("error creating leakage behaviour element 1: %v", err)
	}

	// Group the FRBCLeakageBehaviourElements
	frbcLeakageBehaviourElements := []*frbc.FRBCLeakageBehaviourElement{frbcleakagebehaviourelement1, frbcleakagebehaviourelement2}

	time := time.Now()

	//Createe new FRBCLeakageBehaviour
	frbcLeakageBehaviour, err := frbc.NewFRBCLeakageBehaviour(frbcLeakageBehaviourElements, time)
	if err != nil {
		log.Fatalf("Error creating leakage behaviour element 2: %v", err)
	}

	//Creating the JSON FRBCLEakageBehaviour message
	frbcLeakageBehaviourJson, err := json.MarshalIndent(frbcLeakageBehaviour, "", " ")
	if err != nil {
		log.Fatalf("Error marshalling the frbc leakage behavior: %v", err)
	}

	fmt.Printf("FRBCLeakageBehavior: %s", frbcLeakageBehaviourJson)

}

/* OUTPUT:
FRBCLeakageBehavior: {                                                                                                                                                                                             "elements": [                                                                                                                                                                                                      {                                                                                                                                                                                                                  "fill_level_range": {                                                                                                                                                                                              "end_of_range": 20,                                                                                                                                                                                               "start_of_range": 10                                                                                                                                                                                             },                                                                                                                                                                                                                "leakage_rate": 10                                                                                                                                                                                               },                                                                                                                                                                                                                {                                                                                                                                                                                                                  "fill_level_range": {                                                                                                                                                                                              "end_of_range": 40,                                                                                                                                                                                               "start_of_range": 30                                                                                                                                                                                             },                                                                                                                                                                                                                "leakage_rate": 20                                                                                                                                                                                               }                                                                                                                                                                                                                ],                                                                                                                                                                                                                "message_id": {                                                                                                                                                                                                    "value": "656c84c8821a4bc3948c7212ba2b119d"                                                                                                                                                                      },                                                                                                                                                                                                                "message_type": "FRBC.LeakageBehaviour",                                                                                                                                                                          "valid_from": "2024-06-02T01:10:01.281232+02:00"                                                                                                                                                                 }

*/
