/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var BASE_URL string
var TOKEN string
var CHAT int64
var COMPETITIE string
var POOL_ID int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "A client to send voetbalpoules pool info to Telegram",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	rootCmd.PersistentFlags().StringVar(&BASE_URL, "BASE_URL", viper.GetString("BASE_URL"), "Base url")
	rootCmd.PersistentFlags().StringVar(&TOKEN, "TOKEN", viper.GetString("TOKEN"), "Telegram token")
	rootCmd.PersistentFlags().Int64Var(&CHAT, "CHAT", viper.GetInt64("CHAT"), "Telegram chat")
	rootCmd.PersistentFlags().StringVar(&COMPETITIE, "COMPETITIE", viper.GetString("COMPETITIE"), "voetbalpoules competitie name")
	rootCmd.PersistentFlags().IntVar(&POOL_ID, "POOL_ID", viper.GetInt("POOL_ID"), "voetbalpoules pool id")

}
