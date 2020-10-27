package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
)

func main() {

	for i := 1; i < 1143; i++ {
		c := colly.NewCollector(
			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
			colly.AllowedDomains("baneks.ru"),
		)

		// On every a element which has href attribute call callback
		c.OnHTML("article", func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})

		// Before making a request print "Visiting ..."
		c.OnRequest(func(r *colly.Request) {

		})

		c.Visit("https://baneks.ru/" + strconv.Itoa(i))
	}
}
