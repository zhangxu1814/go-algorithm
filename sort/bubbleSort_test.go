package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{
			in:   []int{5, 4, 1, 2, 3},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			in:   []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			in:   []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			in:   []int{5, 6, 7, 9, 3, 5, 7, 2, 4, 2, 9, 6},
			want: []int{2, 2, 3, 4, 5, 5, 6, 6, 7, 7, 9, 9},
		},
	}

	for _, tt := range tests {
		out := BubbleSort(tt.in)
		for i, v := range out {
			if v != tt.want[i] {
				t.Errorf("got: %d, want: %d", v, tt.want[i])
			}
		}
	}
}
