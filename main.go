package main

import (
	//"fmt"
	"github.com/codegangsta/cli"
	//dclient "github.com/fsouza/go-dockerclient"
	//"log"
	"os"
)

const VERSION string = "0.1.0"

type Config struct {
	Port         string
	ContainerDir string
	LocalDir     string
	DevImage     string
}

var DefaultConfig = Config{
	Port:     "10001:10001",
	DevImage: "btburke/golang-dev",
}

// var Client *dclient.Client

// func init() {
// 	Client, err := NewDockerClient()
// 	if err != nil {
// 		log.Fatal("Error: Could not get a connection to the Docker API.  If you're using boot2docker, check the status of the VM.")
// 	}
// }

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
	}
	app.Run(os.Args)
}
