package main

import (
	"os"

	"github.com/jaxmef/booklover-backend/cmd"
	"github.com/jaxmef/booklover-backend/config"
	"github.com/jaxmef/booklover-backend/info"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "A " + config.ServiceName + " service"
	app.Version = info.App.Version
	app.Flags = cmd.GetFlags()
	app.Commands = cmd.GetCommands()
	app.Run(os.Args)
}
