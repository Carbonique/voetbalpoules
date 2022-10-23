package main

import (
	"fmt"
	"time"

	"github.com/Carbonique/voetbalpoules/client"
)

func main() {

	client := client.NewClient("https://www.voetbalpoules.nl/")

	t2 := time.Now()
	t1 := t2.Add(time.Hour * -4)
	w, err := client.Wedstrijden.Get("eredivisie", t1, t2)
	if err != nil {
		fmt.Errorf("Error")
	}

	d := client.Pool.GetDeelnemers(18173, "eredivisie")

	for _, d2 := range d {
		v, _ := client.Voorspellingen.Get(d2.ID, w[0])

		fmt.Printf("%s: %s - %s (Voorspelling: %d - %d) (Uitslag: %s) \n", d2.Naam, v.ThuisTeam, v.UitTeam, v.DoelpuntenThuis, v.DoelpuntenUit, w[0].Uitslag)
	}

}
