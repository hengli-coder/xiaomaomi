package example

import (
	"testing"
)

func TestExampleMutex(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "ok"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_or(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "ok"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			or()
		})
	}
}