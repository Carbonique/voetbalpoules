package voetbalpoulesclient

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gocolly/colly/v2"
)

//wedstrijdService handles communication related to Wedstrijden
type wedstrijdService service

type Wedstrijd struct {
	Datum                time.Time
	Competitie           string
	ThuisTeam            string
	UitTeam              string
	Uitslag              string
	Wvdw                 bool
	ThuisDoelpuntenMaker string
	UitDoelpuntenMaker   string
}

// Get returns Wedstrijden for a Competitie within a specified timerange
func (w *wedstrijdService) get(competitie string, t1 time.Time, t2 time.Time) ([]Wedstrijd, error) {

	var wedstrijden []Wedstrijd
	var wedstrijd Wedstrijd

	// First fetch the wedstrijdTabel
	t, err := w.getWedstrijdTabel(competitie)
	if err != nil {
		return []Wedstrijd{}, err
	}
	// Now loop through the wedstrijdTabel to get the wedstrijdRijen
	wRijen := t.getWedstrijdRijen()
	for i, rij := range wRijen {

		datum, err := rij.datum(w.time)
		if err != nil {
			log.Debug("Error on getting datum from rij")
			continue
		}

		if !inTimeSpan(t1, t2, datum) {
			log.Debugf("Found wedstrijd at %s, but is not within range %s - %s", datum.Format(time.RFC822), t1.Format(time.RFC822), t2.Format(time.RFC822))
			continue
		}

		// Als wvdw dan voegen we een extra rij toe
		rijen := []wedstrijdRij{rij}
		if rij.DOM.HasClass("wvdw") {
			if i+1 <= len(wRijen) {
				rijen = append(rijen, wRijen[i+1])
			}
		}

		wedstrijd, err = newWedstrijd(competitie, w.time, rijen...)

		if err != nil {
			return []Wedstrijd{}, err
		}

		wedstrijden = append(wedstrijden, wedstrijd)

	}
	return wedstrijden, nil
}

type wedstrijdTabel struct {
	*colly.HTMLElement
}

//getWedstrijdTabel returns the wedstrijdTabel for a competitie
func (w *wedstrijdService) getWedstrijdTabel(c string) (wedstrijdTabel, error) {
	var elem colly.HTMLElement
	w.client.OnHTML("table.wedstrijden", func(tabel *colly.HTMLElement) {
		// maak een wedstrijdTabel van tabel, om receiver methods toe te kunnen passen
		elem = *tabel
	})

	url := fmt.Sprintf("%swedstrijd/index/%s", w.baseURL, c)
	log.Infof("Visiting url: %s", url)
	w.client.Visit(url)
	return wedstrijdTabel{&elem}, nil
}

func (w wedstrijdTabel) getWedstrijdRijen() []wedstrijdRij {
	rij := "tr:not(:first-child)"

	var wRijen []wedstrijdRij

	w.ForEach(rij, func(_ int, r *colly.HTMLElement) {
		//Maak een nieuwe wedstrijdRij aan, zodat we daar receiver methods op toe kunnen passen
		wRij, err := newWedstrijdRij(r)
		if err != nil {
			return
		}
		wRijen = append(wRijen, wRij)
	})

	return wRijen
}

type wedstrijdRij struct {
	*colly.HTMLElement
}

func isWedstrijdRij(e *colly.HTMLElement) (b bool) {
	//Dit is geen briljant criterium, een losse cel met alleen '.vp-team' zou nu ook als wedstrijdRij gezien worden
	return e.ChildText(".vp-team") != "" || e.DOM.HasClass("wvdw")

}

func newWedstrijdRij(e *colly.HTMLElement) (wedstrijdRij, error) {
	if !isWedstrijdRij(e) {
		log.Debugf("Element %s is geen wedstrijdrij", strings.Fields(e.Text))
		return wedstrijdRij{}, fmt.Errorf("is geen wedstrijdrij")
	}
	return wedstrijdRij{e}, nil
}

//newWedstrijd creates a Wedstrijd from a wedstrijdrij
func newWedstrijd(competitie string, vandaag time.Time, wRij ...wedstrijdRij) (Wedstrijd, error) {

	w := Wedstrijd{}
	w.ThuisTeam = wRij[0].team(2)
	w.UitTeam = wRij[0].team(3)

	log.Infof("Getting wedstrijd %s - %s", w.ThuisTeam, w.UitTeam)

	datum, err := wRij[0].datum(vandaag)
	if err != nil {
		return Wedstrijd{}, err
	}
	w.Datum = datum

	w.Competitie = competitie
	w.Uitslag = wRij[0].uitslag()

	w.Wvdw = wRij[0].wvdw()
	if len(wRij) > 1 {
		w.ThuisDoelpuntenMaker = wRij[1].doelpuntenMaker(2)
		w.UitDoelpuntenMaker = wRij[1].doelpuntenMaker(3)
	}

	return w, nil

}

func (r *wedstrijdRij) team(cel int) string {
	t := r.ChildText(fmt.Sprintf("td:nth-child(%d) .vp-team", cel))
	return t
}

func (r *wedstrijdRij) doelpuntenMaker(cel int) string {
	d := r.ChildText(fmt.Sprintf("td:nth-child(%d)", cel))
	return d
}

func (r *wedstrijdRij) uitslag() string {
	u := r.ChildText("td:nth-child(4)")
	return u
}

func (r *wedstrijdRij) wvdw() bool {
	return r.DOM.HasClass("wvdw")
}

//datum extracts the date from a wedstrijdRij
func (r *wedstrijdRij) datum(vandaag time.Time) (t time.Time, err error) {

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
func (r *wedstrijdRij) datumCel() (s string, err error) {
	cel := r.ChildText("td.nowrap:first-child")

	if cel == "" {
		return "", fmt.Errorf("datum: failed parsing %w", err)
	}

	return cel, err

}

func inTimeSpan(t1 time.Time, t2 time.Time, check time.Time) bool {
	var start time.Time
	var end time.Time

	switch {
	case t1.Before(t2):
		start = t1
		end = t2

	case t1.After(t2):
		start = t2
		end = t1

	case t1.Equal(t2):
		start = t1
		end = t2
	}

	log.Debugf("Checking if %s is within %s and %s", check.String(), start.String(), end.String())
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}
