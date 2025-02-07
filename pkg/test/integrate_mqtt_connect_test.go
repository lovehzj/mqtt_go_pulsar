package test

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestMqttConnect steps:
// - create pulsar
// - create mqsar
// - create mqtt client
// - mqtt client connect
func TestMqttConnect(t *testing.T) {
	setupPulsar()
	_, port := setupMqsar()
	ops := mqtt.NewClientOptions().SetUsername("username").SetClientID("foo").AddBroker(MqttConnAddr(port))
	mqttCli := mqtt.NewClient(ops)
	token := mqttCli.Connect()
	token.Wait()
	assert.True(t, mqttCli.IsConnected())
}

// TestMqttUnConnected steps:
// - create pulsar
// - create mqsar
// - create mqtt client
// - mqtt client connect
// - mqtt client check connection
func TestMqttUnConnected(t *testing.T) {
	setupPulsar()
	_, port := setupMqsar()
	ops := mqtt.NewClientOptions().SetUsername("username").SetPassword("wrong_password").SetClientID("foo").AddBroker(MqttConnAddr(port))
	mqttCli := mqtt.NewClient(ops)
	token := mqttCli.Connect()
	token.Wait()
	assert.True(t, !mqttCli.IsConnected())
}
