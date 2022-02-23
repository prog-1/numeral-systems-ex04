package main

import (
	"fmt"
	"strings"
)

func NumberExistsInBaseN(num string, base int) bool {
	base36 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, v := range num {
		if !strings.Contains(base36[:base], string(v)) {
			return false
		}
	}
	return true
}

func LongMul(a, b string, base int) string {
	var s string
	m := make([]byte, len(a)+len(b))
	if 2 <= base && base <= 36 && NumberExistsInBaseN(a, base) && NumberExistsInBaseN(b, base) {
		for i, l := len(a)-1, 0; i >= 0; i-- {
			var add uint8 // uint8 = byte
			k := 0
			for j := len(b) - 1; j >= 0; j-- {
				var num byte
				if '0' <= a[i] && a[i] <= '9' && '0' <= b[j] && b[j] <= '9' {
					num = (a[i]-'0')*(b[j]-'0') + m[l+k] + add // '0' = 48
				} else {
					num = (a[i]-'7')*(b[j]-'7') + m[l+k] + add // '7' = 55
				}
				add = num / byte(base)
				m[l+k] = num % 10
				k++
			}
			if add != 0 {
				m[l+k] = m[l+k] + add
			}
			l++
		}
		for i := len(m) - 1; i >= 0; i-- {
			s = s + fmt.Sprint(m[i])
		}
		if s[0] == '0' {
			s = s[1:]
		}
	} else {
		s = "-1"
	}
	return s
}

func main() {
	var a, b string
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	var base int
	fmt.Print("Enter base: ")
	fmt.Scan(&base)
	if LongMul(a, b, base) != "-1" {
		fmt.Println("The product of a and b:", LongMul(a, b, base))
	} else {
		fmt.Println("ERROR: Incorrect base")
	}
}
