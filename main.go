package main

import (
	"fmt"
	"time"

	"github.com/Carbonique/voetbalpoules-messenger/voetbalpoules"
)

func main() {

	client := voetbalpoules.NewClient("https://www.voetbalpoules.nl/")

	t1 := time.Now().Local()
	t2 := t1.Add(time.Hour * -7)
	w, err := client.Wedstrijden.Get("eredivisie", t2, t1)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(w)
}
