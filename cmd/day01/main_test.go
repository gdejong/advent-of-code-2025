package main

import (
	"testing"

	"github.com/gdejong/advent-of-code-2025/internal/input"
)

func TestPart1(t *testing.T) {
	testData := input.Content("test_input.txt")

	testAnswer := part1(testData)

	if testAnswer != 3 {
		t.Errorf("wrong test answer, got: %d", testAnswer)
	}

	// and let's also assert the real input

	realData := input.Content("real_input.txt")

	realAnswer := part1(realData)

	if realAnswer != 1018 {
		t.Errorf("wrong real answer, got: %d", realAnswer)
	}
}

func TestPart2(t *testing.T) {
	testData := input.Content("test_input.txt")

	testAnswer := part2(testData)

	if testAnswer != 6 {
		t.Errorf("wrong test answer, got: %d", testAnswer)
	}

	// and let's also assert the real input

	realData := input.Content("real_input.txt")

	realAnswer := part2(realData)

	if realAnswer != 5815 {
		t.Errorf("wrong real answer, got: %d", realAnswer)
	}
}
