/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	voetbalpoulestelegram "github.com/Carbonique/voetbalpoules/pkg/telegram"
	voetbalpoules "github.com/Carbonique/voetbalpoules/pkg/voetbalpoules-client"
	"github.com/spf13/cobra"
)

var timeStand int

// standCmd represents the stand command
var standCmd = &cobra.Command{
	Use:   "stand",
	Short: "sends the stand for a pool",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		client := voetbalpoules.NewClient(BASE_URL)
		bot := voetbalpoulestelegram.NewBot(TOKEN, CHAT)

		// If flag is changed the user wants to only send stand if a match
		// has started in the timerange
		if cmd.Flags().Lookup("time").Changed {
			t1 := time.Now()
			t2 := t1.Add(time.Minute * time.Duration(timeStand))

			w, _ := client.GetWedstrijden(t1, t2, POOL_ID, COMPETITIE)
			fmt.Printf("w: %v\n", w)

			if len(w) == 0 {
				return
			}

		}

		deelnemers := client.GetStand(POOL_ID, COMPETITIE)
		bot.StuurStand(deelnemers)
	},
}

func init() {
	rootCmd.AddCommand(standCmd)
	standCmd.PersistentFlags().IntVar(&timeStand, "time", 0, "Time in minutes from now to look for matches. Will only print stand if a match started within this timeframe")

}
