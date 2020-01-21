package main

import "strconv"

func main() {
	fizzbuzz(5)
}

func fizzbuzz(n int) string {

	var name string

	if n%3 == 0 {
		name = "Fizz"
	}
	if n%5 == 0 {
		name += "Buzz"
	}
	if name == "" {
		name = strconv.Itoa(n)
	}

	return name
}
