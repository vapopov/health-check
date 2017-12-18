package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	"checker"
	"checker/config"
	"checker/storage/postgres"
)

func main() {
	flagSet := flag.NewFlagSet("web-resources health checker service", flag.ExitOnError)
	dbConfig := config.NewDbConfigFlagSet(flagSet)
	chConfig := config.NewCheckerFlagSet(flagSet)

	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		flagSet.Usage()
		os.Exit(1)
	}

	flagSet.Parse(os.Args[1:])

	store, err := postgres.NewStore(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	signalInterrupt := make(chan os.Signal)
	signal.Notify(signalInterrupt, os.Interrupt)
	go func() {
		select {
		case <-signalInterrupt:
			cancelFunc()
		}
	}()

	service := checker.NewService(store)
	service.Watch(ctx, chConfig)
}
