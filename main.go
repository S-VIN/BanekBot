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
	for i := 1; i < 1141; i++ {
		c := colly.NewCollector()

		// After making a request get "url" from
		// the context of the request
		c.OnResponse(func(r *colly.Response) {
			fmt.Println("OK")
		})

		c.OnHTML("article", func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})

		c.Visit("https://baneks.ru/" + string(i))
		c.Wait()
	}
}
