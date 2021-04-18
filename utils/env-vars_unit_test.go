// +build !integration

package utils_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-postgres-cicd/utils"
)

const (
	stringKey      = "DB_EXAMPLE_STRING"
	stringValue    = "sample"
	stringFallback = "fallback"

	intKey      = "DB_EXAMPLE_INT"
	intValue    = 10
	intFallback = 42

	boolKey      = "DB_EXAMPLE_BOOL"
	boolValue    = true
	boolFallback = false
)

func TestGetStringEnv_Success(t *testing.T) {
	setErr := os.Setenv(stringKey, stringValue)
	require.NoError(t, setErr)

	value := utils.GetStringEnv(stringKey, stringFallback)

	assert.Equal(t, stringValue, value)

	unsetErr := os.Unsetenv(stringKey)
	require.NoError(t, unsetErr)
}

func TestGetStringEnv_Fallback(t *testing.T) {
	value := utils.GetStringEnv(stringKey, stringFallback)

	assert.NotEqual(t, stringValue, value)
	assert.Equal(t, stringFallback, value)
}

func TestGetIntEnv_Success(t *testing.T) {
	setErr := os.Setenv(intKey, strconv.Itoa(intValue))
	require.NoError(t, setErr)

	value := utils.GetIntEnv(intKey, intFallback)

	assert.Equal(t, intValue, value)

	unsetErr := os.Unsetenv(intKey)
	require.NoError(t, unsetErr)
}

func TestGetIntEnv_Fallback_NotSet(t *testing.T) {
	value := utils.GetIntEnv(intKey, intFallback)

	assert.NotEqual(t, intValue, value)
	assert.Equal(t, intFallback, value)
}

func TestGetIntEnv_Fallback_Format(t *testing.T) {
	setErr := os.Setenv(intKey, "42s")
	require.NoError(t, setErr)

	value := utils.GetIntEnv(intKey, intFallback)

	assert.NotEqual(t, intValue, value)
	assert.Equal(t, intFallback, value)

	unsetErr := os.Unsetenv(intKey)
	require.NoError(t, unsetErr)
}

func TestGetBoolEnv_Success(t *testing.T) {
	setErr := os.Setenv(boolKey, strconv.FormatBool(boolValue))
	require.NoError(t, setErr)

	value := utils.GetBoolEnv(boolKey, boolFallback)

	assert.Equal(t, boolValue, value)

	unsetErr := os.Unsetenv(boolKey)
	require.NoError(t, unsetErr)
}

func TestGetBoolEnv_Fallback_NotSet(t *testing.T) {
	value := utils.GetBoolEnv(boolKey, boolFallback)

	assert.NotEqual(t, boolValue, value)
	assert.Equal(t, boolFallback, value)
}

func TestGetBoolEnv_Fallback_Format(t *testing.T) {
	setErr := os.Setenv(boolKey, "unknown")
	require.NoError(t, setErr)

	value := utils.GetBoolEnv(boolKey, boolFallback)

	assert.NotEqual(t, boolValue, value)
	assert.Equal(t, boolFallback, value)

	unsetErr := os.Unsetenv(boolKey)
	require.NoError(t, unsetErr)
}
