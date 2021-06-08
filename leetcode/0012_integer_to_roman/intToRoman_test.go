package integer_to_roman

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{3},
			want: "III",
		},
		{
			name: "case2",
			args: args{4},
			want: "IV",
		},
		{
			name: "case3",
			args: args{9},
			want: "IX",
		},
		{
			name: "case4",
			args: args{58},
			want: "LVIII",
		},
		{
			name: "case5",
			args: args{1994},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
