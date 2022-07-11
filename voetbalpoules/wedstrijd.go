package voetbalpoules

import (
	"fmt"
	"log"
	"time"

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

func isWedstrijdRij(e *colly.HTMLElement) (b bool) {
	//Dit is geen briljant criterium, een losse cel met alleen '.vp-team' zou nu ook als wedstrijdRij gezien worden
	return e.ChildText(".vp-team") != ""

}

//NaarWedstrijd creates a Wedstrijd from a HTMLElement. If the HTMLElement cannot be converted into a Wedstrijd,
//NaarWedstrijd will return a nil value
func NaarWedstrijd(e *colly.HTMLElement) (w *Wedstrijd, err error) {

	if !isWedstrijdRij(e) {
		log.Println("Element is geen wedstrijdrij")
		return nil, nil
	}

	//Maak een nieuwe wedstrijdRij aan, zodat we daar receiver methods op toe kunnen passen
	wRij := wedstrijdRij{e}
	fmt.Println(wRij.ChildText(".vp-team"))

	return &Wedstrijd{}, err

}

//datum extracts the date from a wedstrijdRij
func (r *wedstrijdRij) datum() (t time.Time, err error) {

	cel, err := r.datumCel()
	if err != nil {
		return t, fmt.Errorf("datum: failed parsing %w", err)

	}
	fmt.Println(cel)
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

////NewWedstrijd maakt een wedstrijd struct op basis van ingevoerde info
//func NewWedstrijd(datum time.Time, competitie string, ronde string, thuisTeam string, uitTeam string, wvdw bool, thuisDoelpuntenMaker string, uitDoelpuntenMaker string) (w *Wedstrijd, err error) {
//
//	//w.Datum =
//	//	w.Competitie =
//	//	w.Ronde =
//	//	w.ThuisTeam = getTeam(e)
//	//	w.UitTeam = getTeam(e)
//	//	w.Wvdw =
//	//	w.ThuisDoelpuntenMaker =
//	//	w.UitDoelpuntenMaker =
//
//	return w, nil
//
//}
