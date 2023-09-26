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
