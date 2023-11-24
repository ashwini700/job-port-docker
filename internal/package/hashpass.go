package pkgs

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

)

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {

	// Generate a bcrypt hash from the provided password.
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error is hashing the password : %w", err)
	}
	return string(hashedPass), nil

}

// CheckHashedPassword compares a plain password with a bcrypt hashed password.
func CheckHashedPassword(password string, hashedPassword string) error {
	// Compare the provided plain password with the bcrypt hashed password.
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err		// Passwords do not match.
	}
	return nil
}
