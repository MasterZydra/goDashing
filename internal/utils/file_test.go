package utils

import (
	"godashing"
	"path"
	"testing"
)

func TestExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr string
	}{
		{
			name: "Existing",
			args: args{path: path.Join(godashing.TestPath, "testfile.txt")},
			want: true,
			wantErr: "",
		},
		{
			name: "None existing",
			args: args{path: path.Join(godashing.TestPath, "noneexisitingtestfile.txt")},
			want: false,
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Exists(tt.args.path)
			if !ErrorContains(err, tt.wantErr) {
				t.Errorf("Exists() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exists() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
