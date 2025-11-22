package recursion

// IterativeFactorial calculates the factorial of a number using a loop
func IterativeFactorial(number int) int {
	if number < 0 {
		return 1
	}
	total := 1
	for n := number; n > 1; n-- {
		total = total * n
	}
	return total
}
