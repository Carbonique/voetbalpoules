package main

import (
	"fmt"
	"time"

	"github.com/Carbonique/voetbalpoules-messenger/voetbalpoules"
)

func main() {

	client := voetbalpoules.NewClient("https://www.voetbalpoules.nl/")

	t1 := time.Now().Local()
	t2 := t1.AddDate(0, 0, 1)
	t := client.Wedstrijden.Get("ek_vrouwen_2022", t1, t2)
	fmt.Println(t)
}
