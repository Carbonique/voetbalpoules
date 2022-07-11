package voetbalpoules

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly/v2"
)

//CompetitieService handles communication related to Competities
type CompetitieService service

// GetWedstrijden returns Wedstrijden for a Competitie that fall within a specified timerange
func (w *WedstrijdService) GetWedstrijden(competitie string, t1 time.Time, t2 time.Time) string {

	w.OnHTML("table.wedstrijden", func(wedstrijdenTabel *colly.HTMLElement) {

		rij := "tr:not(:first-child)"
		wedstrijdenTabel.ForEachWithBreak(rij, func(_ int, r *colly.HTMLElement) bool {
			w, err := NaarWedstrijd(r)

			if err != nil {
				log.Panic("GetWedstrijden: error in NaarWedstrijd()")
			}
			if w == nil {
				//continue
				return true
			}

			return true
		})

	})

	url := fmt.Sprintf("%swedstrijd/index/%s", w.url, competitie)
	w.Visit(url)

	return "ja"
}
