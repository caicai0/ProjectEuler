package main

import "fmt"

func main() {
	sum := 0
	for i := 2; i<=10000 ;i++ {
		j := sumOfProperDivisors(i)
		if j != i && j <= 10000 {
			if sumOfProperDivisors(j) == i {
				sum = sum + j +i
				fmt.Println(sum,j,i)
			}
		}
	}
	fmt.Println(sum/2)
}

func sumOfProperDivisors(x int) int {
	sum := 0
	for i := 1; i < x; i++ {
		if x%i == 0 {
			sum = sum + i
		}
	}
	return sum
}