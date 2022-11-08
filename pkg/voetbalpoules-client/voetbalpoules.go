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

	wedstrijden    *wedstrijdService
	pool           *poolService
	voorspellingen *voorspellingenService
}

type service struct {
	*Client
}

type VoorspeldeWedstrijd struct {
	Wedstrijd               Wedstrijd
	DeelnemerVoorspellingen map[Deelnemer]Voorspelling
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

	c.wedstrijden = &wedstrijdService{c}
	c.pool = &poolService{c}
	c.voorspellingen = &voorspellingenService{c}

	return c

}

func (c Client) GetStand(poolid int, competitie string) []Deelnemer {
	return c.pool.getStand(poolid, competitie)
}

func (c Client) GetPoolVoorspelling(t1 time.Time, t2 time.Time, poolid int, competitie string) ([]VoorspeldeWedstrijd, error) {

	var vw []VoorspeldeWedstrijd

	w, err := c.wedstrijden.get(competitie, t1, t2)
	if err != nil {
		return []VoorspeldeWedstrijd{}, err
	}

	d := c.pool.getDeelnemers(poolid, competitie)

	for _, w2 := range w {

		m := make(map[Deelnemer]Voorspelling)
		for _, d2 := range d {
			v, _ := c.voorspellingen.get(d2, w2)

			m[d2] = v
		}
		tempVW := VoorspeldeWedstrijd{
			Wedstrijd:               w2,
			DeelnemerVoorspellingen: m,
		}
		vw = append(vw, tempVW)
	}

	return vw, nil
}
