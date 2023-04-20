package model

type Attachment struct {
	Model
	StorageID uint64 `gorm:"not null"` // ID of the storage
	FilePath  string `gorm:"not null"` // path of the file

	UploaderIP string `gorm:"not null"` // IP address of the uploader

	FileName        string `gorm:"not null"` // original file name
	FileSize        int64  `gorm:"not null"` // in bytes
	FileContentType string `gorm:"not null"` // MIME type
	FileSha256Sum   string `gorm:"not null"` // SHA256 sum of the file
}
