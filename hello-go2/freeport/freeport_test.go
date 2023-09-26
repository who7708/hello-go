// Taken from https://github.com/phayes/freeport

package freeport

import (
	"reflect"
	"testing"
)

func TestGetFreePort(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Hellow", want: 8080, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFreePort()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFreePort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFreePort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFreePorts(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Hello", args: args{count: 10}, want: []int{1, 2, 34}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFreePorts(tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFreePorts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFreePorts() = %v, want %v", got, tt.want)
			}
		})
	}
}
