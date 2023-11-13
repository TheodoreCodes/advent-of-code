package main

import "testing"

const testInput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func Test_part1(t *testing.T) {
	t.Run("day 05 part 1", func(t *testing.T) {
		expected := "CMZ"

		if actual := part1(testInput); actual != expected {
			t.Errorf("day 05 part 1: expected: %s, actual %s", expected, actual)
		}
	})
}

func Test_part2(t *testing.T) {
	t.Run("day 05 part 2", func(t *testing.T) {
		expected := "MCD"

		if actual := part2(testInput); actual != expected {
			t.Errorf("day 05 part 2: expected: %s, actual %s", expected, actual)
		}
	})
}
