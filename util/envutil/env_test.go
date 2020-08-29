package envutil

import (
	"reflect"
	"testing"

	"github.com/dhanusaputra/somewhat-server/util/testutil"
)

func TestGetEnv(t *testing.T) {
	type args struct {
		key        string
		defaultVal string
	}
	tests := []struct {
		name string
		args args
		want string
		mock func()
	}{
		{
			name: "happy path",
			args: args{
				key:        "MOCK",
				defaultVal: "mock",
			},
			want: "mock",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&GetEnv}).Restore()
			if tt.mock != nil {
				tt.mock()
			}
			got := GetEnv(tt.args.key, tt.args.defaultVal)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnv() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvAsInt(t *testing.T) {
	type args struct {
		key        string
		defaultVal int
	}
	tests := []struct {
		name string
		args args
		want int
		mock func()
	}{
		{
			name: "happy path - default",
			args: args{
				key:        "MOCK",
				defaultVal: 1,
			},
			want: 1,
			mock: func() {
				GetEnv = func(key string, defaultVal string) string {
					return ""
				}
			},
		},
		{
			name: "happy path",
			args: args{
				key:        "MOCK",
				defaultVal: 1,
			},
			want: 2,
			mock: func() {
				GetEnv = func(key string, defaultVal string) string {
					return "2"
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&GetEnv}).Restore()
			if tt.mock != nil {
				tt.mock()
			}
			got := GetEnvAsInt(tt.args.key, tt.args.defaultVal)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnvAsInt() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvAsBool(t *testing.T) {
	type args struct {
		key        string
		defaultVal bool
	}
	tests := []struct {
		name string
		args args
		want bool
		mock func()
	}{
		{
			name: "happy path - default",
			args: args{
				key:        "MOCK",
				defaultVal: true,
			},
			want: true,
			mock: func() {
				GetEnv = func(key string, defaultVal string) string {
					return ""
				}
			},
		},
		{
			name: "happy path",
			args: args{
				key:        "MOCK",
				defaultVal: true,
			},
			want: false,
			mock: func() {
				GetEnv = func(key string, defaultVal string) string {
					return "false"
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&GetEnv}).Restore()
			if tt.mock != nil {
				tt.mock()
			}
			got := GetEnvAsBool(tt.args.key, tt.args.defaultVal)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnvAsBool() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvAsSlice(t *testing.T) {
	type args struct {
		key        string
		defaultVal []string
	}
	tests := []struct {
		name string
		args args
		want []string
		mock func()
	}{
		{
			name: "happy path - default",
			args: args{
				key:        "MOCK",
				defaultVal: []string{"1", "2"},
			},
			want: []string{"1", "2"},
			mock: func() {
				GetEnv = func(key string, defaultVal string) string {
					return ""
				}
			},
		},
		{
			name: "happy path",
			args: args{
				key:        "MOCK",
				defaultVal: []string{"1", "2"},
			},
			want: []string{"3", "4"},
			mock: func() {
				GetEnv = func(key string, defaultVal string) string {
					return "3,4"
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&GetEnv}).Restore()
			if tt.mock != nil {
				tt.mock()
			}
			got := GetEnvAsSlice(tt.args.key, tt.args.defaultVal, ",")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnvAsSlice() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvAsMapBool(t *testing.T) {
	type args struct {
		key        string
		defaultVal map[string]bool
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
		mock func()
	}{
		{
			name: "happy path - default",
			args: args{
				key:        "MOCK",
				defaultVal: map[string]bool{"1": true, "2": true},
			},
			want: map[string]bool{"1": true, "2": true},
			mock: func() {
				GetEnv = func(key string, defaultVal string) string {
					return ""
				}
			},
		},
		{
			name: "happy path",
			args: args{
				key:        "MOCK",
				defaultVal: map[string]bool{"1": true, "2": true},
			},
			want: map[string]bool{"3": true, "4": true},
			mock: func() {
				GetEnv = func(key string, defaultVal string) string {
					return "3,4"
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&GetEnv}).Restore()
			if tt.mock != nil {
				tt.mock()
			}
			got := GetEnvAsMapBool(tt.args.key, tt.args.defaultVal, ",")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnvAsMapBool() = %v, want = %v", got, tt.want)
			}
		})
	}
}
