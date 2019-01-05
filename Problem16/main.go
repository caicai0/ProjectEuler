package main

import "fmt"

func main() {
	var dgts [1000]int
	dgts[999] = 2
	for i:=2 ;i<=1000 ; i++ {
		quotient := 0
		remainder := 0
		for j := 0; j < len(dgts); j++ {
			remainder = dgts[999-j]*2+quotient
			quotient = remainder/10
			remainder = remainder%10
			dgts[999-j]=remainder
		}
	}

	sum := 0
	for i:=0;i< len(dgts);i++ {
		sum = sum + dgts[i]
	}
	fmt.Println(sum)
}
