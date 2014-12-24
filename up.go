package main

import (
	"fmt"
	"log"
	"os/exec"
)

func figCmd(cmd []string) (out []byte, execErr error) {
	fig, err := exec.LookPath("fig")
	if err != nil {
		log.Fatalf("Error: Fig does not appear to be installed.")
	}

	out, execErr = exec.Command(fig, cmd...).Output()
	return
}

func dockerPs() (out []byte, execErr error) {
	docker, err := exec.LookPath("docker")
	if err != nil {
		log.Fatalf("Error: Docker does not appear to be installed.")
	}

	argsBase := []string{"ps"}
	out, execErr = exec.Command(docker, argsBase...).Output()
	return
}

func Up() {
	_, err := figCmd([]string{"up", "-d"})
	if err != nil {
		log.Fatalf("Error: Fig up failed to start. %v", err)
	}
	out, err := dockerPs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updating dependencies...")
	Dep()
	fmt.Printf("\nYour dev environment is running:\n\n%s", string(out[:]))
}

func Stop() {
	_, err := figCmd([]string{"stop"})
	if err != nil {
		log.Fatalf("Error: Failed to stop running containers. %v", err)
	}
	fmt.Println("All dev-managed containers stopped.\n")
	out, err := dockerPs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nYour environment is now:\n\n%s", string(out[:]))
}

func Rm() {
	out, err := figCmd([]string{"rm", "-force"})
	if err != nil {
		log.Fatalf("Error: Failed to remove stopped containers. %v", err)
	}
	fmt.Printf("%s", string(out[:]))
}
