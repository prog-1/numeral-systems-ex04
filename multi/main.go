package main

import "fmt"

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func LongMul(a, b string) string {
	if len(a) > len(b) {
		t := a
		a = b
		b = t
	}
	if len(a) == 0 && len(b) == 0 {
		return "0"
	}
	var car byte
	a = Reverse(a)
	b = Reverse(b)
	res := make([]byte, len(b)+1)
	for i := range a {
		sum := car + a[len(a)-1-i] - '0' + b[len(b)-1-i] - '0'
		res[i], car = sum%10+'0', sum/10
	}
	for i := len(a); i < len(b); i++ {
		sum := car + b[len(b)-1-i] - '0'
		res[i], car = sum%10+'0', sum/10
	}
	if car > 0 {
		res[len(res)-1] = car + '0'
	} else {
		res = res[:len(res)-1]
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func main() {
	fmt.Println(LongMul("177", "3"))
}
