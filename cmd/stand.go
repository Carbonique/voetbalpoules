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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := voetbalpoules.NewClient("https://www.voetbalpoules.nl/")
		bot := voetbalpoulestelegram.NewBot()
		deelnemers := client.Pool.GetDeelnemers(18173, "eredivisie")
		bot.StuurStand(deelnemers)
	},
}

func init() {
	rootCmd.AddCommand(standCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// standCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// standCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
