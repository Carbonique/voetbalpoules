/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	voetbalpoulestelegram "github.com/Carbonique/voetbalpoules/pkg/telegram"
	voetbalpoules "github.com/Carbonique/voetbalpoules/pkg/voetbalpoules-client"
	"github.com/spf13/cobra"
)

// standCmd represents the stand command
var standCmd = &cobra.Command{
	Use:   "stand",
	Short: "sends the stand for a pool",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := voetbalpoules.NewClient(BASE_URL)
		bot := voetbalpoulestelegram.NewBot(TOKEN, CHAT)
		deelnemers := client.GetStand(POOL_ID, COMPETITIE)
		bot.StuurStand(deelnemers)
	},
}

func init() {
	rootCmd.AddCommand(standCmd)

}
