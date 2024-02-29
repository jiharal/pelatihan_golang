package daysatu

import "testing"

func TestMasukkan_Bagi(t *testing.T) {
	tests := []struct {
		name string
		c    Masukkan
		want float64
	}{
		{
			name: "Test 1",
			c: Masukkan{
				prm1: 10,
				prm2: 10,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Bagi(); got != tt.want {
				t.Errorf("Masukkan.Bagi() = %v, want %v", got, tt.want)
			}
		})
	}
}
