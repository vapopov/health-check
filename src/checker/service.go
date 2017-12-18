package checker

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"checker/config"
	"checker/statistic"
	"checker/storage"
)

const (
	parallelRunMaxQty = 2000
)

// Service implements resources availability checking and preparing base statistics
type Service struct {
	store storage.Store
}

// NewService constructor of service
func NewService(store storage.Store) *Service {
	return &Service{
		store: store,
	}
}

// Watch runs process of checking resources availabilities taken from list of resources in file and by specified interval
func (s *Service) Watch(ctx context.Context, chConfig *config.CheckerConfig) {
	sourcesFile, err := os.Open(chConfig.Source())
	if err != nil {
		log.Fatal(err)
	}
	defer sourcesFile.Close()

	var parallelRun int
	scanner := bufio.NewScanner(sourcesFile)
	for scanner.Scan() && parallelRun < parallelRunMaxQty {
		s.spawnCheck(ctx, scanner.Text(), chConfig.Interval())
		parallelRun++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	select {
	case <-ctx.Done():
		fmt.Println("Service shutdown.")
		return
	}
}

// GetResourceStats calculates detail stats for resource like:
// percentage of availability, and longest period when resource was offline
func (s *Service) GetResourceStats(url string, start, end time.Time) (*statistic.DetailStats, error) {
	result := statistic.NewDetailStats()

	sLogs, err := s.store.FetchLogs(start, end, url)
	if err != nil {
		return nil, err
	}

	for len(sLogs) > 0 {
		err = result.CalculateLogs(sLogs)
		if err != nil {
			return nil, err
		}

		start = sLogs[len(sLogs)-1].CreatedAt

		sLogs, err = s.store.FetchLogs(start, end, url)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// GetResourcesStats calculates base stats for all resource in specified time range
func (s *Service) GetResourcesStats(start, end time.Time) (*statistic.BaseStats, error) {
	result := statistic.NewBaseStats()

	sLogs, err := s.store.FetchLogs(start, end, "")
	if err != nil {
		return nil, err
	}

	for len(sLogs) > 0 {
		err = result.CalculateLogs(sLogs)
		if err != nil {
			return nil, err
		}

		start = sLogs[len(sLogs)-1].CreatedAt

		sLogs, err = s.store.FetchLogs(start, end, "")
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (s *Service) spawnCheck(ctx context.Context, url string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case requestTime := <-ticker.C:
				s.store.AddLog(&storage.Log{CreatedAt: requestTime, Url: url, IsHealthy: s.checkResource(url)})
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}

func (s *Service) checkResource(url string) bool {
	transport := &http.Transport{
		MaxIdleConns:    1,
		IdleConnTimeout: 10 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 0,
			DualStack: true,
		}).DialContext,
	}

	client := &http.Client{Transport: transport}
	resp, err := client.Head(url)

	return err == nil && resp.StatusCode < 400
}
