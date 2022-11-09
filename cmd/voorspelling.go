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

// voorspellingenCmd represents the voorspellingen command
var voorspellingCmd = &cobra.Command{
	Use:   "voorspelling",
	Short: "sends the voorspelling for all wedstrijden within a timerange",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := voetbalpoules.NewClient(BASE_URL)
		bot := voetbalpoulestelegram.NewBot(TOKEN, CHAT)
		t1 := time.Now()
		t2 := t1.Add(time.Minute * 65)
		vw, _ := client.GetPoolVoorspelling(t1, t2, POOL_ID, COMPETITIE)

		for _, vw2 := range vw {
			bot.StuurVoorspelling(vw2, BASE_URL)
		}
	},
}

func init() {
	rootCmd.AddCommand(voorspellingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// voorspellingenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// voorspellingenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
