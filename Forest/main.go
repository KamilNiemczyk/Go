package main

import "fmt"

func main() {
	board := createBoard(20, 20)         //basic boar 20x20
	generateForest(board)                //generate forest
	boardbefore := visualizeBoard(board) //visualize board
	fmt.Println(boardbefore)             //print board
	randomStrike(board)                  //random strike of lightning
	newboard := visualizeBoard(board)    //visualize board
	fmt.Println(newboard)                //print board

	// procent := percentageOfForestBurned(board) //to check how many trees are burned
	// fmt.Println("Procent of burned trees: ", procent)

	// procent2 := percentageOfTreesBeforeBurned(board) //to check how many trees were in forest before they were burned
	// fmt.Println("Procent of trees before burned: ", procent2)

	// notBurned := whichTreesNotBurned(board)
	// fmt.Println("Trees that are not burned: ", notBurned) //to check how many trees are not burned but they were in fire

	// result := multiplyBurnSimulationIncreasingProbability(10000) //to check how many trees are burned with increasing forest density
	// fmt.Println(result)
	// generateChart1(result)

	// result2 := howManyTreesStayedSimulation(10000) //to check how many trees stayed in forest
	// fmt.Println(result2)
	// generateChart2(result2)

}
