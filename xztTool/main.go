package main

import (
	"os"
	"xzt/help"

	"github.com/urfave/cli"
)

func main() {
	// fmt.Println(help.X)
	app := cli.NewApp()
	app.Name = "xzt"
	app.Usage = "my personal tool"
	app.Commands = []cli.Command{
		help.TestCommand,
		help.TicCrud,
	}

	app.Run(os.Args)
}
