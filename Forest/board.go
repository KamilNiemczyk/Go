package main

import (
	"math/rand"
)

func createBoard(x int, y int) [][]map[string]int {
	board := make([][]map[string]int, x)
	for i := 0; i < x; i++ {
		board[i] = make([]map[string]int, y)
		for j := 0; j < y; j++ {
			board[i][j] = make(map[string]int)
			board[i][j]["is_tree"] = 0
		}
	}
	return board
}

func generateTree(board [][]map[string]int, x int, y int) {
	random := rand.Intn(10) + 1
	board[x][y]["is_tree"] = 1               //0 means grass, 1 means tree, 2 means burned tree
	board[x][y]["age"] = random              //age of tree
	board[x][y]["burned_percentage"] = 0     //burned percentage of tree
	board[x][y]["how_many_times_burnes"] = 0 //how many times tree burned (max 3)
}

func generateForest(board [][]map[string]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			random := rand.Intn(2)
			if random == 1 {
				generateTree(board, i, j)
			}
		}
	}
}

func visualizeBoard(board [][]map[string]int) string {
	var result string
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j]["is_tree"] == 1 && board[i][j]["burned_percentage"] == 0 {
				result += "🌲"
			} else if board[i][j]["is_tree"] == 1 && board[i][j]["burned_percentage"] > 0 {
				result += "🪵"
			} else if board[i][j]["is_tree"] == 0 {
				result += "🟩"
			} else {
				result += "🔥"
			}
		}
		result += "\n"
	}
	return result
}

func generateForestWithProbability(board [][]map[string]int, x int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			random := rand.Intn(100)
			if random < x {
				generateTree(board, i, j)
			}
		}
	}
}
