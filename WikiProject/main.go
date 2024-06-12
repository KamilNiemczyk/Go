package main

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	var clubs []Club
	c.OnHTML("table.wikitable:nth-of-type(4) > tbody", func(h *colly.HTMLElement) {
		counter := 0
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			if counter < 18 && counter > 0 {
				stringplace := el.ChildText("td:nth-child(4)")
				place, _ := strconv.Atoi(stringplace)
				club := Club{
					short: el.ChildText("td:nth-child(1)"),
					full:  el.ChildText("td:nth-child(2)"),
					city:  el.ChildText("td:nth-child(3)"),
					place: place,
				}
				clubs = append(clubs, club)
			}
			counter++
		})
	})

	c.Visit("https://pl.wikipedia.org/wiki/Premier_League")
	exportDataToCsv(clubs)
	fmt.Println(clubs)
	fmt.Print("end")

}
