package cliargs

import (
	"fmt"
	"os"

	"github.com/bornone/clair-connector/version"
	"github.com/codegangsta/cli"
)

func Run(name string, commands ...cli.Command) {
	app := cli.NewApp()
	app.Name = name

	versionCommand := cli.Command{
		Name:      "version",
		ShortName: "v",
		Usage:     "Show the version",
		Action: func(_ *cli.Context) {
			fmt.Printf("%s-%s\n", version.VERSION, version.REVISION)
		},
	}

	app.Commands = append(commands, versionCommand)
	app.Version = fmt.Sprintf("%s-%s", version.VERSION, version.REVISION)
	app.Run(os.Args)
}
