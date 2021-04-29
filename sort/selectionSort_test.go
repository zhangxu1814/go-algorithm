package sort

import (
	"reflect"
	"testing"
)

func Test_selectionSort(t *testing.T) {
	type args struct {
		sequence []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{[]int{5, 6, 7, 9, 3, 5, 7, 2, 4, 2, 9, 6}},
			want: []int{2, 2, 3, 4, 5, 5, 6, 6, 7, 7, 9, 9},
		},
		{
			name: "case2",
			args: args{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.sequence); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
