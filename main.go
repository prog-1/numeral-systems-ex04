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
func LongMul(a, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	var carry byte
	var cnt int
	var res string
	for i := range b {
		var tmp []byte
		for c := 0; c < cnt; c++ {
			tmp = append(tmp, '0')
		}
		for i1 := range a {
			mul := carry + (a[len(a)-1-i1]-'0')*(b[len(b)-1-i]-'0')
			tmp = append(tmp, mul+'0')

		}
		for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
			tmp[i], tmp[j] = tmp[j], tmp[i]
		}

		res = LongAdd(string(res), string(tmp))

		cnt++
	}
	return string(res)

}
func main() {
	for _, t := range []struct {
		a, b string
	}{
		{"98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222"},
	} {
		fmt.Printf("%v * %v = %v\n", t.a, t.b, LongMul(t.a, t.b))
	}
}
