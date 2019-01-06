package main

import (
	"fmt"
)

func main() {
	//假设月日  都是从0开始数
	y := 1901
	m := 0
	d := 6
	sum := 0
	for ; y<=2000 && m<=11 && d<=30;  {
		fmt.Println(y,m,d)
		addSevenDay(&y,&m,&d)
		if d == 1 {
			sum = sum + 1
		}
	}
	fmt.Println(sum)
}

func addSevenDay(y *int,m *int,d *int)  {
	dayInManth := 30
	if *m == 1-1 || *m == 3-1 || *m == 5-1 || *m == 7-1 || *m == 8-1 || *m == 10-1 || *m == 12-1 {
		dayInManth = 31
	}else if *m==2-1 {
		if leapYear(*y) {
			dayInManth = 29
		}else {
			dayInManth = 28
		}
	}
	newManth := (*d+7)/dayInManth
	newYear := (*m+newManth)/12
	*d = (*d+7)%dayInManth
	*m = (*m+newManth)%12
	*y = *y+newYear
}

func leapYear(y int) bool {
	return ((y%4 == 0 && y%100 != 0) || y%400 == 0)
}
