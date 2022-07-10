package main

import (
	"github.com/gocolly/colly/v2"
)

type Scraper struct {
	*colly.Collector
	url string
}

//Creates a new Scraper instance
func NewScraper(url string) *Scraper {
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	s := Scraper{
		c,
		url,
	}

	return &s

}
