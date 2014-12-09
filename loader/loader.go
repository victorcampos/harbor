package loader

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type HarborConfig struct {
	ImageTag       string
	BucketRootPath string
	DownloadPath   string `yaml:",omitempty"`
	Files          []struct {
		Path string
	}
	Commands []string
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
