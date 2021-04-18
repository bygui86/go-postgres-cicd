package config

func (c *Config) EnableMonitoring() bool {
	return c.enableMonitoring
}

func (c *Config) EnableTracing() bool {
	return c.enableTracing
}

func (c *Config) TracingTech() string {
	return c.tracingTech
}
