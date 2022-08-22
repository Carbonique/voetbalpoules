package voetbalpoules

import "time"

const competitie = "ek_vrouwen_2022"

var vandaag = time.Date(2022, 7, 10, 0, 0, 0, 0, time.Local)

var PortugalZwitserland = Wedstrijd{
	Datum:                time.Date(2022, 7, 9, 18, 0, 0, 0, time.Local),
	Competitie:           competitie,
	ThuisTeam:            "Portugal",
	UitTeam:              "Zwitserland",
	Uitslag:              "2 - 2",
	Wvdw:                 false,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
}

var NederlandZweden = Wedstrijd{
	Datum:                time.Date(2022, 7, 9, 21, 0, 0, 0, time.Local),
	Competitie:           competitie,
	ThuisTeam:            "Nederland",
	UitTeam:              "Zweden",
	Uitslag:              "1 - 1",
	Wvdw:                 true,
	ThuisDoelpuntenMaker: "Roord",
	UitDoelpuntenMaker:   "Andersson",
}

var BelgiëIJsland = Wedstrijd{
	Datum:                time.Date(2022, 7, 10, 18, 0, 0, 0, time.Local),
	Competitie:           competitie,
	ThuisTeam:            "België",
	UitTeam:              "IJsland",
	Uitslag:              "-",
	Wvdw:                 false,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
}

var FrankrijkItalië = Wedstrijd{
	Datum:                time.Date(2022, 7, 10, 21, 0, 0, 0, time.Local),
	Competitie:           competitie,
	ThuisTeam:            "Frankrijk",
	UitTeam:              "Italië",
	Uitslag:              "-",
	Wvdw:                 false,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
}

var OostenrijkNoordIerland = Wedstrijd{
	Datum:                time.Date(2022, 7, 11, 18, 0, 0, 0, time.Local),
	Competitie:           competitie,
	ThuisTeam:            "Oostenrijk",
	UitTeam:              "Noord Ierland",
	Uitslag:              "-",
	Wvdw:                 false,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
}

var EngelandNoorwegen = Wedstrijd{
	Datum:                time.Date(2022, 7, 11, 21, 0, 0, 0, time.Local),
	Competitie:           competitie,
	ThuisTeam:            "Engeland",
	UitTeam:              "Noorwegen",
	Uitslag:              "-",
	Wvdw:                 false,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
}
