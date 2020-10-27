package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"time"
)

func main() {
	t0 := time.Now()

	for i := 1; i < 1143; i++ {
		c := colly.NewCollector()

		c.OnHTML("article", func(e *colly.HTMLElement) {
			//fmt.Println(e.Text)
		})

		c.Visit("https://baneks.ru/" + strconv.Itoa(i))

	}

	t1 := time.Now()
	fmt.Println(t1.Sub(t0))

}
