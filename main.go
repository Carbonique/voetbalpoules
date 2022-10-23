package main

import (
	"fmt"

	"github.com/Carbonique/voetbalpoules-messenger/voetbalpoules"
)

func main() {

	client := voetbalpoules.NewClient("https://www.voetbalpoules.nl/")

	d := client.Pool.GetDeelnemers(18173, "eredivisie")
	fmt.Printf("d: %v\n", d)
}
