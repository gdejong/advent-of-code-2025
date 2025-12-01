package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gdejong/advent-of-code-2025/internal/input"
)

func main() {
	in := input.Content("cmd/day01/real_input.txt")

	fmt.Println(part1(in))
}

func part1(lines []string) int {
	dial := 50 // The dial starts by pointing at 50
	resetsToZero := 0

	for _, line := range lines {
		direction := line[0:1] // Either "L" or "R"
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		if direction == "L" {
			dial = (dial - rotation + 100) % 100
		} else if direction == "R" {
			dial = (dial + rotation) % 100
		} else {
			panic("invalid direction")
		}

		if dial == 0 {
			resetsToZero++
		}
	}

	return resetsToZero
}
