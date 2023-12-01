package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var results int

	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine)`)

	mapping := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, s := range utils.LoadInput("1") {
		var currStr, result string

		items := strings.Split(s, "")

		for i := 0; i < len(s); i++ {
			if _, err := strconv.Atoi(items[i]); err == nil {
				result += items[i]
				break
			}

			currStr += items[i]

			if val, ok := mapping[re.FindString(currStr)]; ok {
				result += val
				break
			}
		}

		currStr = ""
		for i := len(items) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(items[i]); err == nil {
				result += items[i]
				break
			}

			currStr = items[i] + currStr

			if val, ok := mapping[re.FindString(currStr)]; ok {
				result += val
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
