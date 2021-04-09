package watsonconn

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config have all configuration settings for watson assitant.
type Config struct {
	Region      string `yaml:"region"`
	Instance    string `yaml:"instance"`
	AssistantID string `yaml:"assistant"`
	Version     string `yaml:"version"`
	Credentials string `yaml:"credentials"`
}

// ParseFromYAML method reads the YAML file and set the values to Config.
func (cfg *Config) ParseFromYAML(fileName string) error {

	if fileName == "" {
		return errors.New("Please provide yaml file path")
	}

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return err
	}

	return nil
}
