package main

import (
	"fmt"
	"strings"
)

func inputNick() string {
	var name string
	var lastname string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	fmt.Println("Enter your last name:")
	fmt.Scanln(&lastname)
	nameToLower := strings.ToLower(name)[:3]
	lastnameLower := strings.ToLower(lastname)[:3]
	nick := nameToLower + lastnameLower
	return nick
}
