package utils

import "github.com/gofrs/uuid"

func GetUUID() (string, error) {
	ul, err := uuid.NewV4()
	return ul.String(), err
}
