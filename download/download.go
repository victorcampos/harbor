package download

import (
	"fmt"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"github.com/victorcampos/harbor/config"
	"io/ioutil"
	"os"
	"path/filepath"
	"errors"
)

func FromS3(harborConfig config.HarborConfig) error {
	var region aws.Region
	var exists bool
	fileListLength := len(harborConfig.Files)

	if fileListLength > 0 {
		if harborConfig.S3.Region == "" {
			region = aws.USEast
		} else {
			region, exists = aws.Regions[harborConfig.S3.Region]
			if !exists {
				return errors.New("This region is not valid: " + harborConfig.S3.Region)
			}
		}

		bucket, err := getBucket(harborConfig.S3.Bucket, region)

		if err != nil {
			return err
		}

		fmt.Printf("--- Downloading from bucket: %s\r\n", harborConfig.S3.Bucket)
		fmt.Printf("--- Bucket region is set to %s\r\n", region.Name)
		fmt.Printf("--- Files to be downloaded: %d\r\n", fileListLength)
		fmt.Printf("--- Downloading to: %s\r\n", harborConfig.DownloadPath)
		for key, value := range harborConfig.Files {
			fmt.Printf("--- Downloading file %d of %d...\r\n", key+1, fileListLength)

			// TODO: Enable custom download path per file
			err := downloadFile(bucket, harborConfig.S3.BasePath, harborConfig.DownloadPath, value)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getBucket(bucketName string, region aws.Region) (*s3.Bucket, error) {
	var bucket *s3.Bucket

	awsAuthentication, err := aws.EnvAuth()
	if err != nil {
		return bucket, err
	}

	s3Connection := s3.New(awsAuthentication, region)
	bucket = s3Connection.Bucket(bucketName)
	if err != nil {
		return bucket, err
	}

	return bucket, nil
}

func downloadFile(bucket *s3.Bucket, s3BasePath string, downloadPath string, file config.HarborFile) error {
	s3FilePath := filepath.Join(s3BasePath, file.S3Path)
	outputFilePath := filepath.Join(downloadPath, file.FileName)
	outputDirectory := filepath.Dir(outputFilePath)

	os.MkdirAll(outputDirectory, 0755)

	fmt.Printf("S3 Path: %s\r\n", s3FilePath)
	fmt.Printf("File: %s\r\n", outputFilePath)

	// FIXME: Use GetReader to stream file contents instead of loading all the file to memory before writing
	contents, err := bucket.Get(s3FilePath)
	if err != nil {
		return err
	}

	// Sets default permission if not configured in YAML
	if file.Permission == 0 {
		file.Permission = 0644
	}

	filemode := os.FileMode(file.Permission & 0777)
	err = ioutil.WriteFile(outputFilePath, contents, filemode)
	if err != nil {
		return err
	}

	err = os.Chmod(outputFilePath, filemode)
	if err != nil {
		return err
	}

	return nil
}
