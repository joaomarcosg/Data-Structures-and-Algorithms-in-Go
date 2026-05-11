package recursion

func IterativeFibonacci(n int) int {
	if n < 1 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	fibNMinus2 := 0
	fibNMinus1 := 1
	fibN := n

	for i := 2; i <= n; i++ {
		fibN = fibNMinus1 + fibNMinus2
		fibNMinus2 = fibNMinus1
		fibNMinus1 = fibN
	}

	return fibN
}
