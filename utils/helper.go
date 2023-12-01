package utils

import (
	"fmt"
	"os"
	"strings"
)

func LoadInput(part string) []string {
	f, err := os.ReadFile(fmt.Sprintf("./input/%s.txt", part))
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), "\n")
}
