package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func readJSON() ([]SharkAttack, error) {
	jsonFile, err := os.Open("global-shark-attack.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var attacks []SharkAttack
	err = json.Unmarshal(byteValue, &attacks)
	if err != nil {
		fmt.Println(err)
	}
	return random10(attacks), err
}
