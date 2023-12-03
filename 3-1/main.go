package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

type Number struct {
	Value        string
	NestedPoints []Point
}

type Point struct {
	Value string
	X     int
	Y     int
}

func (n *Number) AddPoint(matrix [][]string, x, y int) {
	if y >= 0 && y <= len(matrix)-1 && x >= 0 && x <= len(matrix[y])-1 {
		n.NestedPoints = append(n.NestedPoints, Point{Value: matrix[y][x], X: x, Y: y})
	}
}

func (n *Number) HasNestedSymbol() bool {
	for _, point := range n.NestedPoints {
		if !utils.IsInt(point.Value) && point.Value != "." {
			return true
		}
	}

	return false
}

func (n *Number) ToInt() int {
	return utils.MustAtoi(n.Value)
}

func main() {
	var results int
	var numbers []Number

	var matrix [][]string
	for _, line := range utils.LoadInput("3") {
		matrix = append(matrix, strings.Split(line, ""))
	}

	for y, line := range matrix {
		var number Number

		for x, v := range line {
			if !utils.IsInt(v) {
				if number.Value != "" {
					numbers = append(numbers, number)
					number = Number{}
				}
				continue
			}

			if number.Value == "" {
				// premier chiffre
				number.AddPoint(matrix, x-1, y)   // x 1
				number.AddPoint(matrix, x-1, y-1) // x^ 1
				number.AddPoint(matrix, x-1, y+1) // xv 1
			}

			if x+1 <= len(line)-1 && !utils.IsInt(line[x+1]) {
				// dernier chiffre
				number.AddPoint(matrix, x+1, y)   // 1 x
				number.AddPoint(matrix, x+1, y-1) // 1 ^x
				number.AddPoint(matrix, x+1, y+1) // 1 vx
			}

			number.AddPoint(matrix, x, y-1) // ^1
			number.AddPoint(matrix, x, y+1) // v1

			number.Value += v
		}

		if number.Value != "" {
			numbers = append(numbers, number)
			number = Number{}
		}
	}

	for _, number := range numbers {
		if number.HasNestedSymbol() {
			results += number.ToInt()
		}
	}

	fmt.Println(results)
}
