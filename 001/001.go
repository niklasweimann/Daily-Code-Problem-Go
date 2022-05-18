package day001

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(getSolution(4))
	fmt.Println(getSolution(6))
	fmt.Println(getSolution(8))
	fmt.Println(getSolution(10))
}

func getSolution(input int) (int, int) {
	for i := input; i > 0; i-- {
		var n = input - i
		if big.NewInt(int64(i)).ProbablyPrime(0) && big.NewInt(int64(n)).ProbablyPrime(0) {
			return n, i
		}
	}
	return 0, 0
}
