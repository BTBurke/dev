package main

import (
	"fmt"
	//"github.com/codegangsta/cli"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

const figTmpl string = `dev:
    image: {{.DevImage}}
    ports:
        - "{{.Port}}"
    volumes_from:
        - code
    working_dir: {{.ContainerDir}}
code:
    image: busybox
    volumes:
        - {{.LocalDir}}:{{.ContainerDir}}`

func populateConfig() (*Config, error) {
	config := DefaultConfig
	var err error
	config.LocalDir, err = os.Getwd()
	if err != nil {
		return &config, fmt.Errorf("Error: Could not get current path.")
	}

	localGopath := os.Getenv("GOPATH")
	if len(localGopath) == 0 {
		return &config, fmt.Errorf("Error: GOPATH not set.")
	}

	relPath, err := filepath.Rel(localGopath, config.LocalDir)
	if err != nil {
		return &config, fmt.Errorf("Error: Current working directory is not a subdirectory of the GOPATH.")
	}

	config.ContainerDir = filepath.Join("/golang/", relPath)
	return &config, nil
}

func renderToFig(c *Config, wr io.Writer) error {
	t := template.Must(template.New("fig").Parse(figTmpl))
	err := t.Execute(wr, c)
	if err != nil {
		return fmt.Errorf("Error: Could not write configuration to fig.yml.")
	}
	return nil
}

func InitNewWorkspace() {
	if _, err := os.Stat("fig.yml"); err == nil {
		log.Fatal("Error: fig.yml already exists.  If you want to re-init this project, delete fig.yml first.")
	}

	f, err := os.Create("fig.yml")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	config, err := populateConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = renderToFig(config, f)
	if err != nil {
		log.Fatal("Error: Could not write to fig.yml")
	}

}
