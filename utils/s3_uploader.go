package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Uploader handles file uploads to AWS S3
type S3Uploader struct {
	session    *session.Session
	bucketName string
	region     string
}

// AllowedMimeTypes defines which file types can be uploaded
var AllowedMimeTypes = map[string]bool{
	"application/pdf":    true, // PDF
	"image/jpeg":         true, // JPG/JPEG
	"image/png":          true, // PNG
	"image/gif":          true, // GIF
	"image/webp":         true, // WebP
	"application/msword": true, // DOC
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true, // DOCX
}

// IsValidFileType checks if the MIME type is allowed
func IsValidFileType(mimeType string) bool {
	// Extract base MIME type (ignore parameters like charset)
	baseMimeType := strings.Split(mimeType, ";")[0]
	baseMimeType = strings.TrimSpace(baseMimeType)
	return AllowedMimeTypes[baseMimeType]
}

// NewS3Uploader creates a new S3 uploader instance
func NewS3Uploader(region, bucketName string) (*S3Uploader, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	return &S3Uploader{
		session:    sess,
		bucketName: bucketName,
		region:     region,
	}, nil
}

// UploadFile uploads a file to S3 and returns the URL
func (u *S3Uploader) UploadFile(file multipart.File, header *multipart.FileHeader, folder string) (string, error) {
	// Validate file type
	contentType := header.Header.Get("Content-Type")
	if !IsValidFileType(contentType) {
		return "", fmt.Errorf("unsupported file type '%s'. allowed types: PDF, PNG, JPG, GIF, WebP, DOC, DOCX", contentType)
	}

	// Read file content using io.ReadAll for better efficiency
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// Validate file size (max 50MB)
	const maxFileSize = 50 * 1024 * 1024 // 50MB
	if int64(len(fileBytes)) > maxFileSize {
		return "", fmt.Errorf("file size exceeds maximum limit of 50MB")
	}

	// Generate unique filename with timestamp
	ext := filepath.Ext(header.Filename)
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%s/%d_%s%s", folder, timestamp, filepath.Base(header.Filename[:len(header.Filename)-len(ext)]), ext)

	// Create S3 service client
	svc := s3.New(u.session)

	// Upload file to S3 with explicit ACL and proper Content-Disposition
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(u.bucketName),
		Key:                  aws.String(filename),
		Body:                 bytes.NewReader(fileBytes),
		ContentType:          aws.String(contentType),
		ContentDisposition:   aws.String("attachment; filename=" + filepath.Base(header.Filename)),
		ServerSideEncryption: aws.String("AES256"),
		ACL:                  aws.String("private"), // Explicitly set to private
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload to S3: %w", err)
	}

	// Return the S3 URL
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.bucketName, u.region, filename)
	return url, nil
}

// UploadFileBytes uploads raw file bytes to S3 and returns the URL
func (u *S3Uploader) UploadFileBytes(fileBytes []byte, filename, contentType, folder string) (string, error) {
	// Validate file type
	if !IsValidFileType(contentType) {
		return "", fmt.Errorf("unsupported file type '%s'. allowed types: PDF, PNG, JPG, GIF, WebP, DOC, DOCX", contentType)
	}

	// Validate file size (max 50MB)
	const maxFileSize = 50 * 1024 * 1024 // 50MB
	if int64(len(fileBytes)) > maxFileSize {
		return "", fmt.Errorf("file size exceeds maximum limit of 50MB")
	}

	// Generate unique filename with timestamp
	ext := filepath.Ext(filename)
	timestamp := time.Now().Unix()
	key := fmt.Sprintf("%s/%d_%s%s", folder, timestamp, filepath.Base(filename[:len(filename)-len(ext)]), ext)

	// Create S3 service client
	svc := s3.New(u.session)

	// Upload file to S3 with explicit ACL and proper Content-Disposition
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(u.bucketName),
		Key:                  aws.String(key),
		Body:                 bytes.NewReader(fileBytes),
		ContentType:          aws.String(contentType),
		ContentDisposition:   aws.String("attachment; filename=" + filepath.Base(filename)),
		ServerSideEncryption: aws.String("AES256"),
		ACL:                  aws.String("private"), // Explicitly set to private
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload to S3: %w", err)
	}

	// Return the S3 URL
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.bucketName, u.region, key)
	return url, nil
}
