package voetbalpoules

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

//PoolService handles communication related to Pools
type PoolService service

type Pool struct {
	Deelnemers []Deelnemer
}

func (p *PoolService) GetDeelnemers(id int, competitie string) []Deelnemer {
	deelnemers := []Deelnemer{}

	p.OnHTML("table.stand", func(stand *colly.HTMLElement) {
		stand.ForEach("tr:not(:first-child)", func(_ int, rij *colly.HTMLElement) {
			d := Deelnemer{}
			split := strings.Split(rij.ChildAttr("a.links", "href"), "/")
			d.ID, _ = strconv.Atoi(split[2])
			fmt.Println(d.ID)
		})
	})
	url := fmt.Sprintf("%s/poule/%d/stand/%s", p.baseURL, id, competitie)
	log.Infof("Visiting url: %s", url)

	p.Visit(url)
	return deelnemers
}
