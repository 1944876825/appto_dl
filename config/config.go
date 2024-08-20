package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var Conf = Model{}

type Model struct {
	Port     int      `yaml:"port"`
	Token    string   `yaml:"token"`
	Endata   string   `yaml:"data"`
	EnMac    string   `yaml:"mac"`
	Features []string `yaml:"features"`
}

func Load() {
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		return
	}
	if err := yaml.Unmarshal(content, &Conf); err != nil {
		panic(err)
	}
}

func Save() {
	y, err := yaml.Marshal(Conf)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("config.yaml", y, 0644); err != nil {
		panic(err)
	}
}
