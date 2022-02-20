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

func AddZeros(s string, count int) string {
	for ; count > 0; count-- {
		s = s + "0"
	}
	return s
}

func MakeNormal(a, b string) (string, string) {
	i := 0
	if len(a) > len(b) {
		b = AddZeros(b, len(a)-len(b)+i)
	} else {
		a = AddZeros(a, len(b)-len(a)+i)
	}
	i++
	return a, b
}

func LongAdd(a, b string, base int) string {
	var s string
	var add uint8 // uint8 = byte
	a, b = MakeNormal(a, b)
	base36 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := len(a) - 1; i >= 0; i-- {

		num := (a[i] - '0') + (b[i] - '0') + add // '0' = 48
		add = num / uint8(base)
		if '0' <= uint8(base) && uint8(base) <= '9' {
			s = string((base36[num%uint8(base)])+'0') + s // '0' = 48
		} else {
			s = string((base36[num%uint8(base)])+'7') + s // '7' = 55
		}
	}
	if add != 0 {
		s = "1" + s
	}
	return s
}

func LongMul(a, b string, base int) string {
	var s, result, sum string
	var num, add uint8 // uint8 = byte
	base36 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if len(a) < len(b) {
		a, b = b, a
	}
	if 2 < base && base < 36 && NumberExistsInBaseN(a, base) && NumberExistsInBaseN(b, base) {
		for i, k := len(a)-1, len(b)-1; i >= 0 && k >= 0; i-- {
			if '0' <= a[i] && a[i] <= '9' && '0' <= b[k] && b[k] <= '9' {
				num = (a[i]-'0')*(b[k]-'0') + add // '0' = 48
			} else if 'A' <= a[i] && a[i] <= 'Z' && '0' <= b[k] && b[k] <= '9' {
				num = (a[i]-'7')*(b[k]-'0') + add // '7' = 55
			} else if '0' <= a[i] && a[i] <= '9' && 'A' <= b[k] && b[k] <= 'Z' {
				num = (a[i]-'0')*(b[k]-'7') + add
			} else {
				num = (a[i]-'7')*(b[k]-'7') + add
			}
			add = num / uint8(base)
			if '0' <= base && base <= '9' {
				result = string(base36[num%uint8(base)+'0']) + result // '0' = 48
			} else {
				result = string(base36[num%uint8(base)]+'7') + result // '7' = 55
			}
			if i <= 0 {
				sum = LongAdd(string(sum), result, base)
				k = k - 1
				i = len(a) - 1
				result = ""
			}
		}
		s = sum + s
		if add != 0 {
			s = string(add+'0') + s // '0' = 48
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
