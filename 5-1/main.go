package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type Converter struct {
	Map []ConverterMap
}

type ConverterMap struct {
	Source      int
	Destination int
	Range       int
}

func (c *Converter) Convert(value int) int {
	for _, m := range c.Map {
		if value > m.Source && value < m.Source+m.Range {
			return m.Destination + (value - m.Source)
		}
	}
	return value
}

func main() {
	//var results int

	re := regexp.MustCompile(`(?m)([a-z- ]+):|([\d ]+)`)

	var seeds []string
	converters := make(map[string]*Converter)

	currName := ""
	for _, item := range re.FindAllStringSubmatch(utils.LoadRawInput("5"), -1) {
		if item[1] == "seeds" {
			currName = "seeds"
			continue
		} else if strings.HasSuffix(item[1], "map") {
			currName = strings.TrimSuffix(item[1], " map")
			converters[currName] = &Converter{}
			continue
		}

		if currName == "seeds" {
			seeds = strings.Fields(item[2])
		} else {
			line := strings.Fields(item[2])

			converters[currName].Map = append(converters[currName].Map, ConverterMap{
				Source:      utils.MustAtoi(line[1]),
				Destination: utils.MustAtoi(line[0]),
				Range:       utils.MustAtoi(line[2]),
			})
		}
	}

	//Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
	//Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
	//Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, humidity 82, location 86.
	//Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, humidity 35, location 35.

	var results []int
	for _, seed := range seeds {
		v := converters["seed-to-soil"].Convert(utils.MustAtoi(seed))
		v = converters["soil-to-fertilizer"].Convert(v)
		v = converters["fertilizer-to-water"].Convert(v)
		v = converters["water-to-light"].Convert(v)
		v = converters["light-to-temperature"].Convert(v)
		v = converters["temperature-to-humidity"].Convert(v)
		v = converters["humidity-to-location"].Convert(v)

		results = append(results, v)
	}

	fmt.Println(slices.Min(results))
}
