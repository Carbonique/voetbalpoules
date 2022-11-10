/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"time"

	voetbalpoulestelegram "github.com/Carbonique/voetbalpoules/pkg/telegram"
	voetbalpoules "github.com/Carbonique/voetbalpoules/pkg/voetbalpoules-client"
	"github.com/spf13/cobra"
)

var timeVoorspelling int

// voorspellingenCmd represents the voorspellingen command
var voorspellingCmd = &cobra.Command{
	Use:   "voorspelling",
	Short: "sends the voorspelling for all wedstrijden within a timerange",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := voetbalpoules.NewClient(BASE_URL)
		bot := voetbalpoulestelegram.NewBot(TOKEN, CHAT)
		t1 := time.Now()
		t2 := t1.Add(time.Minute * time.Duration(timeVoorspelling))
		vw, _ := client.GetPoolVoorspelling(t1, t2, POOL_ID, COMPETITIE)

		for _, vw2 := range vw {
			switch {
			case t1.Before(t2):
				bot.StuurVoorspellingVoorlopig(vw2, BASE_URL)

			case t1.After(t2):
				bot.StuurVoorspellingDefinitief(vw2, BASE_URL)

			case t1.Equal(t2):
				bot.StuurVoorspellingDefinitief(vw2, BASE_URL)

			}

		}
	},
}

func init() {
	rootCmd.AddCommand(voorspellingCmd)
	voorspellingCmd.PersistentFlags().IntVar(&timeVoorspelling, "time", 45, "Time in minutes from now to look for voorspellingen (counting from current time)")

}
