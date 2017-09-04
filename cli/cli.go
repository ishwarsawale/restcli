package cli

import (
  "os"
  "github.com/urfave/cli"
  "restcli/commands"
)

func CliApp() {
    app := cli.NewApp()
    app.Commands = commands.REST_Commands
    app.Flags = commands.Puppet_Flags
    app.Action = func(c *cli.Context) error {
    return nil
  }
    app.Run(os.Args)
}
