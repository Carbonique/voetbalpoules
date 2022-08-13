package main

import (
	"fmt"
	"time"

	"github.com/Carbonique/voetbalpoules-messenger/voetbalpoules"
)

func main() {

	client := voetbalpoules.NewClient("https://www.voetbalpoules.nl/")

	t1 := time.Now().Local()
	t2 := t1.Add(time.Hour * 10)
	t := client.Wedstrijden.Get("eredivisie", t1, t2)
	fmt.Println(t)
}
