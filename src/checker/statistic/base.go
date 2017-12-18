package statistic

import (
	"fmt"
	"time"

	"checker/storage"
)

// BaseStats struct of availability information for each resource in percentage
type BaseStats struct {
	resourcesInfo map[string]*resourceInfo
}

// NewBaseStats constructor of BaseStats
func NewBaseStats() *BaseStats {
	return &BaseStats{
		resourcesInfo: make(map[string]*resourceInfo),
	}
}

// CalculateLogs calculates logs for all resources
func (s *BaseStats) CalculateLogs(sLogs []*storage.Log) error {
	if len(sLogs) == 0 {
		return nil
	}

	for _, sLog := range sLogs {
		info, ok := s.resourcesInfo[sLog.Url]
		if !ok {
			info = &resourceInfo{Url: sLog.Url}
			s.resourcesInfo[sLog.Url] = info
		}

		if err := info.CalculateLog(sLog); err != nil {
			return err
		}
	}

	return nil
}

// String formats base output of statistic for all resources
func (s *BaseStats) String() string {
	var result string
	for _, info := range s.resourcesInfo {
		result += info.String() + "\n"
	}

	return result
}

type resourceInfo struct {
	Url     string
	prevLog *storage.Log
	active  time.Duration
	allTime time.Duration
}

// CalculateLog calculates base statistic for resource
func (s *resourceInfo) CalculateLog(sLog *storage.Log) error {
	if s.prevLog == nil {
		s.prevLog = sLog
		return nil
	}

	diffDuration := sLog.CreatedAt.Sub(s.prevLog.CreatedAt)
	if diffDuration < 0 {
		return fmt.Errorf(
			"logs must be sorted by ascending, prev: %v, next: %v",
			s.prevLog.CreatedAt,
			sLog.CreatedAt,
		)
	}

	if s.prevLog.IsHealthy {
		s.active += diffDuration
	}
	s.allTime += diffDuration
	s.prevLog = sLog

	return nil
}

// Percentage returns percent of resource availability
func (s *resourceInfo) Percentage() float64 {
	if s.allTime == 0 {
		return 0
	}

	return (s.active.Seconds() / s.allTime.Seconds()) * 100
}

// String formats base output of statistic for resource
func (s *resourceInfo) String() string {
	return fmt.Sprintf("resource: %s available %.2f%% of time", s.Url, s.Percentage())
}
