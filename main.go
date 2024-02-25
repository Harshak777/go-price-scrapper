package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type item struct {
	name, price string
}

func main() {
	c := colly.NewCollector()
	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15"

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("visitinf")
		r.Headers.Set("User-Agent", userAgent)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.Headers.Get("User-Agent"))
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML("div", func(e *colly.HTMLElement) {
		// fmt.Print("%v", e)
		// // currentPrice := e.ChildText("span.w_iUH7")

		fmt.Println(e.Text)
	})

	c.OnHTML("a[automation-id^=\"productDescriptionLink_\"]", func(e *colly.HTMLElement) {
		// fmt.Print("%v", e)
		// currentPrice := e.ChildText("span.w_iUH7")

		fmt.Println(e.Text)
	})

	c.Visit("https://www.costco.com/CatalogSearch?dept=All&keyword=rice")

}
