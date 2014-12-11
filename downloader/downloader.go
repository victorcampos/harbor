package downloader

import (
	"fmt"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"github.com/victorcampos/harbor/config"
	"io/ioutil"
	"path/filepath"
)

func DownloadFromS3(harborConfig config.HarborConfig) error {
	fileListLength := len(harborConfig.Files)

	if fileListLength > 0 {
		awsAuthentication, err := aws.EnvAuth()
		if err != nil {
			return err
		}

		s3Connection := s3.New(awsAuthentication, aws.USEast)
		bucket := s3Connection.Bucket(harborConfig.S3.Bucket)
		if err != nil {
			return err
		}

		fmt.Printf("Downloading from bucket: %s\r\n", harborConfig.S3.Bucket)
		fmt.Printf("Files to be downloaded: %d\r\n", fileListLength)
		fmt.Printf("Downloading to: %s\r\n\r\n", harborConfig.DownloadPath)
		for key, value := range harborConfig.Files {
			outputFilePath := filepath.Join(harborConfig.DownloadPath, value.FileName)

			fmt.Printf("Downloading file number %d of %d...\r\n", key+1, fileListLength)
			fmt.Printf("File: %s\r\n", outputFilePath)

			contents, err := bucket.Get(harborConfig.S3.BasePath + value.S3Path)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(outputFilePath, contents, 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
