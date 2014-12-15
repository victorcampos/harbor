package download

import (
	"fmt"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"github.com/victorcampos/harbor/config"
	"io/ioutil"
	"path/filepath"
)

func FromS3(harborConfig config.HarborConfig) error {
	fileListLength := len(harborConfig.Files)

	if fileListLength > 0 {
		bucket, err := getBucket(harborConfig.S3.Bucket)

		if err != nil {
			return err
		}

		fmt.Printf("--- Downloading from bucket: %s\r\n", harborConfig.S3.Bucket)
		fmt.Printf("--- Files to be downloaded: %d\r\n", fileListLength)
		fmt.Printf("--- Downloading to: %s\r\n", harborConfig.DownloadPath)
		for key, value := range harborConfig.Files {
			fmt.Printf("--- Downloading file %d of %d...\r\n", key+1, fileListLength)

			err := downloadFile(bucket, harborConfig.S3.BasePath, harborConfig.DownloadPath, value)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getBucket(bucketName string) (*s3.Bucket, error) {
	var bucket *s3.Bucket

	awsAuthentication, err := aws.EnvAuth()
	if err != nil {
		return bucket, err
	}

	s3Connection := s3.New(awsAuthentication, aws.USEast)
	bucket = s3Connection.Bucket(bucketName)
	if err != nil {
		return bucket, err
	}

	return bucket, nil
}

func downloadFile(bucket *s3.Bucket, s3BasePath string, downloadPath string, file config.HarborFile) error {
	outputFilePath := filepath.Join(downloadPath, file.FileName)
	s3FilePath := filepath.Join(s3BasePath, file.S3Path)

	fmt.Printf("File: %s\r\n", outputFilePath)

	contents, err := bucket.Get(s3FilePath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFilePath, contents, 0644)
	if err != nil {
		return err
	}

	return nil
}
