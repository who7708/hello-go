package section05

import (
	"testing"
)

func TestSetConfig(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Hello", args: args{file: "test.txt"}},
		{name: "World", args: args{file: "test.cfg"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetConfig(tt.args.file)
		})
	}
}

func TestSetHomeDir(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "Hello", args: args{root: "./"}},
		{name: "World", args: args{root: "/a/b/c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetHomeDir(tt.args.root)
		})
	}
}

func Test_path_Resolve(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		p    *path
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Test1", p: &path{homeDir: "/", configFile: "test.a"}, args: args{path: "/test"}, want: "/test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Resolve(tt.args.path); got != tt.want {
				t.Errorf("path.Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
