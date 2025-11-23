package recursion

// Factorial calculates the factorial of a number using recursion
func Factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
