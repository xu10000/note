package main

import (
	"os"
	"os/exec"
	"fmt"
	"github.com/urfave/cli"


)

func main() {
	app := cli.NewApp()
	app.Name = "xzt"
	app.Usage = "my personal tool"
	app.Commands = []cli.Command{
		{
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
					fmt.Println("bash err: ", string(shCmd))
				}
				return nil
			},
		},
	}
	
  
	app.Run(os.Args)
  }