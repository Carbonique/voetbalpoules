package voetbalpoulestelegram

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	voetbalpoulesclient "github.com/Carbonique/voetbalpoules/pkg/voetbalpoules-client"
)

var Deelnemer1 = voetbalpoulesclient.Deelnemer{
	Naam:        "Deelnemer1",
	ID:          1,
	Punten:      20,
	PuntenRonde: 1,
}

var Deelnemer2 = voetbalpoulesclient.Deelnemer{
	Naam:        "Deelnemer2",
	ID:          2,
	Punten:      24,
	PuntenRonde: 1,
}

var Deelnemer3 = voetbalpoulesclient.Deelnemer{
	Naam:        "Deelnemer3",
	ID:          3,
	Punten:      20,
	PuntenRonde: 4,
}

var Deelnemer4 = voetbalpoulesclient.Deelnemer{
	Naam:        "Deelnemer4",
	ID:          5,
	Punten:      20,
	PuntenRonde: 4,
}

var Deelnemer5 = voetbalpoulesclient.Deelnemer{
	Naam:        "Deelnemer5",
	ID:          5,
	Punten:      20,
	PuntenRonde: 4,
}
var Deelnemer6 = voetbalpoulesclient.Deelnemer{
	Naam:        "Deelnemer6",
	ID:          6,
	Punten:      20,
	PuntenRonde: 4,
}

func stringPointer(s string) *string {
	return &s
}

func intPointer(i int) *int {
	return &i
}

var Voorspelling3_1 = voetbalpoulesclient.Voorspelling{
	DoelpuntenThuis: intPointer(3),
	DoelpuntenUit:   intPointer(1),
}

var Voorspelling1_0 = voetbalpoulesclient.Voorspelling{
	DoelpuntenThuis: intPointer(1),
	DoelpuntenUit:   intPointer(0),
}

var Voorspelling0_1 = voetbalpoulesclient.Voorspelling{
	DoelpuntenThuis: intPointer(0),
	DoelpuntenUit:   intPointer(1),
}

var Voorspelling0_4 = voetbalpoulesclient.Voorspelling{
	DoelpuntenThuis: intPointer(0),
	DoelpuntenUit:   intPointer(4),
}

var Voorspelling0_0 = voetbalpoulesclient.Voorspelling{
	DoelpuntenThuis: intPointer(0),
	DoelpuntenUit:   intPointer(0),
}

var VoorspellingNietIngevuld = voetbalpoulesclient.Voorspelling{
	Datum:     time.Date(2022, 7, 9, 21, 0, 0, 0, time.Local),
	ThuisTeam: "Land A",
	UitTeam:   "Land B",
}

func TestSortDeelnemers(t *testing.T) {

	expectedList := []voetbalpoulesclient.Deelnemer{
		Deelnemer6,
		Deelnemer1,
		Deelnemer2,
		Deelnemer4,
		Deelnemer5,
		Deelnemer3,
	}
	m := make(map[voetbalpoulesclient.Deelnemer]voetbalpoulesclient.Voorspelling)

	m[Deelnemer1] = Voorspelling3_1
	m[Deelnemer2] = Voorspelling0_1
	m[Deelnemer3] = VoorspellingNietIngevuld
	m[Deelnemer4] = Voorspelling0_4
	m[Deelnemer5] = Voorspelling0_0
	m[Deelnemer6] = Voorspelling1_0

	sortedList := sortDeelnemers(m)

	for i := range sortedList {
		if !reflect.DeepEqual(sortedList[i], expectedList[i]) {
			fmt.Print("Result: ")
			fmt.Println(sortedList[i])
			fmt.Print("Expected: ")
			fmt.Println(expectedList[i])
			t.Error("Is not equal")
		}
	}
}
