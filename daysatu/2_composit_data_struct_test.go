package daysatu

import "testing"

func TestStructType(t *testing.T) {
	tests := []struct {
		name string
	}{}

	t.Log(tests)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StructType()
		})
	}
}
