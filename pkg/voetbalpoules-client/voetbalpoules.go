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

//A struct containing a Wedstrijd and its DeelnemerVoorspellingen
type VoorspeldeWedstrijd struct {
	Wedstrijd               Wedstrijd
	DeelnemerVoorspellingen map[Deelnemer]Voorspelling
}

//NewClientt returns a new Voetbalpoules Client instance
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

//GetStand returns a sorted ranking for a pool
func (c Client) GetStand(poolID int, competitie string) []Deelnemer {
	return c.pool.getStand(poolID, competitie)
}

//GetPoolVoorspelling returns a wedstrijd and all deelnemer predictions for a pool for a competitie within a specified timerange
func (c Client) GetPoolVoorspelling(t1 time.Time, t2 time.Time, poolid int, competitie string) ([]VoorspeldeWedstrijd, error) {

	var vw []VoorspeldeWedstrijd
	//First check whether any wedstrijden are played within the time range.
	w, err := c.wedstrijden.get(competitie, t1, t2)
	if err != nil {
		return []VoorspeldeWedstrijd{}, err
	}
	if len(w) == 0 {
		return []VoorspeldeWedstrijd{}, nil
	}

	//Now get the deelnemers for the pool
	d := c.pool.getDeelnemers(poolid, competitie)

	//Now iterate through the slice of wedstrijden and create a map
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
