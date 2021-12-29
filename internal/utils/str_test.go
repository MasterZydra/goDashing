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
		{
			name: "Four \\n",
			args: args{input: "\nTest \nwith \none \n"},
			want: "Test with one ",
		},
		{
			name: "Four \\r",
			args: args{input: "\rTest \rwith \rone \r"},
			want: "Test with one ",
		},
		{
			name: "HTML",
			args: args{input: "<a></a>"},
			want: "&lt;a&gt;&lt;/a&gt;",
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
