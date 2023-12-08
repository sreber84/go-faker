package main

import (
	"crypto/rand"
	"fmt"
)

const (
	randomStringLength = 10024
)

func generateRandomString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", randomBytes), nil
}

func main() {
	for {
		randomString, err := generateRandomString(randomStringLength)
		if err != nil {
			fmt.Println("Error generating random string:", err)
			return
		}

		fmt.Printf("Generated random string: %s\n", randomString)
	}
}
