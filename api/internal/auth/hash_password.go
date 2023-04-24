package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes given string and used for hashing users' passwords.
func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 12)

	return string(bytes), err
}

// CheckPasswordHash compares given two hashes.
func ComparePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
