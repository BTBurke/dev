package main

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"log"
)

func Web() {
	arch := DetermineDockerType()
	switch arch {
	case UseBoot2Docker:
		open.Run("http://192.168.59.103:10001")
	case UseDocker:
		open.Run("http://127.0.0.1:10001")
	}
}

func Test() {
	out, err := RunInDevContainerCapture([]string{"go", "test", "-v"})
	if err != nil {
		log.Fatalf("Error: Failed to run go test in dev container. %v", err)
	}
	fmt.Printf("%s", string(out[:]))
}
