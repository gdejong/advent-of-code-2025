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

	fmt.Println(part2(in))
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

func part2(lines []string) int {
	dial := 50 // The dial starts by pointing at 50
	pointedAtZero := 0

	restAtZero := false

	for _, line := range lines {
		direction := line[0:1] // Either "L" or "R"
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		// All full circle rotates count as one. Count those here and reduce the rest of the problem to a rotation between 1 and 99
		if rotation > 100 {
			pointedAtZero += rotation / 100 // for example R640 counts 6 resets...
			rotation %= 100                 // and leaves 40 for the rest of the rotation.
		}

		if direction == "L" {
			dial -= rotation

			if dial <= 0 {
				if !restAtZero {
					pointedAtZero++
				}
				// If we landed on zero exactly don't +100
				if dial < 0 {
					dial += 100
				}
			}

		} else if direction == "R" {
			dial += rotation

			if dial >= 100 {
				if !restAtZero {
					pointedAtZero++
				}
				dial -= 100
			}

		} else {
			panic("invalid direction")
		}

		// mark if we ended exactly on zero, this helps us later to avoid counting resets twice
		restAtZero = dial == 0
	}

	return pointedAtZero
}
