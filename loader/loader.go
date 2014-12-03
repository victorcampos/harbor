package loader

import (
		"io/ioutil"
		yaml "gopkg.in/yaml.v2"
)

type HarborConfig struct {
	BucketRootPath string
	DownloadPath string `yaml:",omitempty"`
	FileList []struct{Path string; Optional bool}
}

func LoadConfig() (HarborConfig, error) {
	harborConfig := HarborConfig{}
	configFile, err := ioutil.ReadFile(".harbor.yml")

	if err != nil {
		return harborConfig, err
	}

	err = yaml.Unmarshal(configFile, &harborConfig)

	if err != nil {
		return harborConfig, err
	}

	return harborConfig, nil
}