package main

import (
	"math/big"
	"strings"
)

func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func factorialString(n int64) string {
	return factorial(n).String()
}

func contains(s string, substr string) bool {
	return strings.Contains(s, substr)
}

func ifContainsAllNumbers(asciiList []string, factorialString string) bool {
	for _, ascii := range asciiList {
		if !contains(factorialString, ascii) {
			return false
		}
	}
	return true
}

func strongNumber(asciiList []string) int64 {
	var i int64 = 1
	for {
		factorialString := factorialString(i)
		if ifContainsAllNumbers(asciiList, factorialString) {
			return i
		}
		i++
	}
}
