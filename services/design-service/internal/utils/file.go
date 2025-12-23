package utils

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// Allowed file types
var allowedMIMEs = map[string]bool{
	"image/jpeg":    true,
	"image/png":     true,
	"application/pdf": true,
}

// SaveFile saves a multipart file to disk and returns the relative path
func SaveFile(file *multipart.FileHeader) (string, error) {
	// Validate MIME type
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return "", err
	}

	mime := http.DetectContentType(buffer)
	if !allowedMIMEs[mime] {
		return "", fmt.Errorf("unsupported file type: %s", mime)
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	uniqueName := uuid.NewString() + ext
	relativePath := filepath.Join("uploads", uniqueName)
	fullPath := filepath.Join("/app", relativePath) // matches WORKDIR

	// Save file
	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = dst.ReadFrom(src); err != nil {
		return "", err
	}

	return "/" + relativePath, nil // e.g., "/uploads/abc123.pdf"
}