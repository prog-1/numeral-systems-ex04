package main

import (
	"strings"
	"testing"
)

var testLongMul = LongMul

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		want string
	}{
		{"0", "0", "0"},
		{"1", "0", "0"},
		{"0", "1", "0"},
		{"1", "7", "7"},
		{"12", "2", "24"},
		{"5", "5", "25"},
		{"1523", "327", "498021"},
		{"1523", "27", "41121"},
		{"2332323", "11", "25655553"},
		{"999", "1001", "999999"},

		{"55555", "55", "3055525"},
		{"98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222",
			"2194787380219478738021947873802194787380219478737978052126197805212619780521261978052126197805212620"},
	} {
		if got := testLongMul(tc.a, tc.b); got != tc.want {
			t.Errorf("LongSum(%v, %v) = %v, want = %v", tc.a, tc.b, got, tc.want)
		}
	}
}

var (
	benchA = strings.Repeat("1", 10000)
	benchB = strings.Repeat("9", 1)
)

func BenchmarkLongMul(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = testLongMul(benchA, benchB)
	}
}
