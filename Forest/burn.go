package main

import (
	"fmt"
	"math/rand"
)

func burnTree(board [][]map[string]int, x int, y int) { //burn tree
	age := board[x][y]["age"]
	if board[x][y]["burned_percentage"] < 100 {
		if age <= 2 {
			board[x][y]["burned_percentage"] = 100
			board[x][y]["is_tree"] = 2
		} else if age == 10 {
			board[x][y]["burned_percentage"] += 25
		} else {
			board[x][y]["burned_percentage"] += (20/age)*25 - age
			if board[x][y]["burned_percentage"] >= 100 {
				board[x][y]["is_tree"] = 2
			}
		}
		board[x][y]["how_many_times_burnes"]++
		expandFire(board, x, y)
	}
}

func expandFire(board [][]map[string]int, x int, y int) { //expand fire
	if x-1 >= 0 {
		if board[x-1][y]["is_tree"] == 1 && board[x-1][y]["how_many_times_burnes"] < 3 {
			burnTree(board, x-1, y)
		}
	}
	if x+1 < len(board) {
		if board[x+1][y]["is_tree"] == 1 && board[x+1][y]["how_many_times_burnes"] < 3 {
			burnTree(board, x+1, y)
		}
	}
	if y-1 >= 0 {
		if board[x][y-1]["is_tree"] == 1 && board[x][y-1]["how_many_times_burnes"] < 3 {
			burnTree(board, x, y-1)
		}
	}
	if y+1 < len(board[x]) {
		if board[x][y+1]["is_tree"] == 1 && board[x][y+1]["how_many_times_burnes"] < 3 {
			burnTree(board, x, y+1)
		}
	}
	if x-1 >= 0 && y-1 >= 0 {
		if board[x-1][y-1]["is_tree"] == 1 && board[x-1][y-1]["how_many_times_burnes"] < 3 {
			burnTree(board, x-1, y-1)
		}
	}
	if x-1 >= 0 && y+1 < len(board[x]) {
		if board[x-1][y+1]["is_tree"] == 1 && board[x-1][y+1]["how_many_times_burnes"] < 3 {
			burnTree(board, x-1, y+1)
		}
	}
	if x+1 < len(board) && y-1 >= 0 {
		if board[x+1][y-1]["is_tree"] == 1 && board[x+1][y-1]["how_many_times_burnes"] < 3 {
			burnTree(board, x+1, y-1)
		}
	}
	if x+1 < len(board) && y+1 < len(board[x]) {
		if board[x+1][y+1]["is_tree"] == 1 && board[x+1][y+1]["how_many_times_burnes"] < 3 {
			burnTree(board, x+1, y+1)
		}
	}
}

func randomStrike(board [][]map[string]int) { //random strike of lightning
	x := rand.Intn(len(board))
	y := rand.Intn(len(board[0]))
	if board[x][y]["is_tree"] == 1 {
		burnTree(board, x, y)
		fmt.Println("Fire started at", x, y)
	} else {
		fmt.Println("No tree found at", x, y)
	}
}
