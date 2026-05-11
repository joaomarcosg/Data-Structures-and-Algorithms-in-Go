package deque

import (
	"reflect"
	"testing"
)

func TestDeque_AddFrontAndRemoveFront(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "one element",
			input: []string{"John"},
			want:  []string{"John"},
		},
		{
			name:  "multiple elements",
			input: []string{"John", "Jack", "Camila"},
			want:  []string{"Camila", "Jack", "John"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			deque := NewDeque[string]()

			for _, element := range tt.input {
				deque.AddFront(element)
			}

			var got []string
			for !deque.IsEmpty() {
				item, ok := deque.RemoveFront()
				if !ok {
					t.Fatal("unexpected empty deque")
				}
				got = append(got, item)
			}

			if len(got) != len(tt.want) {
				t.Errorf("fail in the case %s, expected length %d, got %d", tt.name, len(tt.want), len(got))
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected %s, got %s", tt.want, got)
			}

		})
	}
}

func TestDeque_RemoveFront_EmptyDeque(t *testing.T) {
	deque := NewDeque[int]()

	_, ok := deque.RemoveFront()

	if ok {
		t.Error("expected false when removing from empty deque")
	}
}

func TestDeque_MixedOpetations(t *testing.T) {
	deque := NewDeque[string]()

	deque.AddFront("A")
	deque.AddFront("B")

	item, _ := deque.RemoveFront()

	if item != "B" {
		t.Errorf("expected B, got %s", item)
	}

	deque.AddFront("C")

	first, _ := deque.RemoveFront()
	second, _ := deque.RemoveFront()

	if first != "C" || second != "A" {
		t.Errorf("expected C and A, got %v and %v", first, second)
	}
}

func TestDeque_AddBackAndRemoveBack(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "one element",
			input: []string{"John"},
			want:  []string{"John"},
		},
		{
			name:  "multiple elements",
			input: []string{"John", "Jack", "Camila"},
			want:  []string{"Camila", "Jack", "John"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			deque := NewDeque[string]()

			for _, element := range tt.input {
				deque.AddBack(element)
			}

			var got []string
			for !deque.IsEmpty() {
				item, ok := deque.RemoveBack()
				if !ok {
					t.Fatal("unexpected empty deque")
				}
				got = append(got, item)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected %v, got %v", tt.want, got)
			}

			if !deque.IsEmpty() {
				t.Error("expected deque to be empty")
			}

		})
	}
}

func TestDeque_RemoveBack_Empty(t *testing.T) {
	deque := NewDeque[string]()

	_, ok := deque.RemoveBack()

	if ok {
		t.Error("expected false when removing from empty deque")
	}

}

func TestDeque_MixedBackOperations(t *testing.T) {
	deque := NewDeque[string]()

	deque.AddBack("A")
	deque.AddBack("B")

	item, _ := deque.RemoveBack()

	if item != "B" {
		t.Errorf("expected B, got %v", item)
	}

	deque.AddBack("C")

	first, _ := deque.RemoveBack()
	second, _ := deque.RemoveBack()

	if first != "C" || second != "A" {
		t.Errorf("expected C and A, got %v and %v", first, second)
	}
}

func TestDeque_FrontDeque_AddFront(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "deque with one element",
			input: []string{"John"},
			want:  "John",
		},
		{
			name:  "deque with multiple elements",
			input: []string{"John", "Jack", "Camila"},
			want:  "Camila",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			deque := NewDeque[string]()

			for _, element := range tt.input {
				deque.AddFront(element)
			}

			front, ok := deque.FrontDeque()

			if !ok {
				t.Error("unexpected empty deque")
			}

			if front != tt.want {
				t.Errorf("fail in the case %s, expected front %s, got %s", tt.name, tt.want, front)
			}

		})
	}
}

func TestDeque_FrontDeque_AddBack(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "deque with one element",
			input: []string{"John"},
			want:  "John",
		},
		{
			name:  "deque with multiple elements",
			input: []string{"John", "Jack", "Camila"},
			want:  "John",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			deque := NewDeque[string]()

			for _, element := range tt.input {
				deque.AddBack(element)
			}

			front, ok := deque.FrontDeque()

			if !ok {
				t.Error("unexpected empty deque")
			}

			if front != tt.want {
				t.Errorf("fail in the case %s, expected front %s, got %s", tt.name, tt.want, front)
			}

		})
	}
}

func TestDeque_RearDeque_AddFront(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "deque with one element",
			input: []string{"John"},
			want:  "John",
		},
		{
			name:  "deque with multiple elements",
			input: []string{"John", "Jack", "Camila"},
			want:  "John",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			deque := NewDeque[string]()

			for _, element := range tt.input {
				deque.AddFront(element)
			}

			rear, ok := deque.RearDeque()

			if !ok {
				t.Error("unexpected empty deque")
			}

			if rear != tt.want {
				t.Errorf("fail in the case %s, expected front %s, got %s", tt.name, tt.want, rear)
			}

		})
	}
}

func TestDeque_RearDeque_AddBack(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "deque with one element",
			input: []string{"John"},
			want:  "John",
		},
		{
			name:  "deque with multiple elements",
			input: []string{"John", "Jack", "Camila"},
			want:  "Camila",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			deque := NewDeque[string]()

			for _, element := range tt.input {
				deque.AddBack(element)
			}

			rear, ok := deque.RearDeque()

			if !ok {
				t.Error("unexpected empty deque")
			}

			if rear != tt.want {
				t.Errorf("fail in the case %s, expected front %s, got %s", tt.name, tt.want, rear)
			}

		})
	}
}

func TestDeque_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected bool
	}{
		{
			name:     "empty deque",
			input:    []string{},
			expected: true,
		},
		{
			name:     "multiple elements",
			input:    []string{"John", "Jack", "Camila"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			deque := NewDeque[string]()

			for _, element := range tt.input {
				deque.AddBack(element)
			}

			empty := deque.IsEmpty()

			if empty != tt.expected {
				t.Errorf("fail in the case %s, expected true, got %v", tt.name, empty)
			}

		})
	}

}
