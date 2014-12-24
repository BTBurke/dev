package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
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
	_, err = RunInDevContainerCapture([]string{"go", "get", dep})
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
	for _, dep := range summary.DepsErrors {
		for _, stack := range dep.ImportStack {
			switch stack {
			case summary.ImportPath:
				continue
			default:
				err = installDep(stack)
				if err != nil {
					log.Fatalf("Error: Failed to install dependency %v", stack)
				}
				fmt.Printf("Installed: %v\n", stack)
			}
		}
	}
}

func installDepsUsingGodep() {
	err := installDep("github.com/tools/godep")
	if err != nil {
		log.Fatalf("Error: Failed to install godep")
	}
	_, err = RunInDevContainerCapture([]string{"godep", "restore"})
	if err != nil {
		log.Fatalf("Error: Godep restore did not complete successfully. %v", err)
	}
}

func Dep() {
	installDepsUsingGoGet()
}
