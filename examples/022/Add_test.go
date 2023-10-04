package main

import "testing"

func TestAdd(t *testing.T) {
	// Test cases
	cases := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{-1, -1, -2},
	}
	// Iterate over test cases
	for _, c := range cases {
		got := Add(c.a, c.b)
		if got != c.want {
			t.Errorf("Add(%d, %d) == %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
