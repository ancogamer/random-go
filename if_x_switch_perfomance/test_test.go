package test

import "testing"

func TestIf(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"test_if"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			If()
		})
	}
}

func BenchmarkIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		If()
	}
}
func TestSwich(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"test_swicht_"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Swich()
		})
	}
}

func BenchmarkSwich(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Swich()
	}
}
func TestSwich1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"test_swicht_1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Swich1()
		})
	}
}

func BenchmarkSwich1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Swich1()
	}
}
