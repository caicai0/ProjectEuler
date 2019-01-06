package main

import (
	"fmt"
	"math/big"
)

func main() {
	res := new(big.Int).SetUint64(1)
	for i := 1; i <= 100; i++ {
		nub := new(big.Int).SetUint64(uint64(i))
		res.Mul(res,nub)
	}
	fmt.Println(res)
	sum := 0
	ten := new(big.Int).SetInt64(10)
	mod := new(big.Int).SetInt64(0)
	zero := new(big.Int).SetInt64(0)
	for ; zero.Cmp(res) < 0; {
		res.DivMod(res,ten,mod)
		sum = int(int64(sum) + mod.Int64())
		fmt.Println(res,mod,sum)
	}
	fmt.Println(sum)
}
