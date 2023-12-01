package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var results int

	for _, s := range utils.LoadInput("1") {
		var result string

		items := strings.Split(s, "")

		for i := 0; i < len(s); i++ {
			if _, err := strconv.Atoi(items[i]); err == nil {
				result += items[i]
				break
			}
		}

		for i := len(items) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(items[i]); err == nil {
				result += items[i]
				break
			}
		}

		intResult, err := strconv.Atoi(result)
		if err != nil {
			panic(err)
		}

		results += intResult
	}

	fmt.Println(results)
}
