package tests

import (
	"testing"

	"github.com/BuriP/S2-Go/src/common"
	"github.com/BuriP/S2-Go/src/generated"
	"github.com/stretchr/testify/assert"
)

func TestNewHandshake(t *testing.T) {
	t.Parallel()

	handshake, err := common.NewHandshake("CEM", []string{"1.0", "2.0"})
	assert.NoError(t, err)
	assert.NotNil(t, handshake)
	assert.Equal(t, "Handshake", handshake.MessageType)
	assert.Equal(t, generated.CEM, *handshake.Role)
	assert.ElementsMatch(t, []string{"1.0", "2.0"}, *handshake.SupportedProtocolVersions)

	// Test RM role with mandatory supported protocol versions
	handshakeRM, err := common.NewHandshake("RM", []string{"1.0"})
	assert.NoError(t, err)
	assert.NotNil(t, handshakeRM)
	assert.Equal(t, "Handshake", handshakeRM.MessageType)
	assert.Equal(t, generated.EnergyManagementRole("RM"), *handshakeRM.Role)
	assert.ElementsMatch(t, []string{"1.0"}, *handshakeRM.SupportedProtocolVersions)

	// Test invalid role
	_, err = common.NewHandshake("INVALID", []string{"1.0"})
	assert.Error(t, err)
}

func TestNewVarHandshake(t *testing.T) {
	t.Parallel()

	roles := []string{"CEM", "RM"}
	protocols := [][]string{{"1.0"}, {"2.0", "3.0"}}
	handshakes, err := common.NewVarHandshake(roles, protocols)
	assert.NoError(t, err)
	assert.Len(t, handshakes, 2)

	assert.Equal(t, generated.CEM, *handshakes[0].Role)
	assert.ElementsMatch(t, []string{"1.0"}, *handshakes[0].SupportedProtocolVersions)

	assert.Equal(t, generated.RM, *handshakes[1].Role)
	assert.ElementsMatch(t, []string{"2.0", "3.0"}, *handshakes[1].SupportedProtocolVersions)
}

func TestSetMessageID(t *testing.T) {
	t.Parallel()

	handshake, _ := common.NewHandshake("CEM", []string{"1.0"})
	newID := "new1234567890abcdef1234567890abcdef"
	handshake.SetMessageID(newID)
	assert.Equal(t, newID, handshake.MessageID.Value)
}

func TestSetSupportedProtocolVersions(t *testing.T) {
	t.Parallel()

	handshake, _ := common.NewHandshake("CEM", []string{"1.0"})
	newProtocols := []string{"2.0", "3.0"}
	handshake.SetSupportedProtocolVersions(newProtocols)
	assert.ElementsMatch(t, newProtocols, *handshake.SupportedProtocolVersions)
}
