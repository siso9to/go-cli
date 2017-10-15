package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func main() {
	const defaultKeyword = "Ruby"

	var keyword string
	flag.StringVar(&keyword, "keyword", defaultKeyword, "keyword to use")
	flag.StringVar(&keyword, "k", defaultKeyword, "keyword to use (short)")

	flag.Parse()

	fmt.Println("keyword: ", keyword)

	r := regexp.MustCompile(keyword)

	doc, err := goquery.NewDocument("https://www.oreilly.co.jp/ebook/")
	if err != nil {
		fmt.Print("failed")
	}

	doc.Find("table#bookTable > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		isbn := s.Find("td.isbn").Text()
		title := s.Find("td.title > a").Text()
		price := s.Find("td.price").Text()
		published := s.Find("td:nth-child(4)").Text()

		if r.MatchString(title) {
			fmt.Printf("%s, %s, ¥%s, %s\n", isbn, title, price, published)
		}
	})
}
