package config

import (
	"github.com/victorcampos/harbor/commandline"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type HarborFile struct {
	S3Path     string
	FileName   string
	Permission int
}

type HarborConfig struct {
	ImageTag      string
	CliConfigVars commandline.ConfigVarsMap
	S3            struct {
		Bucket   string
		BasePath string
	}
	DownloadPath string `yaml:",omitempty"`
	Files        []HarborFile
	Commands     []string
}

func Load(cliConfigVars commandline.ConfigVarsMap) (HarborConfig, error) {
	harborConfig := HarborConfig{}

	configFile, err := ioutil.ReadFile(".harbor.yml")
	if err != nil {
		return harborConfig, err
	}

	configFile = SetEnv(cliConfigVars, configFile)

	err = yaml.Unmarshal(configFile, &harborConfig)

	if err != nil {
		return harborConfig, err
	}

	return harborConfig, nil
}
