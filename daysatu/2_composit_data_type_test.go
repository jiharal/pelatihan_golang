package daysatu

import "testing"

func TestArrayType(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ArrayType()
		})
	}
}

func TestSliceType(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "TestSliceType",
		},
		{
			name: "TestSliceTypeV1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceType()
		})
	}
}

func TestMapType(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "jello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MapType()
		})
	}
}
