/*package main
import (
	"fmt"
	"github.com/gocolly/colly")
const CollyDocsLink = "https://baneks.ru/1141"

func main() {
	collector := colly.NewCollector()


	collector.Visit(func(e *colly.HTMLElement){
		println(e.Text)
	})
	collector.Visit(CollyDocsLink)
	fmt.Print(collector.OnResponse(func(r *colly.Response) {
		println(r)
	}))
}


package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

const CollyDocsLink = "https://baneks.ru/1141"

func main() {
	for i := 1; i < 1141; i++ {
		c := colly.NewCollector()
		c.OnResponse(func(r *colly.Response) {
			fmt.Println("Visited", r.Request.URL)
			fmt.Println(r.Request.Headers)
		})

		c.Visit("https://baneks.ru/" + string(i))
	}


	for i := 1000; i < 1141; i++ {
		c := colly.NewCollector()
		c.OnResponse(func(r *colly.Response) {
			fmt.Println("Visited", r.Request.URL)
			fmt.Println(r.Request.Headers)
		})

		c.Visit("https://baneks.ru/" + string(i))
	}

}*/

package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// Before making a request put the URL with
	// the key of "url" into the context of the request
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	// After making a request get "url" from
	// the context of the request
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Ctx.Get("url"))
	})

	// Start scraping on https://en.wikipedia.org
	c.Visit("https://en.wikipedia.org/")
}
