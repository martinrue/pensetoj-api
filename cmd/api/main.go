package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/matryer/way"

	"github.com/martinrue/pensetoj-api/api"
	"github.com/martinrue/pensetoj-api/logger"
	"github.com/martinrue/pensetoj-api/store"
)

const usage = `Pensetoj API ($SHA)

Usage:
  api --bind=<bind-address>
`

var bind = flag.String("bind", "", "")

func main() {
	printUsage := func() {
		fmt.Fprint(os.Stderr, strings.Replace(usage, "$SHA", api.Commit, -1))
	}

	flag.Usage = func() {
		printUsage()
		os.Exit(0)
	}

	flag.Parse()

	if *bind == "" {
		printUsage()
		os.Exit(1)
	}

	logger := logger.New("[pensetoj-api]")

	jsonStore := store.NewJSON("db.json", logger)

	if err := jsonStore.Load(); err != nil {
		logger.System("store: failed to load: %v", err)
		os.Exit(2)
	}

	server := &api.Server{
		Logger: logger,
		Router: way.NewRouter(),
		Store:  jsonStore,
	}

	logger.System("starting server: %s", *bind)

	if err := server.Start(*bind); err != nil {
		logger.System("server start failed: %s", err)
		os.Exit(3)
	}
}
