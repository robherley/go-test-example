package highcov

import "strings"

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Sub returns the difference of two integers.
func Sub(a, b int) int {
	return a - b
}

// Mul returns the product of two integers.
func Mul(a, b int) int {
	return a * b
}

// Div returns the quotient of two integers and whether the operation was valid.
func Div(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// Reverse returns s with its characters in reverse order.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Shout returns s uppercased with an exclamation point.
func Shout(s string) string {
	return strings.ToUpper(s) + "!"
}

// FizzBuzz returns the classic FizzBuzz string for n.
func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return Shout(strings.TrimSpace(""))
	}
}
