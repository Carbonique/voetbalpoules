package voetbalpoules

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

//DeelnemerService handles communication related to Deelnemers
type DeelnemerService service

type Deelnemer struct {
	Naam   string
	ID     int
	Punten int
}

type Voorspelling struct {
	Datum                time.Time
	ThuisTeam            string
	UitTeam              string
	DoelpuntenThuis      int
	DoelpuntenUit        int
	Wvdw                 bool
	ThuisDoelpuntenMaker string
	UitDoelpuntenMaker   string
}

type voorspellingTabel struct {
	*colly.HTMLElement
}

// Get returns a voorspelling for a deelnemer for a wedstrijd
func (d *DeelnemerService) GetVoorspelling(id string, w Wedstrijd) ([]Voorspelling, error) {

	var voorspellingen []Voorspelling
	var voorspelling Voorspelling

	// First fetch the wedstrijdTabel
	t, err := d.getVoorspellingTabel(id, w)
	if err != nil {
		return []Voorspelling{}, err
	}
	// Now loop through the voorspellingTabel to get the voorspellingRijen
	vRijen := t.getVoorspellingRijen()
	for i, rij := range vRijen {

		datum, err := rij.datum(d.Time)
		if err != nil {
			log.Debug("Error on getting datum from rij")
			continue
		}
		if datum != w.Datum {
			log.Debug("datum is not wedstrijd datum. Continuing")
			continue
		}
		// Als wvdw dan voegen we een extra rij toe
		rijen := []voorspellingRij{rij}
		if rij.DOM.HasClass("wvdw") {
			if i+1 <= len(vRijen) {
				rijen = append(rijen, vRijen[i+1])
			}
		}

		voorspelling, err = NewVoorspelling(w.Competitie, w.Datum, rijen...)

		if err != nil {
			return []Voorspelling{}, err
		}

		voorspellingen = append(voorspellingen, voorspelling)

	}
	return voorspellingen, nil
}

//getVoorspellingTabel returns the voorspellingtabel for a user
func (d *DeelnemerService) getVoorspellingTabel(id string, w Wedstrijd) (voorspellingTabel, error) {
	var elem colly.HTMLElement
	d.OnHTML("table.wedstrijden", func(tabel *colly.HTMLElement) {
		// maak een wedstrijdTabel van tabel, om receiver methods toe te kunnen passen
		elem = *tabel
	})

	url := fmt.Sprintf("%sdeelnemer/%s/voorspellingen/%s", d.baseURL, id, w.Competitie)
	log.Infof("Visiting url: %s", url)
	d.Visit(url)
	return voorspellingTabel{&elem}, nil
}

type voorspellingRij struct {
	*colly.HTMLElement
}

func (v voorspellingTabel) getVoorspellingRijen() []voorspellingRij {
	rij := "tr:not(:first-child)"

	var vRijen []voorspellingRij

	v.ForEach(rij, func(_ int, r *colly.HTMLElement) {
		//Maak een nieuwe wedstrijdRij aan, zodat we daar receiver methods op toe kunnen passen
		vRij, err := newVoorspellingRij(r)
		if err != nil {
			return
		}
		vRijen = append(vRijen, vRij)
	})

	return vRijen
}

func newVoorspellingRij(e *colly.HTMLElement) (voorspellingRij, error) {
	if !isVoorspellingRij(e) {
		log.Errorf("Element %s is geen voorspellingrij", strings.Fields(e.Text))
		return voorspellingRij{}, fmt.Errorf("is geen voorspellingrij")
	}
	return voorspellingRij{e}, nil
}

//NewVoorspelling creates a voorspelling from a voorspellingrij
func NewVoorspelling(competitie string, vandaag time.Time, vRij ...voorspellingRij) (Voorspelling, error) {

	v := Voorspelling{}
	v.ThuisTeam = vRij[0].thuisTeam()
	v.UitTeam = vRij[0].uitTeam()

	log.Infof("Getting voorspelling %s - %s", v.ThuisTeam, v.UitTeam)

	datum, err := vRij[0].datum(vandaag)
	if err != nil {
		return Voorspelling{}, err
	}
	v.Datum = datum

	v.DoelpuntenThuis = vRij[0].doelpuntenThuis()
	v.DoelpuntenUit = vRij[0].doelpuntenUit()

	v.Wvdw = vRij[0].wvdw()
	if len(vRij) > 1 {
		v.ThuisDoelpuntenMaker = vRij[1].thuisDoelpuntenMaker()
		v.UitDoelpuntenMaker = vRij[1].uitDoelpuntenMaker()
	}

	return v, nil

}

func (r *voorspellingRij) thuisTeam() string {
	t := r.ChildText("td:nth-child(2) .vp-team")

	return t
}

func (r *voorspellingRij) uitTeam() string {
	t := r.ChildText("td:nth-child(3) .vp-team")
	return t
}

func (r *voorspellingRij) thuisDoelpuntenMaker() string {
	d := r.ChildText("td:nth-child(2)")

	return d
}

func (r *voorspellingRij) uitDoelpuntenMaker() string {
	d := r.ChildText("td:nth-child(3)")
	return d
}

func (r *voorspellingRij) doelpuntenThuis() int {
	u := r.ChildText("td:nth-child(4)")
	i, _ := strconv.Atoi(strings.Split(u, "-")[0])
	return i
}

func (r *voorspellingRij) doelpuntenUit() int {
	u := r.ChildText("td:nth-child(4)")
	i, _ := strconv.Atoi(strings.Split(u, "-")[1])
	return i
}

func (r *voorspellingRij) wvdw() bool {
	return r.DOM.HasClass("wvdw")
}

func isVoorspellingRij(e *colly.HTMLElement) (b bool) {
	//Dit is geen briljant criterium, een losse cel met alleen '.vp-team' zou nu ook als wedstrijdRij gezien worden
	return e.ChildText(".vp-team") != "" || e.DOM.HasClass("wvdw")

}

//datum extracts the date from a voorspellingRij
func (r *voorspellingRij) datum(vandaag time.Time) (t time.Time, err error) {

	cel, err := r.datumCel()

	if err != nil {
		return t, fmt.Errorf("datum: failed parsing %w", err)
	}
	dag := strings.Fields(cel)[0]
	tijd := strings.Fields(cel)[1]

	uur, err := strconv.Atoi(strings.Split(tijd, ":")[0])
	if err != nil {
		return t, err
	}

	minuten, err := strconv.Atoi(strings.Split(tijd, ":")[1])
	if err != nil {
		return t, err
	}

	t = time.Date(vandaag.Year(), vandaag.Month(), vandaag.Day(), uur, minuten, 0, 0, vandaag.Location())

	switch dag {

	case "Gisteren":
		t = t.AddDate(0, 0, -1)

	case "Vandaag":
		// do nothing

	case "Morgen":
		t = t.AddDate(0, 0, 1)

	}
	return t, err
}

//datumCel extracts the cel containing a date from a wedstrijdRij
func (r *voorspellingRij) datumCel() (s string, err error) {
	cel := r.ChildText("td.nowrap:first-child")

	if cel == "" {
		return "", fmt.Errorf("datum: failed parsing %w", err)
	}

	return cel, err

}
