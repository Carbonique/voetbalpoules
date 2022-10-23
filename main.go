package main

import (
	"fmt"

	"github.com/Carbonique/voetbalpoules/scraper"
)

func main() {

	client := scraper.NewClient("https://www.voetbalpoules.nl/")

	d := client.Pool.GetDeelnemers(18173, "eredivisie")
	fmt.Printf("d: %v\n", d)
}
