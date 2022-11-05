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
		vw, _ := client.GetPoolVoorspelling(t1, t2, 18173, "eredivisie")
		for d, v := range vw[0].DeelnemerVoorspellingen {
			if v.DoelpuntenThuis != nil && v.DoelpuntenUit != nil {

				fmt.Printf("%s: %d - %d\n", d.Naam, *v.DoelpuntenThuis, *v.DoelpuntenUit)
			}
		}
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
