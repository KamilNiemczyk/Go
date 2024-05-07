package main

import "strconv"

func generateNickASCII(name string) []string {
	var ascii []string
	for i := 0; i < len(name); i++ {
		ascii = append(ascii, strconv.Itoa(int(name[i])))
	}
	return ascii
}
