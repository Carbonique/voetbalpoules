package voetbalpoulesclient

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type Deelnemer struct {
	ID          int
	Naam        string
	Punten      int
	PuntenRonde int
}

type deelnemerRij struct {
	*colly.HTMLElement
}

//newDeelnemer creates a Deelnemer from a deelnemerRij
func newDeelnemer(d deelnemerRij) (Deelnemer, error) {

	punten, err := d.punten()
	if err != nil {
		return Deelnemer{}, err
	}

	puntenRonde, err := d.puntenRonde()
	if err != nil {
		return Deelnemer{}, err
	}

	id, err := d.id()
	if err != nil {
		return Deelnemer{}, err
	}

	naam := d.naam()

	deelnemer := Deelnemer{}
	deelnemer.Punten = punten
	deelnemer.PuntenRonde = puntenRonde
	deelnemer.ID = id
	deelnemer.Naam = naam

	return deelnemer, nil
}

func newDeelnemerRij(e *colly.HTMLElement) (deelnemerRij, error) {
	if !isDeelnemerRij(e) {
		log.Debugf("Element %s is geen deelnemerRij", strings.Fields(e.Text))
		return deelnemerRij{}, fmt.Errorf("is geen deelnemerRij")
	}
	return deelnemerRij{e}, nil
}

//isDeelnemerRij returns true if the colly.HTMLElement is a deelnemerrij
func isDeelnemerRij(e *colly.HTMLElement) (b bool) {
	//Dit is geen briljant criterium, een losse cel met alleen '.rank-deelnemer' zou nu ook als wedstrijdRij gezien worden
	return e.ChildText(".rank-deelnemer") != ""

}

//punten gets the total points for a deelnemer
func (d deelnemerRij) punten() (int, error) {
	puntenTekst := d.ChildText("td.punten")
	strafPunten := d.ChildText("td.punten div")
	puntenZonderPunt, err := strconv.Atoi(strings.ReplaceAll(puntenTekst, strafPunten, ""))

	if err != nil {
		return 0, err
	}

	return puntenZonderPunt, nil
}

//puntenRonde gets the points for a round for a deelnemer
func (d deelnemerRij) puntenRonde() (int, error) {
	puntenTekst := d.ChildText("td:last-child")

	re := regexp.MustCompile("[0-9]+")

	puntenRonde, err := strconv.Atoi(re.FindString(puntenTekst))

	if err != nil {
		return 0, err
	}

	return puntenRonde, nil
}

//naam gets the naam for a deelnemer
func (d deelnemerRij) naam() string {
	return strings.TrimSpace(d.ChildText("a"))
}

//id gets the id for a deelnemer
func (d deelnemerRij) id() (int, error) {
	split := strings.Split(d.ChildAttr("a", "href"), "/")
	id, err := strconv.Atoi(split[2])

	if err != nil {
		return 0, err
	}

	return id, nil
}
