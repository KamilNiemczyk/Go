package main

import (
	"fmt"
)

func main() {
	name := inputNick()
	ascii := generateNickASCII(name)
	strongNumber := int(strongNumber(ascii))
	weakNumberList := weakNumberList(30)
	weakNumber := weakNumberFind(strongNumber, weakNumberList)
	fmt.Println("Nick:", name)
	fmt.Println("ASCII:", ascii)
	fmt.Println("Strong number: ", strongNumber)
	fmt.Println("Fibonnaci calls counter: ", weakNumberList)
	fmt.Println("Weak number: ", weakNumber)
}
