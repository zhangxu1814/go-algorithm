package algorithms4th

import "testing"

func Test_evaluate(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				s: []string{"(", "1", "+", "(", "(", "2", "+", "3", ")", "*", "(", "4", "*", "5", ")", ")", ")"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluate(tt.args.s)
		})
	}
}
