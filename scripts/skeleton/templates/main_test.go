package main

import "testing"

const testInput = ``

func Test_part1(t *testing.T) {
	t.Run("day {{.Day}} part 1", func(t *testing.T) {
		expected := 0

		if actual := part1(testInput); actual != expected {
			t.Errorf("day {{.Day}} part 1: expected: %d, actual %d", expected, actual)
		}
	})
}

func Test_part2(t *testing.T) {
	t.Run("day {{.Day}} part 2", func(t *testing.T) {
		expected := 0

		if actual := part1(testInput); actual != expected {
			t.Errorf("day {{.Day}} part 2: expected: %d, actual %d", expected, actual)
		}
	})
}
