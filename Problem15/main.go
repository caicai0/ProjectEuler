package main

import "fmt"

func main() {
	var count = 0
	dfs(0,0,20,20,&count)
	fmt.Println(count)
}

func dfs(x int, y int, m int, n int, count *int){
	if x == m && y == n {
		*count = *count+1
	}else {
		if x < m {
			dfs(x+1,y,m,n,count)
		}
		if y < n {
			dfs(x,y+1,m,n,count)
		}
	}
}

