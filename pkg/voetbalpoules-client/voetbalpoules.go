package voetbalpoulesclient

import (
	"time"

	"github.com/gocolly/colly/v2"
)

//A client that manages communication with Voetbalpoules
type Client struct {
	client  *colly.Collector
	baseURL string
	time    time.Time

	Wedstrijden    *WedstrijdService
	Pool           *PoolService
	Voorspellingen *VoorspellingenService
}

type service struct {
	*Client
}

//Creates a new Client instance
func NewClient(url string) *Client {
	col := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	c := &Client{
		client:  col,
		baseURL: url + "/"}

	c.time = time.Now()

	c.Wedstrijden = &WedstrijdService{c}
	c.Pool = &PoolService{c}
	c.Voorspellingen = &VoorspellingenService{c}

	return c

}
