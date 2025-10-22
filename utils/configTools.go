package utils

import (
	"ZoranYuan_blog/config"
	"os"

	"github.com/goccy/go-yaml"
)

func BindWithYaml(yamlPath string, config *config.Config) error {
	_, err := os.Stat(yamlPath)

	if err != nil {
		return err
	}

	data, err := os.ReadFile(yamlPath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, config)
}
