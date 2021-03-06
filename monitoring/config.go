package monitoring

import (
	"github.com/bygui86/go-postgres-cicd/logging"
	"github.com/bygui86/go-postgres-cicd/utils"
)

const (
	monitorHostEnvVar = "MONITOR_HOST"
	monitorPortEnvVar = "MONITOR_PORT"

	monitorHostDefault = "0.0.0.0"
	monitorPortDefault = 9090
)

func loadConfig() *Config {
	logging.Log.Debug("Load monitoring configurations")
	return &Config{
		restHost: utils.GetStringEnv(monitorHostEnvVar, monitorHostDefault),
		restPort: utils.GetIntEnv(monitorPortEnvVar, monitorPortDefault),
	}
}
