package voetbalpoulesclient

import (
	"fmt"
	"sort"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

//poolService handles communication related to Pools
type poolService service

type Pool struct {
	Deelnemers []Deelnemer
}

//GetDeelnemers fetches the deelnemers for a pool
func (p *poolService) getDeelnemers(id int, competitie string) []Deelnemer {
	deelnemers := []Deelnemer{}

	p.client.OnHTML("table.stand", func(stand *colly.HTMLElement) {
		stand.ForEach("tr:not(:first-child)", func(_ int, rij *colly.HTMLElement) {
			dRij, err := newDeelnemerRij(rij)
			if err != nil {
				return
			}

			d, err := newDeelnemer(dRij)
			if err != nil {
				return
			}

			deelnemers = append(deelnemers, d)
		})
	})
	url := fmt.Sprintf("%s/poule/%d/stand/%s", p.baseURL, id, competitie)
	log.Infof("Visiting url: %s", url)

	p.client.Visit(url)
	return deelnemers
}

//GetStand returns the stand for a pool
func (p *poolService) getStand(poolID int, competitie string) []Deelnemer {
	deelnemers := p.getDeelnemers(poolID, competitie)
	// In principe returnt GetDeelnemers een gesorteerde stand op basis van de html tabel,
	// maar voor de zekerheid sorteren we alsnog
	return p.sorteerStand(deelnemers)
}

//Sort a slice of deelnemers by their points
func (p *poolService) sorteerStand(d []Deelnemer) []Deelnemer {
	sort.Slice(d, func(i, j int) bool {
		return d[i].Punten > d[j].Punten
	})
	return d
}
