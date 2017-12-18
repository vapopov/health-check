package config

import (
	"flag"
	"fmt"
)

// DSNConfig default config dependencies for database connection
type DSNConfig struct {
	host     *string
	user     *string
	port     *string
	password *string
	database *string
}

// NewDbConfigFlagSet returns flagset and attached dsn configuration
func NewDbConfigFlagSet(flagSet *flag.FlagSet) *DSNConfig {
	config := &DSNConfig{
		host:     flagSet.String("host", "localhost", "Database Host"),
		user:     flagSet.String("user", "postgres", "Database Username"),
		port:     flagSet.String("port", "5432", "Database Username"),
		password: flagSet.String("password", "", "Database Password"),
		database: flagSet.String("database", "postgres", "Database Name"),
	}

	return config
}

// DbName returns database name
func (d *DSNConfig) DbName() string {
	return *d.database
}

// String returns the string representation (DSN) of the configuration
func (d *DSNConfig) String() string {
	return fmt.Sprintf("user='%s' dbname='%s' password='%s' host='%s' port='%s' sslmode='disable'",
		*d.user, *d.database, *d.password, *d.host, *d.port)
}
