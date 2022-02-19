package main

import "fmt"

func addZeros(a string, i int) (res string) {
	for ; i > 0; i-- {
		a = "0" + a
	}
	return a
}

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
			toSum[i][len(toSum[i])-1-i1], carry = sum%10, sum/10
		}
		if carry > 0 {
			toSum[i][0] = carry
		} else {
			toSum[i] = toSum[i][1:]
		}
	}
	fmt.Println(toSum)
	return
}


func main() {
	fmt.Println("Program implements arithmetics multiplication - accepts two decimal numbers encoded as strings.")
	var a string
	fmt.Println("Enter a:")
	fmt.Scan(&a)
	var b string
	fmt.Println("Enter b:")
	fmt.Scan(&b)
	fmt.Println(LongMul(a, b))
}
