package main

import (
	"strings"
	"testing"
)

// testLongAdd allows switching an implementation easily. Though it would be
// better to create different benchmarks for different implementations to
// compare them.
var testLongAdd = LongAdd

func TestLongSum(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		want string
	}{
		{"0", "0", "0"},
		{"1", "0", "1"},
		{"0", "1", "1"},
		{"1", "9", "10"},
		{"9", "1", "10"},
		{"9", "9", "18"},
		{"5", "5", "10"},
		{"55", "66", "121"},
		{"11", "99", "110"},
		{"99", "11", "110"},
		{"999", "1001", "2000"},

		{"123", "210", "333"},
		{"55555", "55", "55610"},
		{"98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222",
			"120987654332098765433209876543320987654332098765432"},
	} {
		if got := testLongAdd(tc.a, tc.b); got != tc.want {
			t.Errorf("LongSum(%v, %v) = %v, want = %v", tc.a, tc.b, got, tc.want)
		}
	}
}

var (
	// benchA = strings.Repeat("1234567890", 10000)
	// benchB = strings.Repeat("99999", 100000)
	benchA = strings.Repeat("1", 10000)
	benchB = strings.Repeat("9", 1)
)

func BenchmarkLongSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = testLongAdd(benchA, benchB)
	}
}
