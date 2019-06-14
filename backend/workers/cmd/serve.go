package cmd

import (
	"github.com/jaxmef/booklover-backend/config"
	"github.com/jaxmef/booklover-backend/initialization"
	"github.com/jaxmef/booklover-backend/workers"
	"github.com/urfave/cli"
)

var serveCommand = cli.Command{
	Name:   "serve",
	Usage:  "starts " + config.ServiceName + " workers",
	Action: serveAction,
}

func serveAction(c *cli.Context) error {
	cfg := initialization.Init(c)

	workers.GetChief().RunAll(cfg.Log.AppName, cfg.Workers...)
	return nil
}
