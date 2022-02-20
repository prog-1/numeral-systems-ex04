package main

import "testing"

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		testname string
		a        string
		b        string
		want     string
	}{
		{"all empty", "", "", ""},
		{"first number empty", "", "12312", ""},
		{"second number empty", "26712", "", ""},
		{"zeros", "0", "0", "0"},
		{"zero and number", "0", "2294311", "0"},
		{"number and zeror", "12675", "0", "0"},
		{"small numbers", "20", "50", "1000"},
		{"only one charecter mult", "55555", "55", "3055525"},
		{"very long numbers", "98765432109876543210987654321098765432109876543210", "22222222222222222222222222222222222222222222222222", "2194787380219478738021947873802194787380219478737978052126197805212619780521261978052126197805212620"},
	} {
		if got := LongMul(tc.a, tc.b); got != tc.want {
			t.Errorf("%v - ERROR: got = %v, want = %v", tc.testname, got, tc.want)
		}
	}
}
