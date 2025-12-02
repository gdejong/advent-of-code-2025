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

	fmt.Println(part2(in))
}

func run(lines []string, validFn func(string) bool) int {
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
			if validFn(strconv.Itoa(id)) {
				total += id
			}
		}
	}

	return total
}

func part1(lines []string) int {
	return run(lines, func(id string) bool {
		// find the middle position, we'll cut the id in half using that
		middlePosition := len(id) / 2

		// you can find the invalid IDs by looking for any ID
		// which is made only of some sequence of digits repeated twice
		if id[:middlePosition] == id[middlePosition:] {
			return true
		}

		return false
	})
}

func part2(lines []string) int {
	return run(lines, func(id string) bool {
		// find the middle position, up to that character position
		// we'll construct smaller strings to check against.
		middlePosition := len(id) / 2

		// 446446 gets broken down into:
		// - 4
		// - 44
		// - 446
		for i := 0; i < middlePosition; i++ {
			testString := id[0 : i+1]
			// an ID is invalid if it is made only of some sequence
			// of digits repeated at least twice
			j := 1
			for {
				testSequence := strings.Repeat(testString, j)
				if testSequence == id {
					return true
				}
				// stop if the constructed id no longer is a prefix of the original
				if !strings.HasPrefix(id, testSequence) {
					break
				}
				j++
			}
		}

		return false
	})
}
