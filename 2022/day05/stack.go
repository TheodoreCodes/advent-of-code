package main

type stack[T comparable] struct {
	data []T
}

func (rv *stack[T]) Push(elem T) {
	rv.data = append(rv.data, elem)
}

func (rv *stack[T]) PushSlice(elems []T) {
	rv.data = append(rv.data, elems...)
}

func (rv *stack[T]) PushBottom(elem T) {
	rv.data = append([]T{elem}, rv.data...)
}

func (rv *stack[T]) Pop() T {
	var zero T
	l := len(rv.data)

	if l == 0 {
		return zero
	}

	e := rv.data[l-1]
	rv.data = rv.data[:l-1]

	return e
}

func (rv *stack[T]) PopSlice(n int) []T {
	l := len(rv.data)

	if l == 0 {
		return []T{}
	}

	elems := rv.data[l-n:]
	rv.data = rv.data[:l-n]

	return elems
}
