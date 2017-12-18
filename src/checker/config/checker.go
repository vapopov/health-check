package config

import (
	"flag"
	"time"
)

// CheckerConfig config
type CheckerConfig struct {
	source   *string
	interval *uint
}

// NewCheckerFlagSet returns flagset to health check configuration
func NewCheckerFlagSet(flagSet *flag.FlagSet) *CheckerConfig {
	config := &CheckerConfig{
		source:   flagSet.String("source", "", "path to file with sources"),
		interval: flagSet.Uint("interval", 30, "interval between check availability"),
	}

	return config
}

// Source file path to list with resources to check
func (c *CheckerConfig) Source() string {
	return *c.source
}

// Interval of checking availability of resources
func (c *CheckerConfig) Interval() time.Duration {
	return time.Duration(*c.interval) * time.Second
}
