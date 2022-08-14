package voetbalpoules

import (
	"time"

	"github.com/gocolly/colly/v2"
)

//A client that manages communication with Voetbalpoules
type Client struct {
	*colly.Collector
	baseURL string
	Time    time.Time

	Wedstrijden *WedstrijdService
	Pool        *PoolService
	Deelnemer   *DeelnemerService
}

type service struct {
	*Client
}

//Creates a new Scraper instance
func NewClient(url string) *Client {
	col := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	c := &Client{
		Collector: col,
		baseURL:   url + "/"}

	c.Time = time.Now()

	c.Wedstrijden = &WedstrijdService{c}
	c.Pool = &PoolService{c}
	c.Deelnemer = &DeelnemerService{c}

	return c

}
