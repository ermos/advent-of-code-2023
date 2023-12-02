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
		sets := strings.Split(strings.Split(game, ":")[1], ";")

		var maxRed, maxBlue, maxGreen int

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

			maxRed = max(red, maxRed)
			maxGreen = max(green, maxGreen)
			maxBlue = max(blue, maxBlue)
		}

		results += maxRed * maxGreen * maxBlue
	}

	fmt.Println(results)
}
