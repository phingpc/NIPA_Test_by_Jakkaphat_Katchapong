// internal/config/config_test.go
package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	config := LoadConfig()
	assert.Equal(t, "localhost", config.DBHost)
	assert.Equal(t, "5432", config.DBPort)
	assert.Equal(t, "nipa_test", config.DBUser)
	assert.Equal(t, "nipa_test123", config.DBPassword)
	assert.Equal(t, "nipa_test", config.DBName)
	assert.Equal(t, "8080", config.APIPORT)
}

