package main

import (
	"fmt"
	"strings"
)

func NumberExistsInBase10(num string) bool {
	base36 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, v := range num {
		if !strings.Contains(base36[:10], string(v)) {
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

func LongAdd(a, b string) string {
	var s string
	var add uint8 // uint8 = byte
	a, b = MakeNormal(a, b)

	for i := len(a) - 1; i >= 0; i-- {
		num := (a[i] - '0') + (b[i] - '0') + add // '0' = 48
		add = num / 10
		s = string((num%10)+'0') + s // '0' = 48
	}
	if add != 0 {
		s = "1" + s
	}
	return s
}

// Right now it only works correctly if both a and b or at least one of them is with len 1 (< 10)

func LongMul(a, b string) string {
	var s, result, sum string
	var add uint8 // uint8 = byte
	if len(a) < len(b) {
		a, b = b, a
	}
	if NumberExistsInBase10(a) && NumberExistsInBase10(b) {
		for i, k := len(a)-1, len(b)-1; i >= 0 && k >= 0; i-- {
			num := (a[i]-'0')*(b[k]-'0') + add // '0' = 48
			add = num / 10
			result = string(num%10+'0') + result // '0' = 48
			if i <= 0 {
				sum = LongAdd(string(sum), result)
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
	if LongMul(a, b) != "-1" {
		fmt.Println("The product of a and b:", LongMul(a, b))
	} else {
		fmt.Println("ERROR: At least one of the entered numbers does not exist in base 10")
	}
}
