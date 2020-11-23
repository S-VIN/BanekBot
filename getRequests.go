package main

import "github.com/gocolly/colly"
import "strconv"

func getAnecFromInternet(index int) string {
	var output string
	c := colly.NewCollector()

	c.OnHTML("article", func(e *colly.HTMLElement) {
		output = e.Text
	})

	c.Visit("https://baneks.ru/" + strconv.Itoa(index))

	return output
}