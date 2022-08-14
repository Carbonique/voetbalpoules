package main

import (
	"github.com/Carbonique/voetbalpoules-messenger/voetbalpoules"
)

func main() {

	client := voetbalpoules.NewClient("https://www.voetbalpoules.nl/")

	client.Pool.GetDeelnemers(180001, "eredivisie")
}
