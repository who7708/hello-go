package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{s: "a:b:c", sep: ":"}, wantResult: []string{"a", "b", "c"}},
		{name: "test2", args: args{s: "a:b:c", sep: ","}, wantResult: []string{"a:b:c"}},
		{name: "test3", args: args{s: "abcd", sep: "bc"}, wantResult: []string{"a", "d"}},
		{name: "test4", args: args{s: "沙河有沙又有河", sep: "沙"}, wantResult: []string{"河有", "又有河"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Split() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
