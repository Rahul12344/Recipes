package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Hash Hashes user password with bcrypt. Simple wrapper for bcrypt hasher.
func Hash(password string) string {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(pass)
}

// UnHashAndCompare Unhashes password and compares to input. Simple wrapper for bycrypt unhasher and comparison.
func UnHashAndCompare(password string, check string) bool {
	errf := bcrypt.CompareHashAndPassword([]byte(check), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}
	return true
}
