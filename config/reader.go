package config


import (
    "gopkg.in/yaml.v3"
    "io/ioutil"
)

type Config struct {
    Port string `yaml:"port"`
    FIASDataPath string `yaml:"fias_data_path"`

}

func Read() (*Config, error) {

    conf := &Config{}

    yamlFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        return nil, err
    }

    err = yaml.Unmarshal(yamlFile, conf)
    if err != nil {
        return nil, err
    }

    return conf, nil
}
