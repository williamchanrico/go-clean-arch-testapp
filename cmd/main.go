package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/williamchanrico/xtest/cmd/xtest"
	"github.com/williamchanrico/xtest/log"
)

var (
	version            string
	showVersionAndExit bool
)

func main() {
	f := xtest.Flags{}
	flag.StringVar(&f.HTTPAddress, "http-address", "0.0.0.0:9000", "HTTP API listener")
	flag.StringVar(&f.RedisAddress, "redis-address", "127.0.0.1:6379", "Redis address")
	flag.StringVar(&f.PostgresDSN, "postgres-dsn", "postgres://postgres:@127.0.0.1/postgres?sslmode=disable", "PostgreSQL DB DSN")
	flag.StringVar(&f.NSQDAddress, "nsqd-address", "127.0.0.1:4150", "NSQd address")
	flag.StringVar(&f.LogLevel, "log-level", "info", "App log level")

	flag.BoolVar(&showVersionAndExit, "version", false, "Show version and exit")
	flag.Parse()

	if showVersionAndExit {
		fmt.Printf("xtest %v\n", version)
		os.Exit(0)
	}

	log.SetLevelString(f.LogLevel)

	exitCode, err := xtest.Run(f)
	if err != nil {
		log.Error(err)
	}

	if exitCode != 0 {
		log.Errorf("xtest exited with code %d", exitCode)
	}
	os.Exit(exitCode)
}
