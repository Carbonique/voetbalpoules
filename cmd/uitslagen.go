/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	voetbalpoules "github.com/Carbonique/voetbalpoules/pkg/voetbalpoules-client"
	"github.com/spf13/cobra"
)

// uitslagenCmd represents the uitslagen command
var uitslagenCmd = &cobra.Command{
	Use:   "uitslagen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := voetbalpoules.NewClient("https://www.voetbalpoules.nl/")
		//bot := voetbalpoulestelegram.NewBot()

		t1 := time.Now()
		t2 := t1.Add(time.Hour * 4)
		w, err := client.Wedstrijden.Get("eredivisie", t1, t2)
		if err != nil {
			fmt.Errorf("Error")
		}
		fmt.Printf("w: %v\n", w)

		d := client.Pool.GetDeelnemers(135352, "eredivisie")
		var voorspellingen []voetbalpoules.Voorspelling
		for _, d2 := range d {
			v, _ := client.Voorspellingen.Get(d2.ID, w[0])
			voorspellingen = append(voorspellingen, v)
		}
		bot.StuurUitslag(w[0], voorspellingen)
	},
}

func init() {
	rootCmd.AddCommand(uitslagenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uitslagenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uitslagenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
