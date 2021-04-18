package database_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/cenkalti/backoff"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/bygui86/go-postgres-cicd/database"
	"github.com/bygui86/go-postgres-cicd/logging"
)

const (
	postgresUser = "postgres"
	postgresPw   = "supersecret"
	postgresDb   = "postgres"
)

func TestMain(m *testing.M) {
	logErr := logging.InitGlobalLogger()
	if logErr != nil {
		panic(logErr) // Panic and fail
	}

	ctx := context.Background()

	postgres, contErr := startPostgres(ctx)
	if contErr != nil {
		panic(contErr) // Panic and fail since there is not much we can do if the container doesn't start
	}
	logging.Log.Info("PostgreSQL container running")
	defer stopPostgres(postgres, ctx)

	host, port := getHostAndPort(postgres, ctx)
	logging.SugaredLog.Infof("PostgreSQL container exposed as: %s:%s", host, port.Port())

	setEnvVars(host, port)

	os.Exit(
		m.Run(),
	)
}

func startPostgres(ctx context.Context) (testcontainers.Container, error) {
	logging.Log.Info("Start PostgreSQL")
	contReq := testcontainers.ContainerRequest{
		Image:        "postgres:13.1-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": postgresPw,
			// "POSTGRES_HOST_AUTH_METHOD":       "trust",
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections"),
	}

	postgres, contErr := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: contReq,
			Started:          true,
		},
	)
	return postgres, contErr
}

func getHostAndPort(postgres testcontainers.Container, ctx context.Context) (string, nat.Port) {
	expPorts, expErr := postgres.Ports(ctx)
	if expErr != nil {
		panic(expErr)
	}
	logging.Log.Debug("PostgreSQL exposed ports:")
	for k, v := range expPorts {
		logging.SugaredLog.Debugf("\t %s -> %v", k, v)
	}

	host, hostErr := postgres.Host(ctx)
	if hostErr != nil {
		panic(hostErr) // Panic and fail since there is not much we can do if we cannot figure out the container ip
	}

	port, portErr := postgres.MappedPort(ctx, "5432")
	if portErr != nil {
		panic(portErr) // Panic and fail since there is not much we can do if we cannot figure out the container port
	}
	return host, port
}

func setEnvVars(host string, port nat.Port) {
	envErr := os.Setenv("DB_HOST", host)
	if envErr != nil {
		panic(envErr) // Panic and fail since there is not much we can do if we cannot set environment variables
	}
	envErr = os.Setenv("DB_PORT", port.Port())
	if envErr != nil {
		panic(envErr) // Panic and fail since there is not much we can do if we cannot set environment variables
	}
	envErr = os.Setenv("DB_USERNAME", postgresUser)
	if envErr != nil {
		panic(envErr) // Panic and fail since there is not much we can do if we cannot set environment variables
	}
	envErr = os.Setenv("DB_PASSWORD", postgresPw)
	// envErr = os.Setenv("DB_PASSWORD", "")
	if envErr != nil {
		panic(envErr) // Panic and fail since there is not much we can do if we cannot set environment variables
	}
	envErr = os.Setenv("DB_NAME", postgresDb)
	if envErr != nil {
		panic(envErr) // Panic and fail since there is not much we can do if we cannot set environment variables
	}
}

func ping(db *sql.DB) error {
	// WARN: connection takes a bit time to be opened, golang application is so fast that the first ping could easily fail
	pingErr := backoff.Retry(
		func() error {
			err := db.Ping()
			if err != nil {
				logging.Log.Debug("PostgreSQL connection not ready, backing off...")
				return err
			}
			logging.Log.Debug("PostgreSQL connection ready")
			return nil
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 10),
	)
	return pingErr
}

func stopPostgres(postgres testcontainers.Container, ctx context.Context) {
	logging.Log.Info("PostgreSQL container stop")
	err := postgres.Terminate(ctx)
	if err != nil {
		logging.SugaredLog.Error("PostgreSQL container stop failed: %s", err.Error())
	}
}

func initConnAndTable(t *testing.T) *sql.DB {
	db, dbErr := database.New()
	require.NoError(t, dbErr)

	pingErr := ping(db)
	require.NoError(t, pingErr)

	initErr := database.InitDb(db)
	require.NoError(t, initErr)
	return db
}
