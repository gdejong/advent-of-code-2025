package main

import (
	"testing"

	"github.com/gdejong/advent-of-code-2025/internal/input"
)

func TestPart1(t *testing.T) {
	lines := input.Content("test_input.txt")

	answer := part1(lines)

	if answer != 3 {
		t.Errorf("wrong answer, got: %d", answer)
	}
}
