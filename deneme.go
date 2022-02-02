package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://www.imdb.com/chart/top")
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		fmt.Println("Hata", res.StatusCode)
		return
	}

	names := []string{}
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	doc.Find(".titleColumn").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		names = append(names, title)
	})

	for i := 0; i < len(names)-1; i++ {
		fmt.Printf("Name:%s ", names[i])
		fmt.Println()
	}

	*******************************

	

}
