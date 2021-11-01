package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstance(t *testing.T) {
	configFilePath = "../config.yml"
	config := Instance()

	config.Server.Port.Http = "100500"

	assert.NotEqual(t, config.Server.Port.Http, Instance().Server.Port.Http)
}
