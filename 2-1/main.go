package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var results int

	re := regexp.MustCompile(`(\d+) (blue|red|green)`)

	for _, game := range utils.LoadInput("2") {
		id := utils.MustAtoi(strings.Fields(strings.Split(game, ":")[0])[1])
		sets := strings.Split(strings.Split(game, ":")[1], ";")

		good := true

		for _, s := range sets {
			var red, green, blue int

			for _, match := range re.FindAllStringSubmatch(s, -1) {
				switch match[2] {
				case "red":
					red += utils.MustAtoi(match[1])
				case "green":
					green += utils.MustAtoi(match[1])
				case "blue":
					blue += utils.MustAtoi(match[1])
				}
			}

			if red > 12 || green > 13 || blue > 14 {
				good = false
			}
		}

		if good {
			results += id
		}
	}

	fmt.Println(results)
}
