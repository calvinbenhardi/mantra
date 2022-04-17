package nanoid

import (
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateNanoID() (string, error) {
	nanoID, err := gonanoid.New()
	if err != nil {
		log.Fatal(err)

		return "", err
	}

	return nanoID, nil
}

func GenerateNonSecureNanoID(strings string, length int) (string, error) {
	defaultStrings := "abcdefghijklmnopqrstuvwxyz1234567890"
	if strings == "" {
		strings = defaultStrings
	}

	nanoID, err := gonanoid.Generate(strings, length)
	if err != nil {
		log.Fatal(err)

		return "", err
	}

	return nanoID, nil
}
