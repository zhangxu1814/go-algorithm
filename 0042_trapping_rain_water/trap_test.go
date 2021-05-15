package trapping_rain_water

import (
	"testing"
)

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
		{
			name: "case2",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			want: 9,
		},
		{
			name: "case3",
			args: args{
				height: []int{1, 2},
			},
			want: 0,
		},
		{
			name: "case4",
			args: args{
				height: []int{},
			},
			want: 0,
		},
		{
			name: "case5",
			args: args{
				height: []int{0, 5, 6, 4, 6, 1, 0, 0, 2, 7},
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapBF(tt.args.height); got != tt.want {
				t.Errorf("trapBF() = %v, want %v", got, tt.want)
			}
			if got := trapDP(tt.args.height); got != tt.want {
				t.Errorf("trapDP() = %v, want %v", got, tt.want)
			}
			if got := trapDP2(tt.args.height); got != tt.want {
				t.Errorf("trapDP2() = %v, want %v", got, tt.want)
			}
		})
	}
}
