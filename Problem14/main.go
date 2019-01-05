package main

import "fmt"

func main() {
	maxlength := 0
	maxi := 0
	for i := 13; i<1000000;i++ {
		len := lengthOfCollatzSequence(i)
		if len > maxlength {
			maxlength = len
			maxi = i
		}
	}
	fmt.Println(maxi,maxlength)
}

func lengthOfCollatzSequence(num int) int {
	var s = num
	length := 1
	for ; s != 1; {
		if s%2 == 0 {
			s = s/2
		}else {
			s = s*3+1
		}
		length = length + 1
	}
	return length
}
