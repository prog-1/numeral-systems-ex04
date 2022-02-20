package main

import "fmt"

// Taken from Jaroslav's program
// https://github.com/prog-1/numeral-systems-ex03/blob/main/benchmarks/main.go
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

func LongMulWithLongAdd(a, b string) string {
	if a == "0" || b == "0" {
		return "0"
	}
	if len(a) > len(b) {
		a, b = b, a
	}
	toSum := make([][]byte, len(a))
	for ia := range a {
		var carry byte
		toSum[ia] = make([]byte, len(b)+1)
		for ib := range b {
			sum := (a[len(a)-1-ia]-'0')*(b[len(b)-1-ib]-'0') + carry
			toSum[ia][len(toSum[ia])-1-ib], carry = sum%10+'0', sum/10
		}
		if carry > 0 {
			toSum[ia][0] = carry + '0'
		} else {
			toSum[ia] = toSum[ia][1:]
		}
	}
	var result string
	for i := range toSum {
		for j := i; j > 0; j-- {
			toSum[i] = append(toSum[i], '0')
		}
		result = LongAdd(result, string(toSum[i]))
	}
	return result
}

// Reconstructed from program from website: https://www.geeksforgeeks.org/multiply-large-numbers-represented-as-strings/
func LongMulWithoutLongAdd(a, b string) string {
	if a == "0" || b == "0" {
		return "0"
	}
	ab := make([]byte, len(a)+len(b))
	for i, ia, ib := len(a)-1, 0, 0; i >= 0; i-- {
		var carry byte
		ib = 0
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
	var result string
	for i := len(ab) - 1; i >= 0; i-- {
		result += fmt.Sprint(ab[i])
	}
	return result
}

func main() {
	var a, b string
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	ab := LongMulWithoutLongAdd(a, b)
	fmt.Println("The multiplication of numbers: ", ab)
}
