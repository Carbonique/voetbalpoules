package voetbalpoulestelegram

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	voetbalpoulesclient "github.com/Carbonique/voetbalpoules/pkg/voetbalpoules-client"
)

var NederlandZweden = voetbalpoulesclient.Wedstrijd{
	Datum:                time.Date(2022, 7, 9, 21, 0, 0, 0, time.Local),
	Competitie:           "ek_vrouwen_2022",
	ThuisTeam:            "Nederland",
	UitTeam:              "Zweden",
	Uitslag:              "1 - 1",
	Wvdw:                 true,
	ThuisDoelpuntenMaker: "Roord",
	UitDoelpuntenMaker:   "Andersson",
}

var PortugalZwitserland = voetbalpoulesclient.Wedstrijd{
	Datum:                time.Date(2022, 7, 9, 18, 0, 0, 0, time.Local),
	Competitie:           "ek_vrouwen_2022",
	ThuisTeam:            "Portugal",
	UitTeam:              "Zwitserland",
	Uitslag:              "2 - 2",
	Wvdw:                 false,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
}

var DeelnemerVoorspellingenMap = map[voetbalpoulesclient.Deelnemer]voetbalpoulesclient.Voorspelling{
	Deelnemer1: Voorspelling3_1,
	Deelnemer2: Voorspelling0_1,
	Deelnemer5: Voorspelling0_4,
	Deelnemer6: VoorspellingNietIngevuld,
}

var voorspeldeWedstrijdNietWVDW = voetbalpoulesclient.VoorspeldeWedstrijd{
	Wedstrijd:               PortugalZwitserland,
	DeelnemerVoorspellingen: DeelnemerVoorspellingenMap,
}

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
	Punten:      19,
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

func TestStandBericht(t *testing.T) {
	deelnemers := []voetbalpoulesclient.Deelnemer{
		Deelnemer2,
		Deelnemer1,
		Deelnemer4,
		Deelnemer3,
	}

	expectedTitel := fmt.Sprintf("*Stand %d - %d:*\n", time.Now().Day(), time.Now().Month())

	expectedInhoud := "1. Deelnemer2 24 (+1) \n" + "2. Deelnemer1 20 (+1) \n" + "3. Deelnemer4 20 (+4) \n" + "4. Deelnemer3 19 (+4) \n"

	expectedBericht := bericht{
		titel:  expectedTitel,
		inhoud: expectedInhoud,
	}

	bericht := newStandBericht(deelnemers)

	if !reflect.DeepEqual(expectedBericht, bericht) {
		fmt.Print("Result: ")
		fmt.Println(bericht)
		fmt.Print("Expected: ")
		fmt.Println(expectedBericht)
		t.Error("Is not equal")
	}

}

func TestUitslagBericht(t *testing.T) {

	expectedTitel := fmt.Sprintf("*Uitslag:\n Portugal - Zwitserland (2 - 2) 18:00*\n")
	expectedInhoud := "Deelnemer1 (3 - 1) " + "\n\n" + "Deelnemer2 (0 - 1) " + "\n" + "Deelnemer5 (0 - 4) " + "\n\n" + "Deelnemer6 (Niet ingevuld) " + "\n"

	expectedBericht := bericht{
		titel:  expectedTitel,
		inhoud: expectedInhoud,
	}

	b := newUitslagBericht(voorspeldeWedstrijdNietWVDW)

	if !reflect.DeepEqual(expectedBericht, b) {
		fmt.Print("Result: ")
		fmt.Println(b)
		fmt.Print("Expected: ")
		fmt.Println(expectedBericht)
		t.Error("Is not equal")
	}
}
