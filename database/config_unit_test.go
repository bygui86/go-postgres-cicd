// +build !integration

package database_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bygui86/go-postgres-cicd/database"
	"github.com/bygui86/go-postgres-cicd/logging"
)

const (
	hostKey   = "DB_HOST"
	hostValue = "postgresql.remote.test"

	portKey   = "DB_PORT"
	portValue = 5433

	userKey   = "DB_USERNAME"
	userValue = "john"

	pwKey   = "DB_PASSWORD"
	pwValue = "myPass"

	nameKey   = "DB_NAME"
	nameValue = "test"

	sslKey   = "DB_SSL_MODE"
	sslValue = "enable"
)

func TestLoadConfig(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	hostErr := os.Setenv(hostKey, hostValue)
	require.NoError(t, hostErr)
	portErr := os.Setenv(portKey, strconv.Itoa(portValue))
	require.NoError(t, portErr)
	userErr := os.Setenv(userKey, userValue)
	require.NoError(t, userErr)
	pwErr := os.Setenv(pwKey, pwValue)
	require.NoError(t, pwErr)
	nameErr := os.Setenv(nameKey, nameValue)
	require.NoError(t, nameErr)
	sslErr := os.Setenv(sslKey, sslValue)
	require.NoError(t, sslErr)

	cfg := database.LoadConfig()

	assert.Equal(t, hostValue, cfg.DbHost())
	assert.Equal(t, portValue, cfg.DbPort())
	assert.Equal(t, userValue, cfg.DbUsername())
	assert.Equal(t, pwValue, cfg.DbPassword())
	assert.Equal(t, nameValue, cfg.DbName())
	assert.Equal(t, sslValue, cfg.DbSslMode())

	err := os.Unsetenv(hostKey)
	require.NoError(t, err)
	err = os.Unsetenv(portKey)
	require.NoError(t, err)
	err = os.Unsetenv(userKey)
	require.NoError(t, err)
	err = os.Unsetenv(pwKey)
	require.NoError(t, err)
	err = os.Unsetenv(nameKey)
	require.NoError(t, err)
	err = os.Unsetenv(sslKey)
	require.NoError(t, err)
}

func TestLoadConfig_Defaults(t *testing.T) {
	logErr := logging.InitGlobalLogger()
	require.NoError(t, logErr)

	cfg := database.LoadConfig()

	assert.Equal(t, database.DbHostDefault(), cfg.DbHost())
	assert.Equal(t, database.DbPortDefault(), cfg.DbPort())
	assert.Equal(t, database.DbUsernameDefault(), cfg.DbUsername())
	assert.Equal(t, database.DbPasswordDefault(), cfg.DbPassword())
	assert.Equal(t, database.DbNameDefault(), cfg.DbName())
	assert.Equal(t, database.DbSslModeDefault(), cfg.DbSslMode())
}
