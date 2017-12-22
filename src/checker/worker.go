package checker

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"checker/storage"
)

type LogChannel chan storage.Log

// Worker checks availability of resources
type Worker struct {
	workerPool chan LogChannel
	logChannel LogChannel

	transport *http.Transport
	store     storage.Store
}

// NewWorker constructor of worker initialise base configurations of http transport
func NewWorker(store storage.Store, workerPool chan LogChannel) Worker {
	return Worker{
		workerPool: workerPool,
		logChannel: make(chan storage.Log),
		transport: &http.Transport{
			MaxIdleConns:    0,
			IdleConnTimeout: 30 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
		},
		store: store,
	}
}

// Start starts the worker loop
func (w *Worker) Start(ctx context.Context) {
	go func() {
		for {
			w.workerPool <- w.logChannel

			select {
			case sLog := <-w.logChannel:
				sLog.IsHealthy = w.checkResource(sLog.Url)

				if err := w.store.AddLog(&sLog); err != nil {
					log.Printf("couldn't add log to store: %s \n", err.Error())
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (w *Worker) checkResource(url string) bool {
	client := &http.Client{
		Transport: w.transport,
		Timeout:   10 * time.Second,
	}
	resp, err := client.Get(url)

	return err == nil && resp.StatusCode < 400
}
