package main

import (
	"strings"
	"testing"
)

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		base int
		want string
	}{
		{"0", "0", 0, "-1"},
		{"0", "0", 1, "-1"},
		{"0", "0", 2, "0"},
		{"0", "0", 10, "0"},
		{"0", "0", 16, "0"},
		{"1", "0", 3, "0"},
		{"0", "1", 21, "0"},
		{"1", "1", 5, "1"},
		{"1", "2", 6, "2"},
		{"6", "6", 7, "36"},
		{"6", "6", 6, "-1"},
		{"6", "6", 37, "-1"},
		{"55555", "55", 10, "3055525"},
		{"123", "-210", 10, "-1"},
		{"1535890", "7", 8, "-1"},
		{"2A", "39B1", 11, "-1"},
		{"98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222",
			16,
			"14540b39e014540b39e014540b39e014540b39e014540b39dfebabf4c61febabf4c61febabf4c61febabf4c61febabf4c620"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if got := LongMul(tc.a, tc.b, tc.base); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

var (
	benchA    = strings.Repeat("1", 10000)
	benchB    = strings.Repeat("9", 1)
	benchBase = 10

	// benchA    = strings.Repeat("5", 10000)
	// benchB    = strings.Repeat("17", 2)
	// benchBase = 4

	// benchA = strings.Repeat("1234567890", 10000)
	// benchB = strings.Repeat("99999", 100000)
	// benchBase = 28
)

func BenchmarkLongSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = LongMul(benchA, benchB, benchBase)
	}
}

// Conclusion: The higher the numbers, the more ns, B and allocs the program needs to complete the task, increasing the base doesn't really affect anything, but decreasing it helps the program complete the task in less ns, B and allocs.
