package Factorial

import (
	"math/big"
	"sync"
)

func factorial(n int) *big.Int {
	fact := big.NewInt(1)
	for i := 2; i <= n; i++ {
		fact.Mul(fact, big.NewInt(int64(i)))
	}
	return fact
}

func CalculateFactorials(a, b int) map[string]*big.Int {
	var wg sync.WaitGroup
	result := make(map[string]*big.Int)

	var a_factorial, b_factorial *big.Int

	wg.Add(2)
	go func() {
		defer wg.Done()
		a_factorial = factorial(a)
	}()
	go func() {
		defer wg.Done()
		b_factorial = factorial(b)
	}()

	wg.Wait()

	result["a_factorial"] = a_factorial
	result["b_factorial"] = b_factorial

	return result
}
