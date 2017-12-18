package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"checker"
	"checker/config"
	"checker/storage/postgres"
)

func main() {
	flagSet := flag.NewFlagSet("health check stats", flag.ExitOnError)
	dbConfig := config.NewDbConfigFlagSet(flagSet)
	sConfig := config.NewResourceStatsFlagSet(flagSet)

	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		flagSet.Usage()
		os.Exit(1)
	}

	flagSet.Parse(os.Args[1:])
	store, err := postgres.NewStore(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	service := checker.NewService(store)

	if sConfig.Url() == "" {
		stats, err := service.GetResourcesStats(sConfig.Start(), sConfig.End())
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(stats.String())
	} else {
		result, err := service.GetResourceStats(sConfig.Url(), sConfig.Start(), sConfig.End())
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(result.String())
	}
}
