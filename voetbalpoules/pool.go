package voetbalpoules

import (
	"fmt"
	"sort"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

//PoolService handles communication related to Pools
type PoolService service

type Pool struct {
	Deelnemers []Deelnemer
}

//GetDeelnemers fetches the deelnemers for a pool sorted by their points
func (p *PoolService) GetDeelnemers(id int, competitie string) []Deelnemer {
	deelnemers := []Deelnemer{}

	p.OnHTML("table.stand", func(stand *colly.HTMLElement) {
		stand.ForEach("tr:not(:first-child)", func(_ int, rij *colly.HTMLElement) {
			dRij, err := newDeelnemerRij(rij)
			if err != nil {
				return
			}

			d, err := NewDeelnemer(dRij)
			if err != nil {
				return
			}

			deelnemers = append(deelnemers, d)
		})
	})
	url := fmt.Sprintf("%s/poule/%d/stand/%s", p.baseURL, id, competitie)
	log.Infof("Visiting url: %s", url)

	p.Visit(url)
	p.SorteerStand(deelnemers)
	return deelnemers
}

//Sort a slice of deelnemers by their points
func (p *PoolService) SorteerStand(d []Deelnemer) []Deelnemer {
	sort.Slice(d, func(i, j int) bool {
		return d[i].Punten > d[j].Punten
	})
	return d
}
