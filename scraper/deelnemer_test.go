package scraper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gocolly/colly/v2"
)

func TestNewDeelnemer(t *testing.T) {
	ts := newDeelnemerTestServer()
	defer ts.Close()

	client := NewClient(ts.URL)

	cases := []struct {
		description string
		expected    Deelnemer
	}{
		{
			description: "Test existing deelnemer with all fields filled",
			expected:    Deelnemer{1, "geel", 277},
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			client.Visit(ts.URL + "/poule")
			client.OnHTML("tr", func(e *colly.HTMLElement) {

				fmt.Println(e.Text)
				rij := deelnemerRij{e}
				result, _ := NewDeelnemer(rij)
				if !reflect.DeepEqual(tt.expected, result) {
					fmt.Print("Result: ")
					fmt.Println(result)
					fmt.Print("Expected: ")
					fmt.Println(tt.expected)
					t.Error("Is not equal")
				}
			})
		})
	}
}

func newDeelnemerTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/poule", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
		<tr title="Lid sinds 5-7-2022" class="">
                    <td>3.</td>
                    <td>
                            2.
                                <img src="/Content/Desktop/images/stand-d.png">
                    </td>
                    <td>
                            <a class="links" href="/deelnemer/1/voorspellingen/ek_vrouwen_2022">
        geel 
    </a>


                                <span class="vp-super"><img class="medaille" title="Rondewinnaar" src="/Content/images/award.png">(2)</span>
                    </td>
                    <td>
                        <div class="afkappuh">geel </div>
                    </td>
                        <td class="punten">
                            277
                                <div style="font-size: 9px; text-align: right; font-weight: normal;">
254                                </div>
                        </td>
                            <td style="font-size: 9px; text-align: right">
                                (+30)
                            </td>
                </tr>
		`))
	})
	return httptest.NewServer(mux)
}
