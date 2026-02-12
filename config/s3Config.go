package config

import (
	"log"
	"os"

	"property-backend/utils"
)

var S3Uploader *utils.S3Uploader

// Valid AWS regions
var validRegions = map[string]bool{
	"us-east-1":      true,
	"us-east-2":      true,
	"us-west-1":      true,
	"us-west-2":      true,
	"eu-west-1":      true,
	"eu-central-1":   true,
	"ap-south-1":     true,
	"ap-southeast-1": true,
	"ap-northeast-1": true,
}

// InitS3 initializes the S3 uploader with configuration from environment variables
func InitS3() {
	// Check AWS credentials
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if accessKeyID == "" || secretAccessKey == "" {
		log.Println("⚠️  AWS_ACCESS_KEY_ID or AWS_SECRET_ACCESS_KEY not set, S3 upload will not work")
		log.Println("   ℹ️  Set credentials in .env file to enable S3 uploads")
		return
	}

	// Validate region
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-east-1" // default region
	}

	if !validRegions[region] {
		log.Printf("⚠️  Invalid AWS_REGION '%s', S3 upload will not work\n", region)
		log.Println("   ℹ️  Valid regions: us-east-1, us-east-2, us-west-1, us-west-2, eu-west-1, eu-central-1, ap-south-1, ap-southeast-1, ap-northeast-1")
		return
	}

	// Check bucket name
	bucketName := os.Getenv("AWS_S3_BUCKET")
	if bucketName == "" {
		log.Println("⚠️  AWS_S3_BUCKET not set, S3 upload will not work")
		log.Println("   ℹ️  Set AWS_S3_BUCKET in .env file")
		return
	}

	// Validate bucket name format
	if !isValidBucketName(bucketName) {
		log.Printf("⚠️  Invalid bucket name '%s', S3 upload will not work\n", bucketName)
		log.Println("   ℹ️  Bucket names must be 3-63 characters, lowercase alphanumeric and hyphens")
		return
	}

	uploader, err := utils.NewS3Uploader(region, bucketName)
	if err != nil {
		log.Printf("⚠️  Failed to initialize S3: %v\n", err)
		return
	}

	S3Uploader = uploader
	log.Println("✅ S3 uploader initialized successfully")
	log.Printf("   📍 Region: %s | 🪣 Bucket: %s\n", region, bucketName)
}

// isValidBucketName validates AWS S3 bucket name format
func isValidBucketName(name string) bool {
	if len(name) < 3 || len(name) > 63 {
		return false
	}

	// Check for valid characters (lowercase letters, numbers, hyphens, dots)
	for _, c := range name {
		if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-' || c == '.') {
			return false
		}
	}

	// Cannot start or end with hyphen
	if name[0] == '-' || name[len(name)-1] == '-' {
		return false
	}

	return true
}
