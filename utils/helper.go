package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadInput(part string) []string {
	f, err := os.ReadFile(fmt.Sprintf("./input/%s.txt", part))
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), "\n")
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}
