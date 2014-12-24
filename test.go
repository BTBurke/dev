package main

import (
	"github.com/skratchdot/open-golang/open"
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
