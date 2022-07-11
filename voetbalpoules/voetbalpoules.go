package voetbalpoules

import "github.com/gocolly/colly/v2"

//A client that manages communication with Voetbalpoules
type Client struct {
	scraper *colly.Collector
	BaseURL string

	Competitie *CompetitieService
	Wedstrijd  *WedstrijdService
	Pool       *PoolService
	Deelnemer  *DeelnemerService
}

type service struct {
	client *Client
}

//Creates a new Scraper instance
func NewClient(url string) *Client {
	col := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	c := &Client{scraper: col, BaseURL: url}

	c.Competitie = &CompetitieService{c}
	c.Wedstrijd = &WedstrijdService{c}
	c.Pool = &PoolService{c}
	c.Deelnemer = &DeelnemerService{c}

	return c

}
