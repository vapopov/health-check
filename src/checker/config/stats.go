package config

import (
	"flag"
	"fmt"
	"log"
	"time"
)

const timeLayout = "02-01-2006"

// StatsResourceConfig config parameters fo detail resource info
type StatsResourceConfig struct {
	url   *string
	start *string
	end   *string
}

// NewResourceStatsFlagSet returns flagset to health check configuration
func NewResourceStatsFlagSet(flagSet *flag.FlagSet) *StatsResourceConfig {
	config := &StatsResourceConfig{
		url:   flagSet.String("url", "", "url to resource"),
		start: flagSet.String("start", "", fmt.Sprintf("start date, in format %s", timeLayout)),
		end:   flagSet.String("end", "", fmt.Sprintf("end date, in format %s", timeLayout)),
	}

	return config
}

// Url store url to checking url
func (s *StatsResourceConfig) Url() string {
	return *s.url
}

// Start time of searching period
func (s *StatsResourceConfig) Start() time.Time {
	start, err := time.Parse(timeLayout, *s.start)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return start
}

// End time of searching period
func (s *StatsResourceConfig) End() time.Time {
	end, err := time.Parse(timeLayout, *s.end)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return end
}
