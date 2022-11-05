package voetbalpoulesclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var Deelnemer1 = Deelnemer{
	ID:          10,
	Naam:        "Deelnemer1",
	Punten:      443,
	PuntenRonde: 14,
}

var Deelnemer2 = Deelnemer{
	ID:          11,
	Naam:        "Deelnemer2",
	Punten:      400,
	PuntenRonde: 14,
}

var Deelnemer3 = Deelnemer{
	ID:          12,
	Naam:        "Deelnemer3",
	Punten:      300,
	PuntenRonde: 14,
}

func TestGetDeelnemers(t *testing.T) {
	ts := newPoolTestServer()
	defer ts.Close()

	client := NewClient(ts.URL)
	cases := []struct {
		description string
		expected    []Deelnemer
	}{
		{
			description: "Pool met deelnemers niet gesorteerd",
			expected:    []Deelnemer{Deelnemer1, Deelnemer3, Deelnemer2},
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result := client.pool.getDeelnemers(1, "eredivisie")
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

func TestGetStand(t *testing.T) {
	ts := newPoolTestServer()
	defer ts.Close()

	client := NewClient(ts.URL)
	cases := []struct {
		description string
		expected    []Deelnemer
	}{
		{
			description: "Pool met deelnemers reeds gesorteerd in html tabel",
			expected:    []Deelnemer{Deelnemer1, Deelnemer2, Deelnemer3},
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result := client.pool.getStand(1, "eredivisie")
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

func newPoolTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/poule/1/stand/eredivisie", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
		<!DOCTYPE html>
		<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="nl" lang="nl">
		<head>
			<table class="stand">
					<tbody>
						<tr>
								<th width="10%">Nr.</th>
								<th width="12%">Vr.</th>
								<th width="62%">Deelnemer</th>
								<th colspan="2" width="16%">
									Punten
								</th>
						</tr>
						<tr title="Lid sinds 8-8-2021" class="">
								<td>1.</td>
								<td>
									2.
									<img src="/Content/Desktop/images/stand-s.png">
								</td>
								<td>
									<img src="/logo.png" class="vp-userfoto klein" title="">
									<div class="rank-deelnemer">
											<a href="/deelnemer/10/voorspellingen/eredivisie">Deelnemer1</a><br>
											<span class="naam">Peter</span>
									</div>
									<img class="medaille" title="Rondewinnaar" src="/Content/images/award.png">
								</td>
								<td class="punten">
									443
								</td>
								<td style="font-size: 9px; text-align: right">
									(+14)
								</td>
						</tr>
						<tr title="Lid sinds 8-8-2021" class="">
								<td>1.</td>
								<td>
									2.
									<img src="/Content/Desktop/images/stand-s.png">
								</td>
								<td>
									<img src="/logo.png" class="vp-userfoto klein" title="">
									<div class="rank-deelnemer">
											<a href="/deelnemer/12/voorspellingen/eredivisie">Deelnemer3</a><br>
											<span class="naam">Piet</span>
									</div>
								</td>
								<td class="punten">
									300
								</td>
								<td style="font-size: 9px; text-align: right">
									(+14)
								</td>
						</tr>
						<tr title="Lid sinds 8-8-2021" class="">
								<td>1.</td>
								<td>
									2.
									<img src="/Content/Desktop/images/stand-s.png">
								</td>
								<td>
									<img src="/logo.png" class="vp-userfoto klein" title="">
									<div class="rank-deelnemer">
											<a href="/deelnemer/11/voorspellingen/eredivisie">Deelnemer2</a><br>
											<span class="naam">Jan</span>
									</div>
								</td>
								<td class="punten">
									400
								</td>
								<td style="font-size: 9px; text-align: right">
									(+14)
								</td>
						</tr>			
					</tbody>
			</table>
			</body>
		</html>`))
	})
	return httptest.NewServer(mux)
}
