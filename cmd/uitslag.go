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

var timeUitslag int

// uitslagenCmd represents the uitslagen command
var uitslagCmd = &cobra.Command{
	Use:   "uitslag",
	Short: "sends the uitslag and the voorspellingen for all wedstrijden within a timerange.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := voetbalpoules.NewClient(BASE_URL)
		bot := voetbalpoulestelegram.NewBot(TOKEN, CHAT)
		t1 := time.Now()
		t2 := t1.Add(time.Minute * time.Duration(timeUitslag))
		vw, _ := client.GetPoolVoorspelling(t1, t2, POOL_ID, COMPETITIE)

		for _, vw2 := range vw {
			bot.StuurUitslag(vw2)
		}
	},
}

func init() {
	rootCmd.AddCommand(uitslagCmd)
	uitslagCmd.PersistentFlags().IntVar(&timeUitslag, "time", -120, "Time in mimutes from now to look for uitslagen (counting from current time)")

}
