package voetbalpoules

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Deelnemer struct {
	ID     int
	Naam   string
	Punten int
}

type DeelnemerRij struct {
	*colly.HTMLElement
}

func NewDeelnemer(d DeelnemerRij) (Deelnemer, error) {
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

func (d DeelnemerRij) punten() (int, error) {
	puntenTekst := d.ChildText("td.punten")
	strafPunten := d.ChildText("td.punten div")
	positieZonderPunt, err := strconv.Atoi(strings.ReplaceAll(puntenTekst, strafPunten, ""))

	if err != nil {
		return 0, err
	}

	return positieZonderPunt, nil
}

func (d DeelnemerRij) teamNaam() string {
	return strings.TrimSpace(d.ChildText("a"))
}

func (d DeelnemerRij) id() (int, error) {
	split := strings.Split(d.ChildAttr("a.links", "href"), "/")
	id, err := strconv.Atoi(split[2])

	if err != nil {
		return 0, err
	}

	return id, nil
}
