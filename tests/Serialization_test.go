package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
)

func TestHandshakeSerialization(t *testing.T) {
	role := generated.CEM
	supportedProtocols := []string{"1.0", "2.0"}
	handshake, err := common.NewHandshake(role, supportedProtocols)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	jsonData, err := json.Marshal(handshake)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var deserializedHandshake common.Handshake
	err = json.Unmarshal(jsonData, &deserializedHandshake)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if deserializedHandshake.Role != role {
		t.Errorf("expected role %v, got %v", role, deserializedHandshake.Role)
	}
}

func TestHandshakeResponseSerialization(t *testing.T) {
	handshake, _ := common.NewHandshake(generated.CEM, []string{"1.0"})
	handshakeResponse, err := common.NewHandshakeResponse(handshake)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	jsonData, err := json.Marshal(handshakeResponse)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var deserializedHandshakeResponse common.HandshakeResponse
	err = json.Unmarshal(jsonData, &deserializedHandshakeResponse)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if deserializedHandshakeResponse.SelectedProtocolVersion == nil {
		t.Errorf("expected selected protocol version, got nil")
	}
}

func TestInstructionStatusUpdateSerialization(t *testing.T) {
	statusType := generated.NEW
	isu, err := common.NewInstructionStatusUpdate(statusType)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	jsonData, err := json.Marshal(isu)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var deserializedISU common.InstructionStatusUpdate
	err = json.Unmarshal(jsonData, &deserializedISU)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if deserializedISU.StatusType != statusType {
		t.Errorf("expected status type %v, got %v", statusType, deserializedISU.StatusType)
	}
}

func TestPowerForecastSerialization(t *testing.T) {
	pfv1, _ := common.NewPowerForecastValue(generated.ELECTRIC_POWER_L1, 100.0)
	duration := common.NewDuration(3600) // 1 hour
	element, _ := common.NewPowerForecastElement(duration, &[]common.PowerForecastValue{*pfv1})
	elements := []common.PowerForecastElement{*element}
	startTime := time.Now().Unix()

	pf, err := common.NewPowerForecast(&elements, startTime)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	jsonData, err := json.Marshal(pf)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var deserializedPF common.PowerForecast
	err = json.Unmarshal(jsonData, &deserializedPF)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(*deserializedPF.Elements) != len(elements) {
		t.Errorf("expected %d elements, got %d", len(elements), len(*deserializedPF.Elements))
	}
}

func TestPowerMeasurementSerialization(t *testing.T) {
	powerValues := []common.PowerValue{
		{
			CommodityQuantity: generated.ELECTRIC_POWER_L1,
			Value:             100.0,
		},
	}
	pm, err := common.NewPowerMeasurement(&powerValues)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	jsonData, err := json.Marshal(pm)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var deserializedPM common.PowerMeasurement
	err = json.Unmarshal(jsonData, &deserializedPM)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(*deserializedPM.Values) != len(powerValues) {
		t.Errorf("expected %d power values, got %d", len(powerValues), len(*deserializedPM.Values))
	}
}

func TestReceptionStatusSerialization(t *testing.T) {
	status := generated.OK
	messageID, _ := generated.NewID()
	label := "Test Label"
	rs, err := common.NewReceptionStatus(&status, messageID, &label)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	jsonData, err := json.Marshal(rs)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var deserializedRS common.ReceptionStatus
	err = json.Unmarshal(jsonData, &deserializedRS)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if *deserializedRS.Status != status {
		t.Errorf("expected status %v, got %v", status, *deserializedRS.Status)
	}
}

func TestResourceManagerDetailsSerialization(t *testing.T) {
	controlTypes := []generated.ControlType{generated.NOT_CONTROLABLE}
	measurementTypes := []generated.CommodityQuantity{generated.ELECTRIC_POWER_L1}
	role, err := common.NewRole(generated.GAS, generated.ENERGY_CONSUMER)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	roles := []*common.Role{role}
	duration := common.NewDuration(5000)
	resourceID, err := generated.NewID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	rmd, err := common.NewResourceManagerDetails(
		controlTypes,
		generated.Currency("USD"),
		"1.0.0",
		"Manufacturer",
		"Model1",
		"ResourceManager1",
		duration,
		true,
		measurementTypes,
		resourceID,
		roles,
		"123456789",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	jsonData, err := json.Marshal(rmd)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var deserializedRMD common.ResourceManagerDetails
	err = json.Unmarshal(jsonData, &deserializedRMD)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if deserializedRMD.ResourceID.Value != resourceID.Value {
		t.Errorf("expected resource ID %v, got %v", resourceID.Value, deserializedRMD.ResourceID.Value)
	}
}

func TestRevokeObjectSerialization(t *testing.T) {
	objectID, _ := generated.NewID()
	revokeObject, err := common.NewRevokeObject(objectID, generated.PEBC_PowerConstraints)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	jsonData, err := json.Marshal(revokeObject)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var deserializedRevokeObject common.RevokeObject
	err = json.Unmarshal(jsonData, &deserializedRevokeObject)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if deserializedRevokeObject.ObjectID.Value != objectID.Value {
		t.Errorf("expected object ID %v, got %v", objectID.Value, deserializedRevokeObject.ObjectID.Value)
	}
}

func TestSelectControlTypeSerialization(t *testing.T) {
	selectControlType, err := common.NewSelectControlType(generated.NOT_CONTROLABLE)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	jsonData, err := json.Marshal(selectControlType)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var deserializedSelectControlType common.SelectControlType
	err = json.Unmarshal(jsonData, &deserializedSelectControlType)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if deserializedSelectControlType.ControlType != generated.NOT_CONTROLABLE {
		t.Errorf("expected control type %v, got %v", generated.NOT_CONTROLABLE, deserializedSelectControlType.ControlType)
	}
}

func TestSessionRequestSerialization(t *testing.T) {
	requestType := generated.RECONNECT
	diagnosticLabel := "Diagnostic Label"
	sessionRequest, err := common.NewSessionRequest(requestType, diagnosticLabel)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	jsonData, err := json.Marshal(sessionRequest)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var deserializedSessionRequest common.SessionRequest
	err = json.Unmarshal(jsonData, &deserializedSessionRequest)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if deserializedSessionRequest.Request != requestType {
		t.Errorf("expected request type %v, got %v", requestType, deserializedSessionRequest.Request)
	}
	if deserializedSessionRequest.DiagnosticLabel == nil || *deserializedSessionRequest.DiagnosticLabel != diagnosticLabel {
		t.Errorf("expected diagnostic label %v, got %v", diagnosticLabel, deserializedSessionRequest.DiagnosticLabel)
	}
}

