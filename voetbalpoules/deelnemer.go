package voetbalpoules

import (
	"time"
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
	Uitslag              string
	Wvdw                 bool
	ThuisDoelpuntenMaker string
	UitDoelpuntenMaker   string
}
