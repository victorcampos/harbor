package config

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type HarborFile struct {
	S3Path   string
	FileName string
}

type HarborConfig struct {
	ImageTag string
	S3       struct {
		Bucket   string
		BasePath string
	}
	DownloadPath string `yaml:",omitempty"`
	Files        []HarborFile
	Commands     []string
}

func Load() (HarborConfig, error) {
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
