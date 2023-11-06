package main

import (
	"advent-of-code/scripts/skeleton"
	"flag"
	"time"
)

func main() {
	year := flag.Int("year", time.Now().Year(), "Advent of Code year to fetch")
	sessionCookie := flag.String("session", "", "Session cookie obtained after logging in via https://adventofcode.com")
	flag.Parse()

	skeleton.Run(*year, *sessionCookie)
}
