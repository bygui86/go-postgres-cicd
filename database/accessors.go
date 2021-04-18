package database

func (c *config) DbHost() string {
	return c.dbHost
}

func (c *config) DbPort() int {
	return c.dbPort
}

func (c *config) DbUsername() string {
	return c.dbUsername
}

func (c *config) DbPassword() string {
	return c.dbPassword
}

func (c *config) DbName() string {
	return c.dbName
}

func (c *config) DbSslMode() string {
	return c.dbSslMode
}

// DEFAULTS

func DbHostDefault() string {
	return dbHostDefault
}

func DbPortDefault() int {
	return dbPortDefault
}

func DbUsernameDefault() string {
	return dbUsernameDefault
}

func DbPasswordDefault() string {
	return dbPasswordDefault
}

func DbNameDefault() string {
	return dbNameDefault
}

func DbSslModeDefault() string {
	return dbSslModeDefault
}
