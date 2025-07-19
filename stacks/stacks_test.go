package stacks

import "testing"

// TestStack for testing the stack with array data structure
func TestStack(t *testing.T) {

	type testCase struct {
		name  string
		input []int
		want  []int
	}

	tests := []testCase{
		{
			name:  "Push and Pop 4 elements",
			input: []int{5, 8, 11, 15},
			want:  []int{15, 11, 8, 5},
		},
		{
			name:  "Push and Pop no elements",
			input: []int{},
			want:  []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stack := NewStack[int]()
			for _, element := range tc.input {
				stack.Push(element)
			}

			var got []int
			for !stack.IsEmpty() {
				got = append(got, stack.Pop())
			}

			if len(got) != len(tc.want) {
				t.Errorf("expected length %d, got %d", len(tc.want), len(got))
			}

			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("expected %v, got %v", tc.want, got)
				}
			}
		})
	}

}
