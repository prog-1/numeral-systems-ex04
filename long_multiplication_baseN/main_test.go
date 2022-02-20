package main

import (
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
		{"1", "1", 6, "2"},
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

func benchmarkLongMul(a, c string, base int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		LongMul(a, c, base)
	}
}

func BenchmarkLongMul1(b *testing.B) { benchmarkLongMul("1", "1", 2, b) }
func BenchmarkLongMul2(b *testing.B) { benchmarkLongMul("3", "3", 4, b) }
func BenchmarkLongMul3(b *testing.B) { benchmarkLongMul("9", "9", 10, b) }
func BenchmarkLongMul4(b *testing.B) { benchmarkLongMul("123", "9", 10, b) }

// Conclusion: the higher the numbers, the more time the program needs to complete the task.
