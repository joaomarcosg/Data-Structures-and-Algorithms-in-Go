package queue

import "testing"

func TestQueue(t *testing.T) {

	type testcase struct {
		name  string
		input []string
		want  []string
	}

	tests := []testcase{
		{
			name:  "Enqueue",
			input: []string{"John", "Jack", "Camila"},
			want:  []string{"John", "Jack", "Camila"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			queue := NewQueue[string]()
			for _, element := range tc.input {
				queue.Enqueue(element)
			}

			front, ok := queue.FrontQueue()

			if !ok {
				t.Errorf("Expected front element %q, but queue was empty", tc.want[0])
			}

			if front != tc.want[0] {
				t.Errorf("Expected front element %q, but got %q", tc.want[0], front)
			}

		})
	}

}
