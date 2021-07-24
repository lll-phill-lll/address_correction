package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Config struct {
	Port         string `yaml:"port"`
	FIASDataPath string `yaml:"fias_data_path"`
}

func Read() (*Config, error) {
	if len(os.Args) != 2 {
		return nil, errors.New("Wrong number of command line arguments, expected path to config file")
	}

	conf := &Config{}

	yamlFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
