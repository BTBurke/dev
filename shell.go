package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func RunInDevContainerInteractive(execArgs []string) {
	docker, err := exec.LookPath("docker")
	if err != nil {
		log.Fatalf("Error: Docker does not appear to be installed.")
	}

	container, err := FindDevContainer("dev")
	if err != nil {
		log.Fatal(err)
	}
	argsBase := []string{"docker", "exec", "-it", container.ID}
	argsAll := append(argsBase, execArgs...)
	env := os.Environ()

	execErr := syscall.Exec(docker, argsAll, env)
	if execErr != nil {
		log.Fatalf("Error: Failed while executing command in container. %v", execErr)
	}
}

func RunInDevContainerCapture(execArgs []string) (out []byte, execErr error) {
	docker, err := exec.LookPath("docker")
	if err != nil {
		log.Fatalf("Error: Docker does not appear to be installed.")
	}

	container, err := FindDevContainer("dev")
	if err != nil {
		log.Fatal(err)
	}
	argsBase := []string{"exec", "-it", container.ID}
	argsAll := append(argsBase, execArgs...)
	out, execErr = exec.Command(docker, argsAll...).Output()
	return
}

func Shell() {
	RunInDevContainerInteractive([]string{"/bin/bash"})
}
