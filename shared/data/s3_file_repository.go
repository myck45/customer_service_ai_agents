package data

type S3FileRepository interface {
	// UploadFile uploads a file to an S3 bucket and returns the URL of the file.
	UploadFile(fileName string, fileBytes []byte, contentType string) (string, error)

	// DeleteFile deletes a file from an S3 bucket.
	DeleteFile(fileName string) error

	// GetFileURL generates a pre-signed URL for for a file
	GetFileURL(fileName string, expirationMins int) (string, error)

	// ListFiles lists all files in a specific directory/prefix
	ListFiles(prefix string) ([]string, error)
}
