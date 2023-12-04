package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	var results int

	for _, s := range utils.LoadInput("4") {
		var result int

		sets := strings.Split(strings.Split(s, ":")[1], "|")
		winningNumbers := strings.Fields(sets[0])
		myNumbers := strings.Fields(sets[1])

		for _, winningNumber := range winningNumbers {
			for _, myNumber := range myNumbers {
				if winningNumber == myNumber {
					if result == 0 {
						result = 1
					} else {
						result *= 2
					}
				}
			}
		}

		results += result
	}

	fmt.Println(results)
}
