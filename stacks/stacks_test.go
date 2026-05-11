package stacks

import "testing"

func TestStack_PushAndPop(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "push and pop one element",
			input: []int{5},
			want:  []int{5},
		},
		{
			name:  "push and pop four elements",
			input: []int{5, 8, 11, 15},
			want:  []int{15, 11, 8, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			stack := NewStack[int]()

			for _, element := range tt.input {
				stack.Push(element)
			}

			var got []int
			for !stack.IsEmpty() {
				got = append(got, stack.Pop())
			}

			if len(got) != len(tt.want) {
				t.Errorf("fail in the case %s, expected length %d, got %d", tt.name, len(tt.want), len(got))
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("expected %v, got %v", tt.want, got)
				}
			}

		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "peek an element",
			input:    []int{5},
			expected: 5,
		},
		{
			name:     "peek multiple elements",
			input:    []int{5, 8, 11, 15},
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			stack := NewStack[int]()

			for _, element := range tt.input {
				stack.Push(element)
			}

			top := stack.Peek()

			if top != tt.expected {
				t.Errorf("fail in the case %s, expected top %d, got %d", tt.name, tt.expected, top)
			}

			if len(stack.items) != len(tt.input) {
				t.Errorf(
					"fail in the case %s, peek should not remove element, expected size %d, got %d",
					tt.name, len(tt.input), len(stack.items))
			}

		})
	}
}

func TestStack_Peek_Empty(t *testing.T) {
	stack := NewStack[int]()

	got := stack.Peek()

	if got != 0 {
		t.Errorf("expected zero value, got %d", got)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "multiple elements",
			input:    []int{5, 8, 11, 15},
			expected: false,
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[int]()

			for _, element := range tt.input {
				stack.Push(element)
			}

			empty := stack.IsEmpty()

			if empty != tt.expected {
				t.Errorf("fail in the case %s, expected true, got %v", tt.name, empty)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "two elements",
			input:    []int{5, 8},
			expected: 2,
		},
		{
			name:     "multiple elements",
			input:    []int{5, 8, 11, 15},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			stack := NewStack[int]()

			for _, element := range tt.input {
				stack.Push(element)
			}

			size := stack.Size()

			if size != tt.expected {
				t.Errorf("fail in the case %s, expected %d, got %d", tt.name, tt.expected, size)
			}
		})
	}
}
