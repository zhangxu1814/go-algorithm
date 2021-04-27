package sort

import (
	"reflect"
	"testing"
)

func Test_insertionSort(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSort(tt.args.sequence); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
