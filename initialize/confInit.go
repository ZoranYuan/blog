package initialize

import (
	"ZoranYuan_blog/config"
	"ZoranYuan_blog/utils"
	"os"
	"path/filepath"
)

func InitConfig() (*config.Config, error) {
	var config config.Config

	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	yamlPath := filepath.Join(rootPath, "config.yaml")
	err = utils.BindWithYaml(yamlPath, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
