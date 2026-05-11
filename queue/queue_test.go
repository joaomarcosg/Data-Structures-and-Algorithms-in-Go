package queue

import (
	"reflect"
	"testing"
)

func TestQueue_EnqueueAndDequeue(t *testing.T) {
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
			want:  []string{"John", "Jack", "Camila"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			queue := NewQueue[string]()

			for _, element := range tt.input {
				queue.Enqueue(element)
			}

			var got []string
			for !queue.IsEmpty() {
				item, _ := queue.Dequeue()
				got = append(got, item)
			}

			if len(got) != len(tt.want) {
				t.Errorf("fail in the case %s, expected length %d, got %d", tt.name, len(tt.want), len(got))
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected %v, got %v", tt.want, got)
			}

		})
	}
}

func TestQueue_DequeueEmpty(t *testing.T) {
	queue := NewQueue[string]()

	got, ok := queue.Dequeue()

	if got != "" && !ok {
		t.Errorf("expected zero value, got %v", got)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	queue := NewQueue[int]()

	if !queue.IsEmpty() {
		t.Error("expected queue to be empty")
	}

	queue.Enqueue(1)

	if queue.IsEmpty() {
		t.Error("expected queue to not be empty")
	}
}

func TestQueue_FrontQueue(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "one element",
			input: []string{"John"},
			want:  "John",
		},
		{
			name:  "multiple elements",
			input: []string{"John", "Jack", "Camila"},
			want:  "John",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			queue := NewQueue[string]()

			for _, element := range tt.input {
				queue.Enqueue(element)
			}

			first, ok := queue.FrontQueue()

			if first != tt.want && !ok {
				t.Errorf("fail in the case %s, expected first %s, got %s", tt.name, tt.want, first)
			}

		})
	}
}

func TestQueue_Size(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "empty queue",
			input: []string{},
			want:  0,
		},
		{
			name:  "one element",
			input: []string{"John"},
			want:  1,
		},
		{
			name:  "multiple elements",
			input: []string{"John", "Jack", "Camila"},
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			queue := NewQueue[string]()

			for _, element := range tt.input {
				queue.Enqueue(element)
			}

			size := queue.Size()

			if size != tt.want {
				t.Errorf("fail in the case %s, expected %d, got %d", tt.name, tt.want, size)
			}

		})
	}
}
