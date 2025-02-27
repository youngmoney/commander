package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"regexp"
)

type Command struct {
	Name           string         `yaml:"name"`
	MatchPathRegex *regexp.Regexp `yaml:"match_path_regex"`
	MatchCommand   string         `yaml:"match_command"`
	Command        string         `yaml:"command"`
}

type Commander struct {
	Commands []Command `yaml:"commands"`
}

type Config struct {
	Commander Commander `yaml:"commander"`
}

func ReadConfig(filename string) Config {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("unable to read config: ", filename)
		os.Exit(1)
	}

	config := Config{}
	if err := yaml.Unmarshal(raw, &config); err != nil {
		fmt.Println("unable to parse config: ", filename)
		os.Exit(1)
	}

	return config
}
