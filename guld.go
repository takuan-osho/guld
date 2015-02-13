package main

import (
	"os"

	"github.com/codegangsta/cli"
)

const (
	BaseDirBaseName = "guld"
	BaseDir         = "." + BaseDirBaseName
)

func main() {
	app := cli.NewApp()
	app.Name = "guld"
	app.Usage = "Use it as you use git"
	app.Commands = []cli.Command{
		{
			Name:        "init",
			Usage:       "Create an empty guld repository",
			Description: "Create an empty guld repository",
			Action: func(c *cli.Context) {
				Init()
			},
		},
		{
			Name:        "hash-object",
			Usage:       "Computes the object ID value for an object",
			Description: "Computes the object ID value for an object with specified type with the contents of the named",
			Flags:       CmdHashObjectFlags,
			Action: func(c *cli.Context) {
				HashObjectAction(c)
			},
		},
	}
	app.Run(os.Args)
}
