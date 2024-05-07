package main

func fibonacci(n int, result map[int]int) int {
	if n <= 1 {
		result[n]++
		return n
	} else {
		result[n]++
	}
	return fibonacci(n-1, result) + fibonacci(n-2, result)
}

func weakNumberList(n int) map[int]int {
	result := make(map[int]int)
	fibonacci(n, result)
	return result
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func weakNumberFind(n int, weakNumberList map[int]int) int {
	key := 0
	diff := 0
	for k, v := range weakNumberList {
		if absolute(n-v) < diff || diff == 0 {
			key = k
			diff = absolute(n - v)
		}
	}
	return key
}
