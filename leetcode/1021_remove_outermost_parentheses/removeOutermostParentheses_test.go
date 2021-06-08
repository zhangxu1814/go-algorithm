package remove_outermost_parentheses

import "testing"

func Test_removeOutermostParenthese(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{"(()())(())"},
			want: "()()()",
		},
		{
			name: "case2",
			args: args{"(()())(())(()(()))"},
			want: "()()()()(())",
		},
		{
			name: "case3",
			args: args{"()()"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOutermostParenthese(tt.args.S); got != tt.want {
				t.Errorf("removeOutermostParenthese() = %v, want %v", got, tt.want)
			}
		})
	}
}
