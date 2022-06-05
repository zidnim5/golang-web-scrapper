package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

const (
	exportUrl string = "export/goquery.csv"
)

type GoqueryPkg struct{}

func NewGoquery() Scrapper {
	return GoqueryPkg{}
}

func (g GoqueryPkg) Fetch() {
	var export [][]string
	export = [][]string{{"title", "galery"}}

	res, err := http.Get("https://otodriver.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".section .medium-post").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		imgTag := s.Find("img")
		imgData, _ := imgTag.Attr("data-src")

		tmp := []string{
			title, imgData,
		}

		export = append(export, tmp)
	})

	csvFile, err := os.Create(exportUrl)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range export {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()

	fmt.Println("Exported on : ", exportUrl)
}
