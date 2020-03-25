package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Hash hashes user password
func Hash(password string) string {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(pass)
}

// UnHashAndCompare unhashes password and compares to input
func UnHashAndCompare(password string, check string) bool {
	errf := bcrypt.CompareHashAndPassword([]byte(check), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}
	return true
}
