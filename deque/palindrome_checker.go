package deque

import (
	"strings"
	"unicode"
)

// PalindromeChecker returns true if a string is palindrome
func PalindromeChecker(s string) bool {

	if s == "" {
		return false
	}

	deque := NewDeque[rune]()
	lowerString := strings.ToLower(s)

	for _, char := range lowerString {
		if unicode.IsSpace(char) {
			continue
		}
		deque.AddBack(char)
	}

	for deque.Size() > 1 {
		firstChar, _ := deque.RemoveFront()
		lastChar, _ := deque.RemoveBack()

		if firstChar != lastChar {
			return false
		}
	}

	return true
}
