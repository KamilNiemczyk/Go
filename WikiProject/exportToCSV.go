package main

import (
	"os"
	"strconv"
)

func exportDataToCsv(clubs []Club) {
	file, err := os.Create("clubs.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Short,Full,City,Place\n")

	for _, club := range clubs {
		file.WriteString(club.short + "," + club.full + "," + club.city + "," + strconv.Itoa(club.place) + "\n")
	}
}
