package main

import (
	"testing"
)

type obj struct {
	test *int
}

func TestPoint(t *testing.T) {
	a := []int{1, 2, 3}
	b := make([]int, len(a))
	copy(b, a)
	a[0] = 2
	t.Log(b[0])
}
