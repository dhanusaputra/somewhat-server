package jsonutil

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		want    interface{}
		wantErr bool
	}{
		{
			name: "happy path",
			file: "./testdata/success.json",
			want: map[string]interface{}{"id": map[string]interface{}{"nothing": "tidak ada yang personal di sini"}},
		},
		{
			name:    "open file failed",
			file:    "./testdata/fake.json",
			wantErr: true,
		},
		{
			name:    "unmarshal failed",
			file:    "./testdata/failed.json",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var in interface{}
			err := ReadFile(tt.file, &in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(in, tt.want) {
				t.Errorf("ReadFile() = %v, want = %v", in, tt.want)
			}
		})
	}
}
