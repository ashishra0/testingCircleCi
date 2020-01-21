package main

import "testing"

func TestFizzbuzz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
	}

	for _, tt := range cases {
		actual := fizzbuzz(tt.in)
		if actual != tt.want {
			t.Errorf("Fizzbuzz(%d): expected %q, actual %q", tt.in, tt.want, actual)
		}
	}

}
