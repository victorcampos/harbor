package main

import (
	"fmt"
	loader "github.com/victorcampos/harbor/loader"
)

type HarborConfig struct {
	BucketRootPath string
	DownloadPath string `yaml:",omitempty"`
	FileList []struct{Path string; Optional bool}
}

func main() {
	harborConfig, err := loader.LoadConfig()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(harborConfig.FileList[0].Path)
}