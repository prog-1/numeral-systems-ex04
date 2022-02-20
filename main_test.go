package main

import "testing"

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		name string
		a, b string
		want string
	}{
		{"example2", "98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222",
			"2194787380219478738021947873802194787380219478737978052126197805212619780521261978052126197805212620"},
		{"empty", "", "", ""},
		{"zeros", "0", "0", "0"},
		{"in one decimal", "3", "2", "6"},
		{"bigger numbers", "40", "30", "1200"},
		{"example1", "55555", "55", "3055525"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongMul(tc.a, tc.b); got != tc.want {
				t.Errorf("LongMul(%v, %v) = %v, want = %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
func BenchmarkLongMul(b *testing.B) {
	// for the program optimization i used reconstructed Jaroslav programm.
	for n := 0; n < b.N; n++ {
		LongMul("58646", "5391")
	}
}
