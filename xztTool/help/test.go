package help

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli"
)

var TestCommand = cli.Command{
	// hello world
	Name:        "hello",
	Aliases:     []string{"hi"},
	Usage:       "I don't the usage.. so, just say it",
	Description: "This is how we describe hello the function",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: "Bob",
			Usage: "Name of the person to greet",
		},
	},
	Action: func(c *cli.Context) error {
		name := c.String("name")
		shCmd := "echo hello world " + name
		fmt.Printf("begin: %s\n", shCmd)

		cmd := exec.Command("sh", "-c", shCmd)
		_, err := cmd.CombinedOutput()
		if err != nil {
			panic("bash err: " + string(shCmd))
		}
		return nil
	},
}
