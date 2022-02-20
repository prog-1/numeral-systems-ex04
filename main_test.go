package main

import "testing"

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		a    string
		b    string
		want string
	}{
		{"0", "0", "0"},
		{"123", "0", "0"},
		{"55555", "55", "3055525"},
		{"98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222",
			"2194787380219478738021947873802194787380219478737978052126197805212619780521261978052126197805212620"},
	} {
		got := LongMulWithoutLongAdd(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("ERR: LongMul(%s, %s): got: %s, want: %s", tc.a, tc.b, got, tc.want)
		}
	}
}

var (
	// benchA = "55555"
	// benchB = "55"
	benchA = "98765432109876543210987654321098765432109876543210"
	benchB = "22222222222222222222222222222222222222222222222222"
)

// To run all benchmarks use: go test -bench Long
// BenchmarkLongAdd is not needed, but interesting for comparison.
func BenchmarkLongAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LongAdd(benchA, benchB)
	}
}

func BenchmarkLongMulWithLongAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LongMulWithLongAdd(benchA, benchB)
	}
}

func BenchmarkLongMulWithoutLongAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LongMulWithoutLongAdd(benchA, benchB)
	}
}

// Conclusion: the program LongMul without function LongAdd works about 2 times faster.
