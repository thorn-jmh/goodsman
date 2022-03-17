package utils

import "github.com/gofrs/uuid"

func GetUID() (string, error) {
	ul, err := uuid.NewV4()
	return ul.String(), err
}
