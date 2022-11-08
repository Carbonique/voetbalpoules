package voetbalpoulesclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestWedstrijdenGet(t *testing.T) {
	ts := newWedstrijdenTestServer()
	defer ts.Close()
	vandaag := time.Date(2022, 7, 10, 0, 0, 0, 0, time.Local)

	client := NewClient(ts.URL)
	client.time = vandaag

	cases := []struct {
		description string
		competitie  string
		t1          time.Time
		t2          time.Time
		expected    []Wedstrijd
	}{
		{
			description: "Wedstrijd gisteren tussen nu en een half uur",
			competitie:  competitie,
			t1:          time.Date(2022, 7, 9, 17, 45, 0, 0, time.Local),
			t2:          time.Date(2022, 7, 9, 18, 15, 0, 0, time.Local),
			expected:    []Wedstrijd{PortugalZwitserland},
		},
		{
			description: "Wedstrijd gisteren in tijdsrange",
			competitie:  competitie,
			t1:          time.Date(2022, 7, 9, 17, 45, 0, 0, time.Local),
			t2:          time.Date(2022, 7, 9, 21, 45, 0, 0, time.Local),
			expected:    []Wedstrijd{PortugalZwitserland, NederlandZweden},
		},
		{
			description: "Wedstrijd gisteren in tijdsrange (t2 ouder dan t1)",
			competitie:  competitie,
			t1:          time.Date(2022, 7, 9, 21, 45, 0, 0, time.Local),
			t2:          time.Date(2022, 7, 9, 17, 45, 0, 0, time.Local),
			expected:    []Wedstrijd{PortugalZwitserland, NederlandZweden},
		},
		{
			description: "Wedstrijden vandaag",
			competitie:  competitie,
			t1:          time.Date(2022, 7, 10, 00, 00, 0, 0, time.Local),
			t2:          time.Date(2022, 7, 10, 23, 59, 0, 0, time.Local),
			expected:    []Wedstrijd{BelgiëIJsland, FrankrijkItalië},
		},
		{
			description: "Wedstrijden morgen",
			competitie:  competitie,
			t1:          time.Date(2022, 7, 11, 00, 00, 0, 0, time.Local),
			t2:          time.Date(2022, 7, 11, 23, 59, 0, 0, time.Local),
			expected:    []Wedstrijd{OostenrijkNoordIerland, EngelandNoorwegen},
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result, _ := client.wedstrijden.get(tt.competitie, tt.t1, tt.t2)
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

func newWedstrijdenTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/wedstrijd/index/ek_vrouwen_2022", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
		<!DOCTYPE html>
		<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="nl" lang="nl">
		<head>
		
		<table class="wedstrijden">
		<tr class="">
				<td class="nowrap">
						8 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Duitsland" src="https://vp-logos.azureedge.net/214238" />
						<label class="vp-team" for="">Duitsland</label>
				</td>
				<td title="">
								<img alt="Denemarken" src="https://vp-logos.azureedge.net/214239" />
						<label class="vp-team" for="">Denemarken</label>
				</td>
				<td class="right nowrap">
								 4 - 0
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						Gisteren	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Portugal" src="https://vp-logos.azureedge.net/215720" />
						<label class="vp-team" for="">Portugal</label>
				</td>
				<td title="">
								<img alt="Zwitserland" src="https://vp-logos.azureedge.net/214241" />
						<label class="vp-team" for="">Zwitserland</label>
				</td>
				<td class="right nowrap">
								 2 - 2
				</td>
		</tr>
		
		<tr class="wvdw">
				<td class="nowrap">
						Gisteren	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Nederland" src="https://vp-logos.azureedge.net/214242" />
						<label class="vp-team" for="">Nederland</label>
				</td>
				<td title="">
								<img alt="Zweden" src="https://vp-logos.azureedge.net/214243" />
						<label class="vp-team" for="">Zweden</label>
				</td>
				<td class="right nowrap">
								 1 - 1
				</td>
		</tr>
				<tr class="wvdw doelpunten">
				<td>1e doelpunt</td>
				<td>
					Roord
				</td>
				<td>
					Andersson
				</td>
				<td></td>
			</tr>
		
		<tr class="">
				<td class="nowrap">
						Vandaag	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Belgi&#235;" src="https://vp-logos.azureedge.net/214244" />
						<label class="vp-team" for="">Belgi&#235;</label>
				</td>
				<td title="">
								<img alt="IJsland" src="https://vp-logos.azureedge.net/214245" />
						<label class="vp-team" for="">IJsland</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						Vandaag	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Frankrijk" src="https://vp-logos.azureedge.net/214246" />
						<label class="vp-team" for="">Frankrijk</label>
				</td>
				<td title="">
								<img alt="Itali&#235;" src="https://vp-logos.azureedge.net/214247" />
						<label class="vp-team" for="">Itali&#235;</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
							<tr>
								<td colspan="4">
																<h2>
																		EK Vrouwen 2022
																				<span style="font-size:0.9em;">Groepsfase 2e ronde</span>
																																</h2>
														</td>
												</tr>
		
		<tr class="">
				<td class="nowrap">
						Morgen	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Oostenrijk" src="https://vp-logos.azureedge.net/214222" />
						<label class="vp-team" for="">Oostenrijk</label>
				</td>
				<td title="">
								<img alt="Noord Ierland" src="https://vp-logos.azureedge.net/214224" />
						<label class="vp-team" for="">Noord Ierland</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						Morgen	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Engeland" src="https://vp-logos.azureedge.net/214221" />
						<label class="vp-team" for="">Engeland</label>
				</td>
				<td title="">
								<img alt="Noorwegen" src="https://vp-logos.azureedge.net/214223" />
						<label class="vp-team" for="">Noorwegen</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		</table>
		</body>
		</html>
		`))
	})
	return httptest.NewServer(mux)
}
