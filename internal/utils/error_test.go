package utils

import (
	"errors"
	"testing"
)

func TestErrorContains(t *testing.T) {
	type args struct {
		err  error
		want string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Is nil",
			args: args{
				err:  nil,
				want: "",
			},
			want: true,
		},
		{
			name: "Should be nil but is not",
			args: args{
				err:  errors.New("Something went wrong"),
				want: "",
			},
			want: false,
		},
		{
			name: "Should be error but is not",
			args: args{
				err:  nil,
				want: "Something went wrong",
			},
			want: false,
		},
		{
			name: "Should be error and matches",
			args: args{
				err:  errors.New("Something went wrong"),
				want: "Something went wrong",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorContains(tt.args.err, tt.args.want); got != tt.want {
				t.Errorf("ErrorContains() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
