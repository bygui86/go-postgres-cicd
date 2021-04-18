// +build !integration

package config_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-postgres-cicd/config"
	"github.com/bygui86/go-postgres-cicd/logging"
)

const (
	monitorKey   = "ENABLE_MONITORING"
	monitorValue = false

	traceKey   = "ENABLE_TRACING"
	traceValue = false

	techKey   = "TRACING_TECH"
	techValue = "zipkin"
)

func TestLoadConfig(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	monitorErr := os.Setenv(monitorKey, strconv.FormatBool(monitorValue))
	require.NoError(t, monitorErr)
	traceErr := os.Setenv(traceKey, strconv.FormatBool(traceValue))
	require.NoError(t, traceErr)
	techErr := os.Setenv(techKey, techValue)
	require.NoError(t, techErr)

	cfg := config.LoadConfig()

	assert.Equal(t, monitorValue, cfg.EnableMonitoring())
	assert.Equal(t, traceValue, cfg.EnableTracing())
	assert.Equal(t, techValue, cfg.TracingTech())

	err := os.Unsetenv(monitorKey)
	require.NoError(t, err)
	err = os.Unsetenv(traceKey)
	require.NoError(t, err)
	err = os.Unsetenv(techKey)
	require.NoError(t, err)
}

func TestLoadConfig_Defaults(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := config.LoadConfig()

	assert.Equal(t, true, cfg.EnableMonitoring())
	assert.Equal(t, true, cfg.EnableTracing())
	assert.Equal(t, config.TracingTechJaeger, cfg.TracingTech())
}

func TestLoadConfig_TracingTechNotSupported(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	techErr := os.Setenv(techKey, "not-supported")
	require.NoError(t, techErr)

	cfg := config.LoadConfig()

	assert.Equal(t, true, cfg.EnableMonitoring())
	assert.Equal(t, true, cfg.EnableTracing())
	assert.Equal(t, config.TracingTechJaeger, cfg.TracingTech())

	err := os.Unsetenv(techKey)
	require.NoError(t, err)
}
