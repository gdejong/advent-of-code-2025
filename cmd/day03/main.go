package main

import (
	"fmt"

	"github.com/gdejong/advent-of-code-2025/internal/input"
	"github.com/gdejong/advent-of-code-2025/internal/must"
)

func main() {
	in := input.Content("cmd/day03/real_input.txt")

	fmt.Println(part1(in))

	fmt.Println(part2(in))
}

func part1(lines []string) int {
	total := 0

	for _, v := range lines {
		// find the biggest tens value
		// note the edge case if it is at the end of the battery bank - in that case we pick the second-largest tens value
		tens := 0
		tensIndex := 0
		secondHighestTens := 0
		secondHighestTensIndex := 0

		for i := 0; i < len(v); i++ {
			asInt := must.Int(string(v[i]))
			if asInt > tens {
				secondHighestTens = tens
				secondHighestTensIndex = tensIndex

				tens = asInt
				tensIndex = i
			}
		}

		// handle the edge-case, when the highest tens value was at the end of the battery bank we use the second-highest tens value
		if tensIndex == len(v)-1 {
			tens = secondHighestTens
			tensIndex = secondHighestTensIndex
		}

		// now find the biggest ones, we only have to search to the right of the tens value
		ones := 0
		for i := tensIndex + 1; i < len(v); i++ {
			asInt := must.Int(string(v[i]))
			if asInt > ones {
				ones = asInt
			}
		}

		total += tens*10 + ones
	}

	return total
}

func part2(lines []string) int {
	return 0
}
