package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// function for bcryptEncoder
func BcryptEncoder(value string) string {
	// enconder value to bcrypt
	encoded, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)

	if err != nil {
		return "error"
	}

	return string(encoded)
}

func BcryptCompare(decoded, encoded string) bool {
	compare := bcrypt.CompareHashAndPassword([]byte(encoded), []byte(decoded))

	fmt.Print(compare)

	return compare == nil

}
