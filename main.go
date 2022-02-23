package main

import (
	"fmt"
	"strings"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsValidNumber(num string, base int) bool {
	for _, v := range num {
		if !strings.Contains(base36[:base], string(v)) {
			return false
		}
	}
	return true
}

func LongAdd(a, b string) (sum string) {
	var carry byte
	if len(a) > len(b) {
		b = strings.Repeat("0", len(a)-len(b)) + b
	} else {
		a = strings.Repeat("0", len(b)-len(a)) + a
	}
	for i := len(a) - 1; i >= 0; i-- {
		num := (a[i] - '0') + (b[i] - '0') + carry
		carry = num / 10
		sum = string(num%10+'0') + sum
	}
	if carry != 0 {
		sum = "1" + sum
	}
	return sum
}
func checkForZero(a, b string) bool {
	if a == "0" || b == "0" {
		return true
	}
	return false
}

func LongMul(a, b string) string {
	var result string
	if checkForZero(a, b) {
		return "0"
	}
	if len(a) > len(b) {
		a, b = b, a
	}
	toSum := make([][]byte, len(a))
	for i := range a {
		var carry uint8
		toSum[i] = make([]byte, len(b)+1)
		for i1 := range b {
			sum := carry + (a[len(a)-1-i]-'0')*(b[len(b)-1-i1]-'0')
			toSum[i][len(toSum[i])-1-i1], carry = sum%10+'0', sum/10
		}
		if carry > 0 {
			toSum[i][0] = carry + '0'
		} else {
			toSum[i] = toSum[i][1:]
		}
	}
	for i := range toSum {
		for j := i; j > 0; j-- {
			toSum[i] = append(toSum[i], '0')
		}
	}
	for _, v := range toSum {
		result = LongAdd(result, string(v))
	}
	return result

}

func main() {
	var a, b string
	base := 10
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)
	fmt.Println(LongMul(a, b))
	if IsValidNumber(a, base) && IsValidNumber(b, base) {
		fmt.Printf("Sum of %v and %v = %v", a, b, LongMul(a, b))
	} else {
		fmt.Printf("The number %v or number %v is represented incorrect in base%v", a, b, base)
	}
}
