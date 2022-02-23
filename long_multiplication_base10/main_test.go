package main

import (
	"strings"
	"testing"
)

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		want string
	}{
		{"0", "0", "0"},
		{"1", "0", "0"},
		{"0", "1", "0"},
		{"1", "1", "1"},
		{"-1", "1", "-1"},
		{"1", "-1", "-1"},
		{"-1", "-1", "-1"},
		{"2", "2", "4"},
		{"3", "3", "9"},
		{"4", "4", "16"},
		{"5", "6", "30"},
		{"10", "10", "100"},
		{"100", "1000", "100000"},
		{"123", "210", "25830"},
		{"123", "-210", "-1"},
		{"998", "12", "11976"},
		{"999", "1001", "999999"},
		{"55555", "55", "3055525"},
		{"2A", "391", "-1"},
		{"24671", "B145C7", "-1"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "0123456789", "-1"},
		{"1346595959131", "9", "12119363632179"},
		{"98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222",
			"2194787380219478738021947873802194787380219478737978052126197805212619780521261978052126197805212620"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if got := LongMul(tc.a, tc.b); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

var (
	benchA = strings.Repeat("1", 10000)
	benchB = strings.Repeat("9", 1)

	// benchA    = strings.Repeat("5", 10000)
	// benchB    = strings.Repeat("17", 2)

	// benchA = strings.Repeat("1234567890", 10000)
	// benchB = strings.Repeat("99999", 100000)
)

func BenchmarkLongSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = LongMul(benchA, benchB)
	}
}

// Conclusion: The higher the numbers, the more ns, B and allocs the program needs to complete the task.
