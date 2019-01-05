package main

import (
	"fmt"
	"strings"
)

func main() {
	strall := ""
	for i := 1; i <= 1000; i++ {
		str := numberToString(i)
		strall = strall+str
		fmt.Println(str)
	}
	res := strings.Replace(strall," ","",-1)
	fmt.Println(len(res))
}

func numberToString(num int) string {
	if num == 1000 {
		return "one thousand"
	}
	switch num {
		case 0:
			return ""
		case 1:
			return "one"
		case 2:
			return "two"
		case 3:
			return "three"
		case 4:
			return "four"
		case 5:
			return "five"
		case 6:
			return "six"
		case 7:
			return "seven"
		case 8:
			return "eight"
		case 9:
			return "nine"
		case 10:
			return "ten"
		case 11:
			return "eleven"
		case 12:
			return "twelve"
		case 13:
			return "thirteen"
		case 14:
			return "fourteen"
		case 15:
			return "fifteen"
		case 16:
			return "sixteen"
		case 17:
			return "seventeen"
		case 18:
			return "eighteen"
		case 19:
			return "nineteen"
		case 20:
			return "twenty"
		case 30:
			return "thirty"
		case 40:
			return "forty"
		case 50:
			return "fifty"
		case 60:
			return "sixty"
		case 70:
			return "seventy"
		case 80:
			return "eighty"
		case 90:
			return "ninety"
	}
	if num < 100 {
		digit2 := num/10*10
		digit1 := num%10
		str1 := numberToString(digit2)
		str2 := numberToString(digit1)
		return str1+" "+str2
	}else {
		digit3 := num/100
		other := num%100
		str1 := numberToString(digit3)
		str2 := numberToString(other)
		if len(str2)==0 {
			return str1+" hundred"
		}else {
			return str1+" hundred and "+str2
		}
	}
}
