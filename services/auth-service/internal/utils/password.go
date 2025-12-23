// Package utils provides helper utilities for authentication.
package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plaintext password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// ComparePassword compares a plaintext password with a bcrypt hash.
func ComparePassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
