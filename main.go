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
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visitinf")
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML("div", func(e *colly.HTMLElement) {
		fmt.Print("%v", e)
		currentPrice := e.ChildText("span.w_iUH7")

		fmt.Println(currentPrice)
	})

	c.Visit("https://www.walmart.com/search?q=banana")

}
