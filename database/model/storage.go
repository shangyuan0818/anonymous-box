package model

type Storage struct {
	Model
	Type StorageType `gorm:"not null"`

	// Common
	Name     string `gorm:"not null"`                     // name of the storage
	UsedSize int64  `gorm:"not null;default:0"`           // in bytes (0 by default)
	MaxSize  int64  `gorm:"not null;default:10737418240"` // in bytes (10GB by default)
	IsInUse  bool   `gorm:"not null;default:true"`        // is in use (true by default)
	IsPublic bool   `gorm:"not null;default:false"`       // is public accessible, not used for local storage

	// S3
	S3Endpoint        string `gorm:"not null"` // Endpoint URL of S3
	S3AccessKeyID     string `gorm:"not null"` // Access Key ID of S3
	S3SecretAccessKey string `gorm:"not null"` // Secret Access Key of S3
	S3Bucket          string `gorm:"not null"` // Bucket name of S3
	S3Region          string `gorm:"not null"` // Region name of S3

	// GCS
	GCSProjectID string `gorm:"not null"` // Project ID of GCP
	GCSBucket    string `gorm:"not null"` // Bucket name of GCS
	GCSRegion    string `gorm:"not null"` // Region name of GCS
	GCSJSONKey   string `gorm:"not null"` // JSON key of GCS service account

	// Local
	LocalPath      string `gorm:"not null"` // Local path to store files
	LocalPublicURL string `gorm:"not null"` // Public URL of the local storage
}

type StorageType string

const (
	StorageTypeLocal StorageType = "local"
	StorageTypeS3    StorageType = "s3"
	StorageTypeGCS   StorageType = "gcs"
)
