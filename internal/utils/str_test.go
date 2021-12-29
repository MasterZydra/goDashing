package utils

import "testing"

func TestEscapeString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "One \\n",
			args: args{input: "Test with one \n"},
			want: "Test with one ",
		},
		{
			name: "One \\r",
			args: args{input: "Test with one \r"},
			want: "Test with one ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EscapeString(tt.args.input)
			if got != tt.want {
				t.Errorf("EscapeString() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
