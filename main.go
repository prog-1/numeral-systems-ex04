package main

import (
	sum "github.com/rchrdsasa207/numeral-systems-ex03"
)

func LongMul(a, b string) (res string) {
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
		res = sum.LongAdd(res, string(toSum[i]))
	}

	return
}

func main() {

}
