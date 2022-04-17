package nanoid

import (
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateNanoID() (string, error) {
	id, err := gonanoid.New()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return id, nil
}
