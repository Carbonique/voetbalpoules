package main

import (
	"fmt"

	"github.com/Carbonique/voetbalpoules/client"
)

func main() {

	client := client.NewClient("https://www.voetbalpoules.nl/")

	d := client.Pool.GetDeelnemers(18173, "eredivisie")
	fmt.Printf("d: %v\n", d)
}
