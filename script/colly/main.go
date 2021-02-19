package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

const URL = "http://www.enjoybar.com"

func main(){
	fmt.Println("start")

	agent := colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36")
	depth := colly.MaxDepth(1)
	index := colly.NewCollector(agent, depth)

	//con := index.Clone()
	con := infoC(index.Clone())

	index = indexC(index, con)
	err := index.Visit(URL)
	if err != nil {
		panic(err)
	}

	fmt.Println("End")
}

// indexC 首页爬取
func indexC(index *colly.Collector, info *colly.Collector) *colly.Collector {
	//请求前调用
	index.OnRequest(func(r *colly.Request) {
		fmt.Println("index爬取：", r.URL)
	})

	//请求发生错误时调用
	index.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	//采集器，获取首页文章列表
	index.OnHTML("div[class='container-bd']", func(e *colly.HTMLElement) {
		csign := 0
		e.ForEach("div[class='wrap']", func(ci int, cv *colly.HTMLElement){
			cName := cv.ChildText("dt[class='title'] h3 span a")
			if cName != "" {
				csign++
				cv.ForEach("li", func(itemI int, itemV *colly.HTMLElement) {
					itemHref := itemV.ChildAttr("a[class='meiwen']", "href")
					itemName := itemV.ChildText("a[class='meiwen']")
					if itemName != "" && itemHref != ""{
						//fmt.Println(itemHref, itemName, csign)
						conUrl := fmt.Sprintf("%s%s", URL, itemHref)
						info.Request("GET", conUrl, nil, nil, nil)
					}
				})
			}
		})
	})

	return index
}

// infoC 内容爬取
func infoC(info *colly.Collector) *colly.Collector{
	//限速
	info.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob:   "*.enjoybar.com/*",
		Delay:        2 * time.Second,
		RandomDelay:  0,
		Parallelism:  1,
	})

	//(内容)请求前调用
	info.OnRequest(func(r *colly.Request) {
		//fmt.Println("con爬取：", r.URL)
	})

	//(内容)错误请求调用
	info.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	//(内容)
	info.OnHTML("div[class='container']", func(art *colly.HTMLElement){
		conDate  := art.ChildText("li[class='pubdate'] span")
		conClick := art.ChildText("li[class='click'] span")
		conArtic := art.ChildText("div[class='text'] p")
		conTitle := art.ChildText("div[class='article'] h1")

		var conImages [4]string
		var conImageInt int = 0

		art.ForEach("ul[class='picture-list'] li", func(conInt int, conImage *colly.HTMLElement){
			imageUrl := conImage.ChildAttr("a[class='meiwen'] img", "src")
			if imageUrl != "" && conImageInt <= 3 {
				conImages[conImageInt] = imageUrl
				conImageInt++
			}
		})

		fmt.Println(conTitle, conClick, conImages, conArtic, conDate)
	})

	return info
}
