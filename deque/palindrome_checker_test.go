package deque

import "testing"

func Test_IsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "a simple character",
			input:    "a",
			expected: true,
		},
		{
			name:     "two characters",
			input:    "aa",
			expected: true,
		},
		{
			name:     "a word",
			input:    "kayak",
			expected: true,
		},
		{
			name:     "another word",
			input:    "level",
			expected: true,
		},
		{
			name:     "a setence",
			input:    "Step on no pets",
			expected: true,
		},
		{
			name:     "not palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "empty string",
			input:    "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := PalindromeChecker(tt.input)

			if got != tt.expected {
				t.Errorf(
					"fail in the case %s, expected %v, got %v for input %q",
					tt.name,
					tt.expected,
					got,
					tt.input,
				)
			}

		})
	}
}