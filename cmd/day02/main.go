package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdejong/advent-of-code-2025/internal/input"
	"github.com/gdejong/advent-of-code-2025/internal/must"
)

func main() {
	in := input.Content("cmd/day02/real_input.txt")

	fmt.Println(part1(in))
}

func part1(lines []string) int {
	if len(lines) != 1 {
		panic("wrong number of lines")
	}

	// split each range by a comma, for example 11-22,95-115 gives 2 ranges
	ranges := strings.Split(lines[0], ",")

	total := 0

	for _, v := range ranges {
		minMax := strings.Split(v, "-")
		minRange := must.NoError(strconv.Atoi(minMax[0]))
		maxRange := must.NoError(strconv.Atoi(minMax[1]))
		for id := minRange; id <= maxRange; id++ {
			if isInvalid(id) {
				total += id
			}
		}
	}

	return total
}

func isInvalid(id int) bool {
	asString := strconv.Itoa(id)

	// find the middle position, up to that character position
	// we'll construct smaller strings to check against.
	middlePosition := len(asString) / 2

	for i := 0; i < middlePosition; i++ {
		testString := asString[0 : i+1]
		// you can find the invalid IDs by looking for any ID
		// which is made only of some sequence of digits repeated twice
		if testString+testString == asString {
			return true
		}
	}

	return false
}
