package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	for i := 1; ; i++ {
		sum = sum+i
		if numberOFDivisors(sum)>500{
			fmt.Println(i,sum)
			break
		}
	}
}

func numberOFDivisors(num int) int {
	sum := 0
	more :=0
	for i:=1;float64(i)<=math.Sqrt(float64(num));i++ {
		if num%i == 0 {
			sum = sum + 1
		}
		if float64(i) == math.Sqrt(float64(num)) {
			more = 1;
		}
	}
	sum = sum*2-more
	fmt.Println(sum)
	return sum
}
