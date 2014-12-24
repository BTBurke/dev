package main

import (
	"github.com/codegangsta/cli"
	"os"
)

const VERSION string = "0.1.0"

type Config struct {
	Port         string
	ContainerDir string
	LocalDir     string
	DevImage     string
	Shell        string
}

var DefaultConfig = Config{
	Port:     "10001:10001",
	DevImage: "btburke/golang-dev",
	Shell:    "/bin/bash",
}

func main() {
	app := cli.NewApp()
	app.Name = "dev"
	app.Version = VERSION
	app.Author = "Bryan Burke"
	app.Email = "bryan@alliedcodes.com"
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "create a stub fig.yml",
			Action: func(c *cli.Context) {
				InitNewWorkspace()
			},
		},
		{
			Name:  "shell",
			Usage: "start a new terminal session in the dev container",
			Action: func(c *cli.Context) {
				Shell()
			},
		},
		{
			Name:  "dep",
			Usage: "install necessary dependencies inside the dev container",
			Action: func(c *cli.Context) {
				Dep()
			},
		},
		{
			Name:  "up",
			Usage: "start your dev environment",
			Action: func(c *cli.Context) {
				Up()
			},
		},
		{
			Name:  "stop",
			Usage: "stop your dev-managed environment",
			Action: func(c *cli.Context) {
				Stop()
			},
		},
		{
			Name:  "rm",
			Usage: "remove stopped dev-managed containers",
			Action: func(c *cli.Context) {
				Rm()
			},
		},
		{
			Name:  "web",
			Usage: "link to the Goconvey web interface",
			Action: func(c *cli.Context) {
				Web()
			},
		},
		{
			Name:        "test",
			Usage:       "run go test in the dev container",
			Description: "By default, the command runs go test -v.  To run more sophisticated go test commands with other flags, run dev shell and then execute go tests directly in the container.",
			Action: func(c *cli.Context) {
				Test()
			},
		},
	}
	app.Run(os.Args)
}
