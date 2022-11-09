package voetbalpoulesclient

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

//voorspellingenService handles communication related to Voorspellingen
type voorspellingenService service

type Voorspelling struct {
	Datum                time.Time
	ThuisTeam            string
	UitTeam              string
	DoelpuntenThuis      *int
	DoelpuntenUit        *int
	Wvdw                 bool
	ThuisDoelpuntenMaker *string
	UitDoelpuntenMaker   *string
}

type voorspellingTabel struct {
	*colly.HTMLElement
}

type voorspellingRij struct {
	*colly.HTMLElement
}

// Get returns a voorspelling for a deelnemer for a wedstrijd
func (v *voorspellingenService) get(d Deelnemer, w Wedstrijd) (Voorspelling, error) {

	log.Infof("Trying to get voorspelling for %s - %s", w.ThuisTeam, w.UitTeam)

	var voorspelling Voorspelling

	// First fetch the wedstrijdTabel
	t, err := v.getVoorspellingTabel(d.ID, w)
	if err != nil {
		return Voorspelling{}, err
	}
	// Now loop through the voorspellingTabel to get the voorspellingRijen
	vRijen := t.getVoorspellingRijen()

	for i, rij := range vRijen {

		datum, err := rij.datum(v.time)
		if err != nil {
			log.Debug("Error on getting datum from rij")
			continue
		}
		if datum != w.Datum {
			log.Debugf("Found datum %s is not equal to wedstrijddatum: %s", datum.String(), w.Datum.String())
			continue
		}
		log.Debugf("Found datum %s is equal to wedstrijddatum: %s", datum.String(), w.Datum.String())

		// Als wvdw dan voegen we een extra rij toe
		rijen := []voorspellingRij{rij}
		if rij.DOM.HasClass("wvdw") {
			if i+1 <= len(vRijen) {
				if vRijen[i+1].DOM.HasClass("wvdw") {
					rijen = append(rijen, vRijen[i+1])
				}
			}
		}
		//we gebruiken tempVoorspelling, omdat we nog moeten checken of dit echt de wedstrijd is die we zoeken
		//het kan zijn dat er anders een wedstrijd die op hetzelfde tijdstip start gepakt wordt.
		tempVoorspelling, err := newVoorspelling(w.Competitie, v.time, rijen...)

		if err != nil {
			return Voorspelling{}, err
		}

		if tempVoorspelling.ThuisTeam != w.ThuisTeam && tempVoorspelling.UitTeam != w.UitTeam {
			log.Debugf("Wedstrijd according to voorspelling: %s - %s", tempVoorspelling.ThuisTeam, tempVoorspelling.UitTeam)
			log.Debugf("Wedstrijd according to Wedstrijd: %s - %s", w.ThuisTeam, w.UitTeam)
			continue
		}
		voorspelling = tempVoorspelling
	}
	return voorspelling, nil
}

//getVoorspellingTabel returns the voorspellingtabel for a user
func (v *voorspellingenService) getVoorspellingTabel(id int, w Wedstrijd) (voorspellingTabel, error) {
	var elem colly.HTMLElement
	v.client.OnHTML("table.voorspellingen", func(tabel *colly.HTMLElement) {
		// maak een voorspellingTabel van tabel, om receiver methods toe te kunnen passen
		elem = *tabel
	})

	url := fmt.Sprintf("%sdeelnemer/%d/voorspellingen/%s", v.baseURL, id, w.Competitie)
	log.Infof("Visiting url: %s", url)
	v.client.Visit(url)
	return voorspellingTabel{&elem}, nil
}

//getVoorspellingRijen fetches the rows from a voorspellingTabel containing info on voorspellingen
func (v voorspellingTabel) getVoorspellingRijen() []voorspellingRij {
	rij := "tr:not(:first-child)"
	var vRijen []voorspellingRij
	v.ForEach(rij, func(_ int, r *colly.HTMLElement) {
		//Maak een nieuwe VoorspellingRij aan, zodat we daar receiver methods op toe kunnen passen
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
		log.Debugf("Element %s is geen voorspellingrij", strings.Fields(e.Text))
		return voorspellingRij{}, fmt.Errorf("is geen voorspellingrij")
	}
	return voorspellingRij{e}, nil
}

//newVoorspelling creates a voorspelling from a voorspellingrij
func newVoorspelling(competitie string, baseDate time.Time, vRij ...voorspellingRij) (Voorspelling, error) {

	v := Voorspelling{}
	v.ThuisTeam = vRij[0].thuisTeam()
	v.UitTeam = vRij[0].uitTeam()

	log.Infof("Getting voorspelling %s - %s", v.ThuisTeam, v.UitTeam)

	datum, err := vRij[0].datum(baseDate)
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
	re, err := regexp.Compile(`\((.*?)\)`)
	if err != nil {
		log.Fatal(err)
	}
	tekst := r.ChildText("td:nth-child(2) .vp-team")

	team := re.ReplaceAllString(tekst, " ")

	return strings.TrimSpace(team)
}

func (r *voorspellingRij) uitTeam() string {
	re, err := regexp.Compile(`\((.*?)\)`)
	if err != nil {
		log.Fatal(err)
	}
	tekst := r.ChildText("td:nth-child(3) .vp-team")

	team := re.ReplaceAllString(tekst, " ")

	return strings.TrimSpace(team)
}

func (r *voorspellingRij) thuisDoelpuntenMaker() *string {
	d := r.ChildText("td:nth-child(2)")

	return &d
}

func (r *voorspellingRij) uitDoelpuntenMaker() *string {
	d := r.ChildText("td:nth-child(3)")
	return &d
}

func (r *voorspellingRij) doelpuntenThuis() *int {
	rawTekst := r.ChildText("td:nth-child(4)")
	if rawTekst == "-" {
		return nil
	}
	if rawTekst == "" {
		return nil
	}

	sanitizedTekst := strings.TrimSpace(strings.ReplaceAll(rawTekst, r.ChildText(".vp-uitslag"), ""))
	stringGoals := strings.TrimSpace(strings.Split(sanitizedTekst, "-")[0])
	i, err := strconv.Atoi(stringGoals)
	if err != nil {
		log.Panic()
	}
	return &i
}

func (r *voorspellingRij) doelpuntenUit() *int {
	rawTekst := r.ChildText("td:nth-child(4)")

	if rawTekst == "-" {
		return nil
	}

	if rawTekst == "" {
		return nil
	}

	sanitizedTekst := strings.TrimSpace(strings.ReplaceAll(rawTekst, r.ChildText(".vp-uitslag"), ""))
	stringGoals := strings.TrimSpace(strings.Split(sanitizedTekst, "-")[1])
	i, err := strconv.Atoi(stringGoals)
	if err != nil {
		log.Panic()
	}
	return &i
}

func (r *voorspellingRij) wvdw() bool {
	return r.DOM.HasClass("wvdw")
}

func isVoorspellingRij(e *colly.HTMLElement) (b bool) {
	//Dit is geen briljant criterium, een losse cel met alleen '.vp-team' zou nu ook als wedstrijdRij gezien worden
	return e.ChildText(".vp-team") != "" || e.DOM.HasClass("wvdw")

}

//datum extracts the date from a voorspellingRij
func (r *voorspellingRij) datum(baseDate time.Time) (t time.Time, err error) {

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

	t = time.Date(baseDate.Year(), baseDate.Month(), baseDate.Day(), uur, minuten, 0, 0, baseDate.Location())

	switch dag {

	case "Gisteren":
		t = t.AddDate(0, 0, -1)
	case "Vandaag":
		// do nothing
	case "Morgen":
		t = t.AddDate(0, 0, 1)
	default:
		// If no match, just add 20 years.
		t = t.AddDate(20, 0, 0)
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
