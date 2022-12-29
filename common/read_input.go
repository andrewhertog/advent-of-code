package common

import (
	"os"
)

func ReadInput(fn string) (string, error) {
	dat, err := os.ReadFile(fn)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
