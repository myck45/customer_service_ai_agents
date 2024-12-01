package data

import (
	"bytes"
	"fmt"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/sirupsen/logrus"
)

type S3FileRepositoryImpl struct {
	s3Client   s3iface.S3API
	bucketName string
	basePath   string
}

// DeleteFile implements S3FileRepository.
func (s *S3FileRepositoryImpl) DeleteFile(fileName string) error {
	fullKey := filepath.Join(s.basePath, fileName)

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(fullKey),
	}

	_, err := s.s3Client.DeleteObject(input)
	if err != nil {
		logrus.WithError(err).Error("[S3FileRepositoryImpl] failed to delete file")
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}

// GetFileURL implements S3FileRepository.
func (s *S3FileRepositoryImpl) GetFileURL(fileName string, expirationMins int) (string, error) {

	fullKey := filepath.Join(s.basePath, fileName)

	input := &s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(fullKey),
	}

	// Generate a pre-signed URL for the file
	req, _ := s.s3Client.GetObjectRequest(input)
	url, err := req.Presign(time.Duration(expirationMins) * time.Minute)
	if err != nil {
		logrus.WithError(err).Error("[S3FileRepositoryImpl] failed to generate pre-signed URL")
		return "", fmt.Errorf("failed to generate pre-signed URL: %v", err)
	}

	return url, nil
}

// ListFiles implements S3FileRepository.
func (s *S3FileRepositoryImpl) ListFiles(prefix string) ([]string, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucketName),
		Prefix: aws.String(filepath.Join(s.basePath, prefix)),
	}

	result, err := s.s3Client.ListObjectsV2(input)
	if err != nil {
		logrus.WithError(err).Error("[S3FileRepositoryImpl] failed to list files")
		return nil, fmt.Errorf("failed to list files: %v", err)
	}

	var files []string
	for _, obj := range result.Contents {
		files = append(files, *obj.Key)
	}

	return files, nil
}

// UploadFile implements S3FileRepository.
func (s *S3FileRepositoryImpl) UploadFile(fileName string, fileBytes []byte, contentType string) (string, error) {
	// Generate a unique file path
	uniqueFileName := fmt.Sprintf("%s_%d%s",
		filepath.Base(fileName),
		time.Now().UnixNano(),
		filepath.Ext(fileName),
	)

	// Construct the full S3 key
	fullKey := filepath.Join(s.basePath, uniqueFileName)

	// Prepare the input for S3 upload
	input := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(fullKey),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
	}

	// Perform the upload
	_, err := s.s3Client.PutObject(input)
	if err != nil {
		logrus.WithError(err).Errorf("*** [UploadFile] Error uploading file %s", fileName)
		return "", fmt.Errorf("error uploading file: %v", err)
	}

	// Return the full S3 path
	return fullKey, nil
}

func NewS3FileRepositoryImpl(s3Client *s3.S3, bucketName, basePath string) S3FileRepository {
	return &S3FileRepositoryImpl{
		s3Client:   s3Client,
		bucketName: bucketName,
		basePath:   basePath,
	}
}
