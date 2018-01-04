package main

import (
	// "github.com/henrylee2cn/pholcus/exec"
	_ "github.com/henrylee2cn/pholcus_lib" // 此为公开维护的spider规则库
	// _ "github.com/henrylee2cn/pholcus_lib_pte" // 同样你也可以自由添加自己的规则库
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func main() {
	// 设置运行时默认操作界面，并开始运行
	// 运行软件前，可设置 -a_ui 参数为"web"、"gui"或"cmd"，指定本次运行的操作界面
	// 其中"gui"仅支持Windows系统
	// exec.DefaultRun("web")
	{
		doc, err := goquery.NewDocument("http://metalsucks.net")
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			band := s.Find("a").Text()
			title := s.Find("i").Text()
			fmt.Printf("Review %d: %s - %s\n", i, band, title)
		})
	}
	{
		doc, err := goquery.NewDocument("https://www.etherchain.org/")
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find(".card-group .card-title").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			band := s.Find("span").Text()

			fmt.Printf("xx %d: %s\n", i, band)
		})
	}
}
