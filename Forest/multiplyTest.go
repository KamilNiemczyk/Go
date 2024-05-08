package main

func multiplyBurnSimulationIncreasingProbability(n int) []map[int]float64 { //how many trees are burned with increasing forest density
	result := make([]map[int]float64, 0)
	for i := 0; i < 100; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			board := createBoard(20, 20)
			generateForestWithProbability(board, i)
			randomStrike(board)
			procent := percentageOfForestBurned(board)
			sum += procent
		}
		result = append(result, map[int]float64{i: float64(sum) / float64(n)})
	}
	return result
}

func howManyTreesStayedSimulation(n int) []map[int]float64 { //how many trees stayed in forest with increasing forest density
	result := make([]map[int]float64, 0)
	for i := 0; i < 100; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			board := createBoard(20, 20)
			generateForestWithProbability(board, i)
			randomStrike(board)
			procent := howManyTreesStayed(board)
			sum += procent
		}
		result = append(result, map[int]float64{i: float64(sum) / float64(n)})
	}
	return result
}
