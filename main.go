package main

import (
	"fmt"
	"strings"
)

//took Sebastian's remake of the program from the internet, but corrected bug and made it a little faster

func LongMul(a, b string) string {
	ab := make([]byte, len(a)+len(b))
	ia := 0
	i := len(a) - 1
	for ; i >= 0; i-- {
		var carry byte
		ib := 0
		for j := len(b) - 1; j >= 0; j-- {
			mul := (a[i]-'0')*(b[j]-'0') + ab[ia+ib] + carry
			carry = mul / 10
			ab[ia+ib] = mul % 10
			ib++
		}
		if carry > 0 {
			ab[ia+ib] += carry
		}
		ia++
	}
	var res string
	for i := len(ab) - 1; i >= 0; i-- {
		res += fmt.Sprint(ab[i])
	}
	if strings.IndexAny(res, "0") == 0 {
		res = strings.TrimPrefix(res, "0")
	}
	return res
}

func main() {
	var a, b string
	fmt.Println("Examples:")
	for _, t := range []struct {
		a, b string
	}{
		{"0", "0"},
		{"1", "0"},
		{"0", "1"},
		{"2", "7"},
		{"7", "1"},
		{"12", "2"},
		{"5555", "55"},
		{"999", "1001"},
	} {
		fmt.Printf("%v * %v = %v\n", t.a, t.b, LongMul(t.a, t.b))
	}
	fmt.Println("Enter two numbers:")
	fmt.Scan(&a, &b)
	fmt.Printf("%v * %v = %v\n", a, b, LongMul(a, b))
}
