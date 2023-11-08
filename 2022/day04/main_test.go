package main

import "testing"

const testInput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func Test_part1(t *testing.T) {
	t.Run("day 04 part 1", func(t *testing.T) {
		expected := 2

		if actual := part1(testInput); actual != expected {
			t.Errorf("day 04 part 1: expected: %d, actual %d", expected, actual)
		}
	})
}

func Test_part2(t *testing.T) {
	t.Run("day 04 part 2", func(t *testing.T) {
		expected := 4

		if actual := part2(testInput); actual != expected {
			t.Errorf("day 04 part 2: expected: %d, actual %d", expected, actual)
		}
	})
}
