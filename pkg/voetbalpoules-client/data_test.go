package voetbalpoulesclient

import "time"

const competitie = "ek_vrouwen_2022"

var NoorwegenNoordIerland = Wedstrijd{
	Datum:                time.Date(2022, 7, 7, 21, 0, 0, 0, time.Local),
	Competitie:           competitie,
	ThuisTeam:            "Noorwegen",
	UitTeam:              "Noord Ierland",
	Uitslag:              "4 - 1",
	Wvdw:                 true,
	ThuisDoelpuntenMaker: "Blakstad",
	UitDoelpuntenMaker:   "Nelson",
}

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

var LandALandB = Wedstrijd{
	Datum:                time.Date(2022, 7, 9, 21, 0, 0, 0, time.Local),
	Competitie:           competitie,
	ThuisTeam:            "Land A",
	UitTeam:              "Land B",
	Uitslag:              "0 - 1",
	Wvdw:                 false,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
}

var LandCLandD = Wedstrijd{
	Datum:                time.Date(2022, 7, 9, 21, 0, 0, 0, time.Local),
	Competitie:           competitie,
	ThuisTeam:            "Land C",
	UitTeam:              "Land D",
	Uitslag:              "2 - 1",
	Wvdw:                 true,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
}

var LandELandF = Wedstrijd{
	Datum:                time.Date(2022, 7, 9, 21, 0, 0, 0, time.Local),
	Competitie:           competitie,
	ThuisTeam:            "Land E",
	UitTeam:              "Land F",
	Uitslag:              "2 - 1",
	Wvdw:                 true,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
}

var LandGLandH = Wedstrijd{
	Datum:                time.Date(2022, 7, 9, 21, 0, 0, 0, time.Local),
	Competitie:           competitie,
	ThuisTeam:            "Land G",
	UitTeam:              "Land H",
	Uitslag:              "2 - 0",
	Wvdw:                 false,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
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

var DuitslandDenemarken = Wedstrijd{
	Datum:                time.Date(2022, 7, 8, 21, 0, 0, 0, time.Local),
	Competitie:           competitie,
	ThuisTeam:            "Duitsland",
	UitTeam:              "Denemarken",
	Uitslag:              "4 - 0",
	Wvdw:                 false,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
}
