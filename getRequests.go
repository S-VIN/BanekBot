package main

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

func getAnecFromInternet(index int) (output string, err error) {
	c := colly.NewCollector()

	c.OnHTML("article", func(e *colly.HTMLElement) {
		output = e.Text
	})
	err = c.Visit("https://baneks.ru/" + strconv.Itoa(index))
	
	fmt.Println("Get Anec from Internet ", output)
	return output, err
}