package main

func percentageOfForestBurned(board [][]map[string]int) float64 { //to check how many trees are burned
	burned := 0
	total := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j]["is_tree"] == 2 {
				burned++
			}
			if board[i][j]["is_tree"] == 1 || board[i][j]["is_tree"] == 2 {
				total++
			}
		}
	}
	return float64(burned) / float64(total) * 100
}

func percentageOfTreesBeforeBurned(board [][]map[string]int) float64 { //to check how many trees were in forest before they were burned
	trees := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j]["is_tree"] == 1 || board[i][j]["is_tree"] == 2 {
				trees++
			}
		}
	}
	return float64(trees) / float64(len(board)*len(board[0])) * 100
}

func whichTreesNotBurned(board [][]map[string]int) []map[string]int { //to check how many trees are not burned but they were in fire
	var result []map[string]int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j]["is_tree"] == 1 && board[i][j]["burned_percentage"] > 0 && board[i][j]["age"] < 10 {
				result = append(result, board[i][j])
			}
		}
	}
	return result
}

func howManyTreesStayed(board [][]map[string]int) int { //to check how many trees stayed in forest
	trees := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j]["is_tree"] == 1 {
				trees++
			}
		}
	}
	return trees
}
