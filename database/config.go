package database

import (
	"github.com/bygui86/go-postgres-cicd/logging"
	"github.com/bygui86/go-postgres-cicd/utils"
)

const (
	dbHostEnvVar     = "DB_HOST"
	dbPortEnvVar     = "DB_PORT"
	dbUsernameEnvVar = "DB_USERNAME"
	dbPasswordEnvVar = "DB_PASSWORD"
	dbNameEnvVar     = "DB_NAME"
	dbSslModeEnvVar  = "DB_SSL_MODE"

	dbHostDefault     = "localhost"
	dbPortDefault     = 5432
	dbUsernameDefault = "username"
	dbPasswordDefault = "password"
	dbNameDefault     = "db"
	dbSslModeDefault  = "disable"
)

func LoadConfig() *config {
	logging.Log.Debug("Load REST configurations")

	return &config{
		dbHost:     utils.GetStringEnv(dbHostEnvVar, dbHostDefault),
		dbPort:     utils.GetIntEnv(dbPortEnvVar, dbPortDefault),
		dbUsername: utils.GetStringEnv(dbUsernameEnvVar, dbUsernameDefault),
		dbPassword: utils.GetStringEnv(dbPasswordEnvVar, dbPasswordDefault),
		dbName:     utils.GetStringEnv(dbNameEnvVar, dbNameDefault),
		dbSslMode:  utils.GetStringEnv(dbSslModeEnvVar, dbSslModeDefault),
	}
}
