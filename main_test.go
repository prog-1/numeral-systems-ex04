package main

import "testing"

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		name string
		a    string
		b    string
		want string
	}{
		{"empty", "", "", ""},
		{"test with 0", "0", "0", "0"},
		{"test with same numbers", "2", "2", "4"},
		{"simple test", "3", "4", "12"},
		{"test from README", "55555", "55", "3055525"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongMul(tc.a, tc.b); got != tc.want {
				t.Errorf("LongMul(%v, %v) = %v, want = %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
