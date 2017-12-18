package statistic

import (
	"fmt"
	"time"

	"checker/storage"
)

// DetailStats struct of resource availability information with longest offline period
type DetailStats struct {
	prevLog         *storage.Log
	maxDuration     time.Duration
	currentDuration time.Duration
	active          time.Duration
	allTime         time.Duration
	inactiveStart   time.Time
	inactiveCurrent time.Time
	inactiveEnd     time.Time
}

// NewDetailStats constructor of DetailStats
func NewDetailStats() *DetailStats {
	return &DetailStats{}
}

// CalculateLogs calculates up-time states and longest offline period
func (s *DetailStats) CalculateLogs(sLogs []*storage.Log) error {
	if len(sLogs) == 0 {
		return nil
	}

	if s.prevLog == nil && !sLogs[0].IsHealthy {
		s.inactiveCurrent = sLogs[0].CreatedAt
	}

	for _, sLog := range sLogs {
		if err := s.calculateLog(sLog); err != nil {
			return err
		}
	}

	return nil
}

func (s *DetailStats) calculateLog(sLog *storage.Log) error {
	if s.prevLog == nil {
		s.prevLog = sLog
		return nil
	}

	if s.prevLog.Url != sLog.Url {
		return fmt.Errorf(
			"wrong log entitny, urls must be equal, prev: %s, current %s",
			s.prevLog.Url,
			sLog.Url,
		)
	}

	diffDuration := sLog.CreatedAt.Sub(s.prevLog.CreatedAt)
	if diffDuration < 0 {
		return fmt.Errorf(
			"logs must be sorted by ascending, prev: %v, next: %v",
			s.prevLog.CreatedAt,
			sLog.CreatedAt,
		)
	}

	s.currentDuration += diffDuration

	if s.prevLog.IsHealthy != sLog.IsHealthy {
		if s.maxDuration < s.currentDuration {
			s.maxDuration = s.currentDuration
			s.inactiveStart = s.inactiveCurrent
			s.inactiveEnd = sLog.CreatedAt
		}

		if sLog.IsHealthy {
			s.currentDuration = diffDuration
		} else {
			s.currentDuration = 0
			s.inactiveCurrent = sLog.CreatedAt
		}
	}

	if s.prevLog.IsHealthy {
		s.active += diffDuration
	}
	s.allTime += diffDuration
	s.prevLog = sLog

	return nil
}

// Percentage returns percent of resource availability
func (s *DetailStats) Percentage() float64 {
	if s.allTime == 0 {
		return 0
	}

	return (s.active.Seconds() / s.allTime.Seconds()) * 100
}

// String formats base output of statistic
func (s *DetailStats) String() string {
	result := fmt.Sprintf("available %.2f%% of time", s.Percentage())

	if s.maxDuration != 0 {
		result += fmt.Sprintf(
			", bigest downtime is %v from %v to %v",
			s.maxDuration,
			s.inactiveStart,
			s.inactiveEnd,
		)
	}

	return result
}
