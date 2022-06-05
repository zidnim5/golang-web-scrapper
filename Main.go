package main

import (
	"goscrapper/pkg/scrapper"
)

func main() {
	var webScrapper scrapper.Scrapper
	webScrapper = scrapper.NewGoquery()

	webScrapper.Fetch()
}
