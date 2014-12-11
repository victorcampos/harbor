package main

import (
	"fmt"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"github.com/victorcampos/harbor/loader"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	harborConfig, err := loader.LoadConfig()
	checkError(err)

	fileListLength := len(harborConfig.Files)

	if fileListLength > 0 {
		awsAuthentication, err := aws.EnvAuth()
		checkError(err)

		s3Connection := s3.New(awsAuthentication, aws.USEast)
		bucket := s3Connection.Bucket(harborConfig.S3.Bucket)
		checkError(err)

		fmt.Printf("Downloading from bucket: %s\r\n", harborConfig.S3.Bucket)
		fmt.Printf("Files to be downloaded: %d\r\n", fileListLength)
		fmt.Printf("Downloading to: %s\r\n\r\n", harborConfig.DownloadPath)
		for key, value := range harborConfig.Files {
			outputFilePath := filepath.Join(harborConfig.DownloadPath, value.FileName)

			fmt.Printf("Downloading file number %d of %d...\r\n", key+1, fileListLength)
			fmt.Printf("File: %s\r\n", outputFilePath)

			contents, err := bucket.Get(harborConfig.S3.BasePath + value.S3Path)
			checkError(err)

			err = ioutil.WriteFile(outputFilePath, contents, 0644)
			checkError(err)
		}
	}

	commandListLength := len(harborConfig.Commands)

	if commandListLength > 0 {
		fmt.Printf("Commands to be executed: %d\r\n\r\n", commandListLength)
		for key, value := range harborConfig.Commands {
			fmt.Printf("Executing command number %d of %d...\r\n", key+1, commandListLength)

			commandParts := strings.Fields(value)
			commandHead := commandParts[0]
			commandArgs := commandParts[1:len(commandParts)]

			command := exec.Command(commandHead, commandArgs...)
			commandOutput, _ := command.CombinedOutput()

			fmt.Printf("Executing: %s\r\n", commandHead)
			fmt.Printf("Output of: %s\r\n%s", commandHead, string(commandOutput))
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
