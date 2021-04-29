package next_greater_element_II

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElementII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 1},
			},
			want: []int{2, -1, 2},
		},
		{
			name: "case1",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: []int{-1, 5, 5, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElementII(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElementII() = %v, want %v", got, tt.want)
			}
		})
	}
}
