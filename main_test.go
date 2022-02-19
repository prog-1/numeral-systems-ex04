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
		{"simple", "2", "2", "4"},
		{"from README", "55555", "55", "3055525"},
		{"from README", "98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222",
			"2194787380219478738021947873802194787380219478737978052126197805212619780521261978052126197805212620"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongMul(tc.a, tc.b); got != tc.want {
				t.Errorf("LongMul(%v, %v) = %v, want = %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
