package util

import (
	"testing"
)

func TestLogGoRoutineCount(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogGoRoutineCount()
		})
	}
}

func TestLogGoroutineStackTrace(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogGoroutineStackTrace()
		})
	}
}
