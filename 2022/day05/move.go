package main

import "strconv"

type move struct {
	From int
	To   int
	Qty  int
}

func (rv *move) SetFrom(s string) *move {
	rv.From, _ = strconv.Atoi(s)
	rv.From -= 1

	return rv
}

func (rv *move) SetTo(s string) *move {
	rv.To, _ = strconv.Atoi(s)
	rv.To -= 1

	return rv
}

func (rv *move) SetQty(s string) *move {
	rv.Qty, _ = strconv.Atoi(s)

	return rv
}
