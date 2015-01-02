package main

import (
	"encoding/json"
	"github.com/BTBurke/clt"
	"log"
	"os"
)

type depError struct {
	ImportStack []string
}

type depSummary struct {
	ImportPath string
	Incomplete bool
	DepsErrors []depError
}

func installDep(dep string) error {
	_, err := RunInDevContainerCapture([]string{"go", "get", dep})
	if err != nil {
		return err
	}
	return nil
}

func installDepsUsingGoGet() {
	out, err := RunInDevContainerCapture([]string{"go", "list", "-e", "-json"})
	if err != nil {
		log.Fatalf("Error: Could not get list of dependencies. %v", err)
	}
	var summary depSummary
	err = json.Unmarshal(out, &summary)
	if err != nil {
		log.Fatalf("Error: Could not get dependencies. %v", err)
	}
	clt.Say("Installing dependencies using go get")
	for _, dep := range summary.DepsErrors {
		for _, stack := range dep.ImportStack {
			switch stack {
			case summary.ImportPath:
				continue
			default:
				p := clt.NewProgressSpinner("%v", stack)
				p.Start()
				err = installDep(stack)
				if err != nil {
					p.Fail()
					log.Fatalf("Error: Failed to install dependency %v", stack)
				}
				p.Success()
			}
		}
	}
}

func installDepsUsingGodep() {
	err := installDep("github.com/tools/godep")
	if err != nil {
		log.Fatalf("Error: Failed to install godep")
	}
	clt.Say("Installing dependencies using Godep restore")
	p := clt.NewProgressSpinner("Godep restore")
	p.Start()
	_, err = RunInDevContainerCapture([]string{"godep", "restore"})
	if err != nil {
		p.Fail()
		log.Fatalf("Error: Godep restore did not complete successfully. %v", err)
	}
	p.Success()
}

func Dep() {
	_, err := os.Stat("Godeps/Godeps.json")
	switch os.IsNotExist(err) {
	case true:
		installDepsUsingGoGet()
	case false:
		installDepsUsingGodep()
	}
}
