package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
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

func (c *Converter) ReverseConvert(value int) int {
	for _, m := range c.Map {
		if value >= m.Destination && value < m.Destination+m.Range {
			return m.Source + (value - m.Destination)
		}
	}
	return value
}

func main() {
	//var results int

	re := regexp.MustCompile(`(?m)([a-z- ]+):|([\d ]+)`)

	var seeds []int
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
			strSeeds := strings.Fields(item[2])
			for _, strSeed := range strSeeds {
				seeds = append(seeds, utils.MustAtoi(strSeed))
			}
		} else {
			line := strings.Fields(item[2])

			converters[currName].Map = append(converters[currName].Map, ConverterMap{
				Source:      utils.MustAtoi(line[1]),
				Destination: utils.MustAtoi(line[0]),
				Range:       utils.MustAtoi(line[2]),
			})
		}
	}

	for i := 0; i < 10000000; i++ {
		v := converters["humidity-to-location"].ReverseConvert(i)
		v = converters["temperature-to-humidity"].ReverseConvert(v)
		v = converters["light-to-temperature"].ReverseConvert(v)
		v = converters["water-to-light"].ReverseConvert(v)
		v = converters["fertilizer-to-water"].ReverseConvert(v)
		v = converters["soil-to-fertilizer"].ReverseConvert(v)
		v = converters["seed-to-soil"].ReverseConvert(v)

		for z := 0; z < len(seeds); z += 2 {
			if v >= seeds[z] && v < seeds[z]+seeds[z+1] {
				fmt.Println(i)
				return
			}
		}
	}
}
