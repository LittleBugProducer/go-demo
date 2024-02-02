package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

func btnIdTest() {
	c := colly.NewCollector()

	c.OnHTML("#su", func(e *colly.HTMLElement) {
		baiduBtn := e.Attr("value")
		fmt.Println(baiduBtn)
	})

	if err := c.Visit("http://www.baidu.com/"); err != nil {
		panic(err)
	}
}

type PageChapter struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Index int    `json:"index"`
}

func bqgTest() {
	collector := colly.NewCollector(
		func(collector *colly.Collector) {
			extensions.RandomUserAgent(collector)
		},
		func(collector *colly.Collector) {
			collector.OnRequest(func(reqeuset *colly.Request) {
				fmt.Println(reqeuset.URL, ", User-Agent:", reqeuset.Headers.Get("User-Agent"))
			})
		},
	)
	collector.SetRequestTimeout(time.Second * 60)

	var cha []PageChapter
	collector.OnHTML(".listmain dl", func(element *colly.HTMLElement) {
		element.DOM.Children().Each(func(i int, selection *goquery.Selection) {
			selection = selection.ChildrenFiltered("a")
			link, _ := selection.Attr("href")
			name := strings.TrimSpace(selection.Text())
			cha = append(cha, PageChapter{
				Index: i,
				Url:   element.Request.AbsoluteURL(link),
				Name:  name,
			})

		})
	})
	if err := collector.Visit("http://www.bqg5200.net/16/16705/"); err != nil {
		panic(err)
	}

	for _, ele := range cha {
		fmt.Println("name:", ele.Name, ", url:", ele.Url, ", index: ", ele.Index)
	}

}

type HotItem struct {
	Link  string
	Img   string
	Title string
	Desc  string
	Hot   int
}

func baiduHotTest() {
	collector := colly.NewCollector(
		func(collector *colly.Collector) {
			extensions.RandomUserAgent(collector)
		},
		func(c *colly.Collector) {
			c.OnRequest(func(request *colly.Request) {
				fmt.Println(request.URL, ", User-Agent:", request.Headers.Get("User-Agent"))
			})
		},
	)
	collector.SetRequestTimeout(time.Second * 60)

	data := []HotItem{}

	collector.OnHTML(".container-bg_lQ801", func(element *colly.HTMLElement) {
		element.ForEach(".category-wrap_iQLoo", func(i int, element *colly.HTMLElement) {
			aLink := element.DOM.ChildrenFiltered("a")
			jumpLink, _ := aLink.Attr("href")
			imgLink, _ := aLink.ChildrenFiltered("img").Attr("src")

			title := element.ChildText(".content_1YWBm .c-single-text-ellipsis")
			desc := element.ChildText(".content_1YWBm .large_nSuFU ")
			data = append(data, HotItem{
				Link:  jumpLink,
				Img:   imgLink,
				Title: title,
				Desc:  desc,
			})
		})
	})

	if err := collector.Visit("https://top.baidu.com/board?tab=realtime"); err != nil {
		panic(err)
	}

	for _, ele := range data {
		fmt.Println("title:", ele.Title, ", desc:", ele.Desc, ", img:", ele.Img, ", link:", ele.Link)
	}

}

func seperate() {
	fmt.Println("============================-------------------------============================-------------------------")
}

func main() {
	btnIdTest()
	seperate()

	bqgTest()
	seperate()

	baiduHotTest()
	seperate()

	fmt.Println("done")

}
