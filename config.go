package main

import (
	"encoding/json"
	"github.com/BTBurke/clt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

type Config struct {
	Port            string
	ContainerDir    string
	LocalDir        string
	DevImage        string
	Shell           string
	ContainerGoPath string
}

var DefaultConfig = Config{
	Port:            "10001:10001",
	DevImage:        "btburke/golang-dev",
	Shell:           "/bin/bash",
	ContainerGoPath: "/golang",
}

func ReadConfig() (*Config, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configFileLocs := []string{
		path.Join(os.Getwd(), ".devrc"),
		path.Join(usr.HomeDir, "/.config/.devrc"),
		path.Join(usr.HomeDir, ".devrc"),
	}
	configJsonFile, err := FindOneOf(configFileLocs...)
	if err != nil {
		c, configErr := NewConfig()
		if configErr != nil {
			clt.Warn("There was a problem processing the new configuration.  Using default configuration: %v", DefaultConfig)
			return &DefaultConfig, nil
		}
		return &c, nil
	}
	configJson, err := ioutil.ReadFile(configJsonFile)
	if err != nil {
		clt.Warn("Problem reading the config file located at %s. Creating new config...", configJsonFile)
		c, configErr := NewConfig()
		if configErr != nil {
			clt.Warn("There was a problem processing the new configuration.  Using default configuration: %v", DefaultConfig)
			return &DefaultConfig, nil
		}
		return &c, nil

	}
	var configFromFile Config
	if err := json.Unmarshal(configJson, &configFromFile); err != nil {
		clt.Warn("Problem reading the configuration located in %s. Creating new config...", configJsonFile)
	}
	return &configFromFile, nil
}

func NewConfig() (*Config, error) {

}

func FindOneOf(places ...string) (string, error) {
	for place := range places {
		_, err := os.Stat(place)
		switch err {
		case os.IsNotExist(err):
			continue
		case nil:
			return place, nil
		}
	}
	return "", fmt.Errorf("Not found")
}
