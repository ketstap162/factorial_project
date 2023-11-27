package HttpUtils

import "math/big"

type FactirualRequestBody struct {
	A int `json:"a"`
	B int `json:"b"`
}

type FactorialResponseBody struct {
	AFactorial *big.Int `json:"a_factorial"`
	BFactorial *big.Int `json:"b_factorial"`
}
