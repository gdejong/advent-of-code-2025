package main

import (
	"testing"

	"github.com/gdejong/advent-of-code-2025/internal/input"
)

func TestPart1(t *testing.T) {
	testData := input.Content("test_input.txt")

	testAnswer := part1(testData)

	if testAnswer != 1227775554 {
		t.Errorf("wrong test answer, got: %d", testAnswer)
	}
}
