package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/TarsTestToolKit/ResFetcher/impl"
	fetchertars "github.com/TarsTestToolKit/ResFetcher/tars-protocol/ResFetcher"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(impl.FetcherImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("fetcherImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(fetchertars.Fetcher)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".fetcherObj")

	// Run application
	tars.Run()
}
