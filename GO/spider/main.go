package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

var visited = map[string]bool{}

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
// onclick="goTeacherHomepage('wangxuan')"
	c.Visit("http://faculty.hitsz.edu.cn/discipline-direction?id=2&browseName=%E6%95%99%E5%B8%88%E5%90%8D%E5%BD%95&browseEnName=TEACHERS")
}