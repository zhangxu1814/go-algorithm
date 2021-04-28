package square_sum

import "testing"

func Test_judgeSquareSum(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{5},
			want: true,
		},
		{
			name: "case1",
			args: args{4},
			want: true,
		},
		{
			name: "case1",
			args: args{3},
			want: false,
		},
		{
			name: "case1",
			args: args{2},
			want: true,
		},
		{
			name: "case1",
			args: args{1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSum(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
