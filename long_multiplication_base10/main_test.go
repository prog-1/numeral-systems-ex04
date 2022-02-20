package main

import (
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

func benchmarkLongMul(a, c string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		LongMul(a, c)
	}
}

func BenchmarkLongMul1(b *testing.B) { benchmarkLongMul("1", "1", b) }
func BenchmarkLongMul2(b *testing.B) { benchmarkLongMul("3", "3", b) }
func BenchmarkLongMul3(b *testing.B) { benchmarkLongMul("9", "9", b) }
func BenchmarkLongMul4(b *testing.B) { benchmarkLongMul("123", "9", b) }

// Conclusion: the higher the numbers, the more time the program needs to complete the task.
