package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

var VoorspellingNoorwegenNoordIerland = Voorspelling{
	Datum:                time.Date(2022, 7, 7, 21, 0, 0, 0, time.Local),
	ThuisTeam:            "Noorwegen",
	UitTeam:              "Noord Ierland",
	DoelpuntenThuis:      3,
	DoelpuntenUit:        1,
	Wvdw:                 true,
	ThuisDoelpuntenMaker: "Minde",
	UitDoelpuntenMaker:   "Beattie",
}

var VoorspellingDuitslandDenemarken = Voorspelling{
	Datum:                time.Date(2022, 7, 8, 21, 0, 0, 0, time.Local),
	ThuisTeam:            "Duitsland",
	UitTeam:              "Denemarken",
	DoelpuntenThuis:      1,
	DoelpuntenUit:        0,
	Wvdw:                 false,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
}

var VoorspellingLandALandB = Voorspelling{
	Datum:                time.Date(2022, 7, 9, 21, 0, 0, 0, time.Local),
	ThuisTeam:            "Land A",
	UitTeam:              "Land B",
	DoelpuntenThuis:      0,
	DoelpuntenUit:        0,
	Wvdw:                 false,
	ThuisDoelpuntenMaker: "",
	UitDoelpuntenMaker:   "",
}

var VoorspellingNederlandZweden = Voorspelling{
	Datum:                time.Date(2022, 7, 9, 21, 0, 0, 0, time.Local),
	ThuisTeam:            "Nederland",
	UitTeam:              "Zweden",
	DoelpuntenThuis:      2,
	DoelpuntenUit:        1,
	Wvdw:                 true,
	ThuisDoelpuntenMaker: "Miedema",
	UitDoelpuntenMaker:   "Blackstenius",
}

func TestGetVoorspellingen(t *testing.T) {
	ts := newVoorspellingenTestServer()
	defer ts.Close()

	client := NewClient(ts.URL)
	cases := []struct {
		description string
		w           Wedstrijd
		baseTime    time.Time
		expected    []Voorspelling
	}{
		{
			baseTime:    time.Date(2022, 7, 8, 20, 30, 0, 0, time.Local),
			description: "Wedstrijd Gisteren",
			w:           NoorwegenNoordIerland,
			expected:    []Voorspelling{VoorspellingNoorwegenNoordIerland},
		},
		{
			baseTime:    time.Date(2022, 7, 8, 20, 30, 0, 0, time.Local),
			description: "Wedstrijd vandaag tussen nu en een half uur",
			w:           DuitslandDenemarken,
			expected:    []Voorspelling{VoorspellingDuitslandDenemarken},
		},
		{
			baseTime:    time.Date(2022, 7, 8, 20, 30, 0, 0, time.Local),
			description: "Wedstrijd Nederland - Zweden morgen",
			w:           NederlandZweden,
			expected:    []Voorspelling{VoorspellingNederlandZweden},
		},
		{
			baseTime:    time.Date(2022, 7, 8, 20, 30, 0, 0, time.Local),
			description: "Wedstrijd Land A - Land B (fictieve wedstrijd die tegelijkertijd gespeeld wordt met Nederland - Zweden)",
			w:           LandALandB,
			expected:    []Voorspelling{VoorspellingLandALandB},
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			client.time = tt.baseTime
			result, _ := client.Voorspellingen.Get("1", tt.w)
			if !reflect.DeepEqual(tt.expected, result) {
				fmt.Print("Result: ")
				fmt.Println(result)
				fmt.Print("Expected: ")
				fmt.Println(tt.expected)
				t.Error("Is not equal")
			}
		})
	}
}

func newVoorspellingenTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/deelnemer/1/voorspellingen/ek_vrouwen_2022", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
		<!DOCTYPE html>
		<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="nl" lang="nl">
		<head>
		
		<table class="voorspellingen">
		<tbody><tr>
				<th>Datum</th>
				<th>Thuisteam</th>
				<th>Uitteam</th>
				<th colspan="2">
						<div class="left">Voorspelling</div>
						<div class="right">Punten</div>
				</th>
				<th></th>
		</tr>
		
		<tr class="titel exclude toggle collapse">
				<td colspan="6" title="Klik hier om je voorspellingen voor deze ronde te tonen/ verbergen.">
						<h2>
	Groepsfase 1e ronde                                    </h2>
				</td>
		</tr>
	<tr>
	<td class="nowrap">
	Morgen	<div>21:00</div>
	
	</td>
	<td>		    
		<img alt="Engeland" src="Voorspellingen%20Ajvdt22_files/214221.svg">
	<div class="vp-team">
		Engeland
			<span>(A)</span>
	</div>
	</td>
	<td>
		<img alt="Oostenrijk" src="Voorspellingen%20Ajvdt22_files/214222.svg">
	<div class="vp-team">
		Oostenrijk
			<span>(A)</span>
	</div>
	</td>
	<td class="nowrap">
	
	2
		-
	0
						<div class="vp-uitslag">
					1 - 0
			</div>
	</td>
	<td>
	7
	</td>
	<td>
	<a rel="nofollow" class="noprint" tabindex="-1" href="javascript:void(0)">
		<img title="Statistieken" alt="Statistieken" src="Voorspellingen%20Ajvdt22_files/icon-info-s.png" data-edit="0" data-id="85468" class="statistieken">
	</a>
	</td>
	</tr>
	<tr class="wvdw" title="Wedstrijd van de dag">
	<td class="nowrap">
	Gisteren	<div>21:00</div>
	
	</td>
	<td>		    
		<img alt="Noorwegen" src="Voorspellingen%20Ajvdt22_files/214223.svg">
	<div class="vp-team">
		Noorwegen
			<span>(A)</span>
	</div>
	</td>
	<td>
		<img alt="Noord Ierland" src="Voorspellingen%20Ajvdt22_files/214224.svg">
	<div class="vp-team">
		Noord Ierland
			<span>(A)</span>
	</div>
	</td>
	<td class="nowrap">
	
	3
		-
	1
						<div class="vp-uitslag">
					4 - 1
			</div>
	</td>
	<td>
	8
	</td>
	<td>
	<a rel="nofollow" class="noprint" tabindex="-1" href="javascript:void(0)">
		<img title="Statistieken" alt="Statistieken" src="Voorspellingen%20Ajvdt22_files/icon-info-s.png" data-edit="0" data-id="85469" class="statistieken">
	</a>
	</td>
	</tr>
	<tr class="wvdw doelpunten" title="Wedstrijd van de dag">
	<td>1e doelpunt</td>
	<td>
		<span class="f">
			Minde
		</span>
	</td>
	<td colspan="3">
		<span class="f">
			Beattie
		</span>
	</td>
	<td></td>
	</tr>
	<tr>
	<td class="nowrap">
	Vandaag	<div>18:00</div>
	
	</td>
	<td>		    
		<img alt="Spanje" src="Voorspellingen%20Ajvdt22_files/214236.svg">
	<div class="vp-team">
		Spanje
			<span>(B)</span>
	</div>
	</td>
	<td>
		<img alt="Finland" src="Voorspellingen%20Ajvdt22_files/214237.svg">
	<div class="vp-team">
		Finland
			<span>(B)</span>
	</div>
	</td>
	<td class="nowrap">
	
	1
		-
	0
						<div class="vp-uitslag">
					4 - 1
			</div>
	</td>
	<td>
	5
	</td>
	<td>
	<a rel="nofollow" class="noprint" tabindex="-1" href="javascript:void(0)">
		<img title="Statistieken" alt="Statistieken" src="Voorspellingen%20Ajvdt22_files/icon-info-s.png" data-edit="0" data-id="85474" class="statistieken">
	</a>
	</td>
	</tr>
	<tr>
	<td class="nowrap">
	Vandaag	<div>21:00</div>
	
	</td>
	<td>		    
		<img alt="Duitsland" src="Voorspellingen%20Ajvdt22_files/214238.svg">
	<div class="vp-team">
		Duitsland
			<span>(B)</span>
	</div>
	</td>
	<td>
		<img alt="Denemarken" src="Voorspellingen%20Ajvdt22_files/214239.svg">
	<div class="vp-team">
		Denemarken
			<span>(B)</span>
	</div>
	</td>
	<td class="nowrap">

	1
		-
	0
						<div class="vp-uitslag">
					4 - 0
			</div>
	</td>
	
	</td>
	<td>
	<a rel="nofollow" class="noprint" tabindex="-1" href="javascript:void(0)">
		<img title="Statistieken" alt="Statistieken" src="Voorspellingen%20Ajvdt22_files/icon-info-s.png" data-edit="" data-id="85475" class="statistieken">
	</a>
	</td>
	</tr>
	<tr>
	<td class="nowrap">
	Morgen	<div>18:00</div>
	
	</td>
	<td>		    
		<img alt="Portugal" src="Voorspellingen%20Ajvdt22_files/215720.svg">
	<div class="vp-team">
		Portugal
			<span>(C)</span>
	</div>
	</td>
	<td>
		<img alt="Zwitserland" src="Voorspellingen%20Ajvdt22_files/214241.svg">
	<div class="vp-team">
		Zwitserland
			<span>(C)</span>
	</div>
	</td>
	<td class="nowrap">
	
	2
		-
	1
			</td>
	<td>
	
	</td>
	<td>
	<a rel="nofollow" class="noprint" tabindex="-1" href="javascript:void(0)">
		<img title="Statistieken" alt="Statistieken" src="Voorspellingen%20Ajvdt22_files/icon-info-s.png" data-edit="" data-id="85480" class="statistieken">
	</a>
	</td>
	</tr>
	<tr class="wvdw" title="Wedstrijd van de dag">
	<td class="nowrap">
	Morgen	<div>21:00</div>
	
	</td>
	<td>		    
		<img alt="Nederland" src="Voorspellingen%20Ajvdt22_files/214242.svg">
	<div class="vp-team">
		Nederland
			<span>(C)</span>
	</div>
	</td>
	<td>
		<img alt="Zweden" src="Voorspellingen%20Ajvdt22_files/214243.svg">
	<div class="vp-team">
		Zweden
			<span>(C)</span>
	</div>
	</td>
	<td class="nowrap">
	
	2
		-
	1
			</td>
	<td>
	
	</td>
	<td>
	<a rel="nofollow" class="noprint" tabindex="-1" href="javascript:void(0)">
		<img title="Statistieken" alt="Statistieken" src="Voorspellingen%20Ajvdt22_files/icon-info-s.png" data-edit="" data-id="85481" class="statistieken">
	</a>
	</td>
	</tr>
	<tr class="wvdw doelpunten" title="Wedstrijd van de dag">
	<td>1e doelpunt</td>
	<td>
		<span class="">
			Miedema
		</span>
	</td>
	<td colspan="3">
		<span class="">
			Blackstenius
		</span>
	</td>
	<td></td>
	</tr>
	<tr>
	<td class="nowrap">
	Morgen	<div>21:00</div>
	
	</td>
	<td>		    
		<img alt="België" src="Voorspellingen%20Ajvdt22_files/214244.svg">
	<div class="vp-team">
		Land A
			<span>(D)</span>
	</div>
	</td>
	<td>
		<img alt="IJsland" src="Voorspellingen%20Ajvdt22_files/214245.svg">
	<div class="vp-team">
		Land B
			<span>(D)</span>
	</div>
	</td>
	<td class="nowrap">
	
	0
		-
	0
			</td>
	<td>
	
	</td>
	<td>
	<a rel="nofollow" class="noprint" tabindex="-1" href="javascript:void(0)">
		<img title="Statistieken" alt="Statistieken" src="Voorspellingen%20Ajvdt22_files/icon-info-s.png" data-edit="" data-id="85486" class="statistieken">
	</a>
	</td>
	</tr>
	<tr>
	<td class="nowrap">
	zo	<div>21:00</div>
	
	</td>
	<td>		    
		<img alt="Frankrijk" src="Voorspellingen%20Ajvdt22_files/214246.svg">
	<div class="vp-team">
		Frankrijk
			<span>(D)</span>
	</div>
	</td>
	<td>
		<img alt="Italië" src="Voorspellingen%20Ajvdt22_files/214247.svg">
	<div class="vp-team">
		Italië
			<span>(D)</span>
	</div>
	</td>
	<td class="nowrap">
	
	2
		-
	0
			</td>
	<td>
	
	</td>
	<td>
	<a rel="nofollow" class="noprint" tabindex="-1" href="javascript:void(0)">
		<img title="Statistieken" alt="Statistieken" src="Voorspellingen%20Ajvdt22_files/icon-info-s.png" data-edit="" data-id="85487" class="statistieken">
	</a>
	</td>
	</tr>
		<tr class="totaal exclude">
				<td colspan="6">
						Ronde totaal:
						<span class="vp-punten">20</span>
						<span class="vp-positie">
								<a href="https://www.voetbalpoules.nl/stand/ek_vrouwen_2022/groepswedstrijden_1e_ronde/12?userID=881053">
										(293e)
								</a>
						</span>
				</td>
		</tr>
		</tbody>
		</table>
		</body>
		</html>
		`))
	})
	return httptest.NewServer(mux)
}
