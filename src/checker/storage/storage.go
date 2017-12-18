package storage

import (
	"time"
)

// Log is representation of storage log entity
type Log struct {
	CreatedAt time.Time
	Url       string
	IsHealthy bool
}

// Store is storage interface
type Store interface {
	// AddLog adds to storage information about resource status
	AddLog(sLog *Log) error

	// FetchLogs retrieves logs from storage in defined range, url is optional
	FetchLogs(start, end time.Time, url string) ([]*Log, error)
}
