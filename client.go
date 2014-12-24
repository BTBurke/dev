package main

import (
	"fmt"
	dclient "github.com/fsouza/go-dockerclient"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"
)

const (
	UseBoot2Docker = iota
	UseDocker
)

type tls struct {
	cert string
	key  string
	ca   string
}

func determineDockerType() int {
	_, err := exec.LookPath("boot2docker")
	if err != nil {
		return UseDocker
	}
	return UseBoot2Docker
}

func NewDockerClient() (*dclient.Client, error) {
	arch := determineDockerType()
	switch arch {
	case UseBoot2Docker:
		client, err := newBoot2DockerClient()
		if err != nil {
			return nil, err
		}
		return client, nil
	case UseDocker:
		log.Fatalf("Sorry, dev is not ready to be used on Linux systems.  Stay tuned.")
	}
	return nil, nil
}

func newBoot2DockerClient() (*dclient.Client, error) {
	ip := os.Getenv("DOCKER_HOST")
	if len(ip) == 0 {
		ip = "tcp://192.168.59.103:2376"
	}
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	useTls := tls{
		cert: path.Join(usr.HomeDir, "/.boot2docker/certs/boot2docker-vm/cert.pem"),
		key:  path.Join(usr.HomeDir, "/.boot2docker/certs/boot2docker-vm/key.pem"),
		ca:   path.Join(usr.HomeDir, "/.boot2docker/certs/boot2docker-vm/ca.pem"),
	}
	if _, err := os.Stat(useTls.ca); os.IsNotExist(err) {
		return nil, fmt.Errorf("Error: Could not find credentials for TLS connection to Docker.")
	}
	if _, err := os.Stat(useTls.cert); os.IsNotExist(err) {
		return nil, fmt.Errorf("Error: Could not find credentials for TLS connection to Docker.")
	}
	if _, err := os.Stat(useTls.key); os.IsNotExist(err) {
		return nil, fmt.Errorf("Error: Could not find credentials for TLS connection to Docker.")
	}
	client, err := dclient.NewTLSClient(ip, useTls.cert, useTls.key, useTls.ca)
	if err != nil {
		return nil, err
	}
	return client, nil

}

func FindDevContainer(figName string) (dclient.APIContainers, error) {
	client, err := NewDockerClient()
	if err != nil {
		return dclient.APIContainers{}, fmt.Errorf("Error: No connection to Docker. %v", err)
	}
	opts := dclient.ListContainersOptions{}
	containers, err := client.ListContainers(opts)
	if err != nil {
		return dclient.APIContainers{}, fmt.Errorf("Error: Could not get running containers. %v", err)
	}
	for _, container := range containers {
		for _, cName := range container.Names {
			if strings.HasSuffix(cName, figName+"_1") {
				return container, nil
			}
		}
	}
	return dclient.APIContainers{}, fmt.Errorf("Error: Cannot find your running development container.")
}
