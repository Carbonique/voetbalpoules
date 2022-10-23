package voetbalpoules

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type Deelnemer struct {
	ID     int
	Naam   string
	Punten int
}

type deelnemerRij struct {
	*colly.HTMLElement
}

func NewDeelnemer(d deelnemerRij) (Deelnemer, error) {

	punten, err := d.punten()
	if err != nil {
		return Deelnemer{}, err
	}

	id, err := d.id()
	if err != nil {
		return Deelnemer{}, err
	}

	teamNaam := d.teamNaam()

	deelnemer := Deelnemer{}
	deelnemer.Punten = punten
	deelnemer.ID = id
	deelnemer.Naam = teamNaam

	return deelnemer, nil
}

func newDeelnemerRij(e *colly.HTMLElement) (deelnemerRij, error) {
	if !isDeelnemerRij(e) {
		log.Debugf("Element %s is geen deelnemerRij", strings.Fields(e.Text))
		return deelnemerRij{}, fmt.Errorf("is geen deelnemerRij")
	}
	return deelnemerRij{e}, nil
}
func isDeelnemerRij(e *colly.HTMLElement) (b bool) {
	//Dit is geen briljant criterium, een losse cel met alleen '.rank-deelnemer' zou nu ook als wedstrijdRij gezien worden
	return e.ChildText(".rank-deelnemer") != ""

}

func (d deelnemerRij) punten() (int, error) {
	puntenTekst := d.ChildText("td.punten")
	strafPunten := d.ChildText("td.punten div")
	positieZonderPunt, err := strconv.Atoi(strings.ReplaceAll(puntenTekst, strafPunten, ""))

	if err != nil {
		return 0, err
	}

	return positieZonderPunt, nil
}

func (d deelnemerRij) teamNaam() string {
	return strings.TrimSpace(d.ChildText("a"))
}

func (d deelnemerRij) id() (int, error) {
	split := strings.Split(d.ChildAttr("a", "href"), "/")
	id, err := strconv.Atoi(split[2])

	if err != nil {
		return 0, err
	}

	return id, nil
}
