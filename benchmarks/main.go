package main

import "fmt"

func LongAdd(a, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	var carry byte
	res := make([]byte, len(b)+1)
	for i := range a {
		sum := carry + a[len(a)-1-i] - '0' + b[len(b)-1-i] - '0'
		res[i], carry = sum%10+'0', sum/10
	}
	for i := len(a); i < len(b); i++ {
		sum := carry + b[len(b)-1-i] - '0'
		res[i], carry = sum%10+'0', sum/10
	}
	if carry > 0 {
		res[len(res)-1] = carry + '0'
	} else {
		res = res[:len(res)-1]
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func main() {
	for _, t := range []struct {
		a, b string
	}{
		{"0", "0"},
		{"1", "0"},
		{"0", "1"},
		{"1", "9"},
		{"9", "1"},
		{"11", "99"},
		{"99", "11"},
		{"999", "1001"},
	} {
		fmt.Printf("%v + %v = %v\n", t.a, t.b, LongAdd(t.a, t.b))
	}
}
