package voetbalpoules

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gocolly/colly/v2"
)

//WedstrijdService handles communication related to Wedstrijden
type WedstrijdService service

type wedstrijdRij struct {
	*colly.HTMLElement
}

type Wedstrijd struct {
	Datum                time.Time
	Competitie           string
	Ronde                string
	ThuisTeam            string
	UitTeam              string
	Uitslag              string
	Wvdw                 bool
	ThuisDoelpuntenMaker string
	UitDoelpuntenMaker   string
}

// GetWedstrijden returns Wedstrijden for a Competitie within a specified timerange
func (w *WedstrijdService) Get(competitie string, t1 time.Time, t2 time.Time) []Wedstrijd {

	var wedstrijden []Wedstrijd
	fmt.Println("out")
	w.OnHTML("table.wedstrijden", func(wedstrijdenTabel *colly.HTMLElement) {
		fmt.Println("in")
		rij := "tr:not(:first-child)"
		wedstrijdenTabel.ForEach(rij, func(_ int, r *colly.HTMLElement) {

			//Maak een nieuwe wedstrijdRij aan, zodat we daar receiver methods op toe kunnen passen
			wRij, err := newWedstrijdRij(r)
			if err != nil {
				return
			}
			datum, err := wRij.datum()
			if err != nil {
				log.Error("Error on datum")
				return
			}

			if !inTimeSpan(t1, t2, datum) {
				log.Debugf("Found wedstrijd at %s, but is not within range %s - %s", datum.Format(time.RFC822), t1.Format(time.RFC822), t2.Format(time.RFC822))
				return
			}

			wedstrijd, err := NewWedstrijd(wRij, competitie, "ronde1")

			if err != nil {
				log.Error("Error in getting wedstrijd")
			}
			if wedstrijd == nil {
				log.Debug("Wedstrijd == nil")
				//continue
				return
			}

			wedstrijden = append(wedstrijden, *wedstrijd)
		})

	})

	url := fmt.Sprintf("%swedstrijd/index/%s", w.baseURL, competitie)
	w.Visit(url)
	fmt.Println(url)
	return wedstrijden
}

func isWedstrijdRij(e *colly.HTMLElement) (b bool) {
	//Dit is geen briljant criterium, een losse cel met alleen '.vp-team' zou nu ook als wedstrijdRij gezien worden
	return e.ChildText(".vp-team") != ""

}

func newWedstrijdRij(e *colly.HTMLElement) (wedstrijdRij, error) {
	if !isWedstrijdRij(e) {
		log.Errorf("Element %s is geen wedstrijdrij", strings.Fields(e.Text))
		return wedstrijdRij{}, fmt.Errorf("is geen wedstrijdrij")
	}
	return wedstrijdRij{e}, nil
}

//NaarWedstrijd creates a Wedstrijd from a HTMLElement. If the HTMLElement cannot be converted into a Wedstrijd,
//NaarWedstrijd will return a nil value
func NewWedstrijd(wRij wedstrijdRij, competitie string, ronde string) (*Wedstrijd, error) {

	w := &Wedstrijd{}
	w.ThuisTeam = wRij.thuisTeam()
	w.UitTeam = wRij.uitTeam()

	log.Infof("Getting wedstrijd %s - %s", w.ThuisTeam, w.UitTeam)

	datum, err := wRij.datum()
	if err != nil {
		return &Wedstrijd{}, err
	}
	w.Datum = datum

	w.Competitie = competitie
	w.Ronde = ronde
	w.Wvdw = wRij.wvdw()
	w.Uitslag = wRij.uitslag()
	w.ThuisDoelpuntenMaker = "jaap"
	w.UitDoelpuntenMaker = "kees"

	return w, nil

}

func (r *wedstrijdRij) thuisTeam() string {
	t := r.ChildText("td:nth-child(2) .vp-team")

	return t
}

func (r *wedstrijdRij) uitTeam() string {
	t := r.ChildText("td:nth-child(3) .vp-team")
	return t
}

func (r *wedstrijdRij) uitslag() string {
	u := r.ChildText("td:nth-child(4)")
	return u
}

func (r *wedstrijdRij) wvdw() bool {
	return r.DOM.HasClass("wvdw")
}

//datum extracts the date from a wedstrijdRij
func (r *wedstrijdRij) datum() (t time.Time, err error) {

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

	vandaag := time.Now()
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

func inTimeSpan(start, end, check time.Time) bool {
	log.Debugf("Checking if %s is within %s and %s", check.String(), start.String(), end.String())
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}
