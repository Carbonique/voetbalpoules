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

	viper.AutomaticEnv()
	var base_url string
	if viper.GetString("BASE_URL") == "" {
		base_url = "https://voetbalpoules.nl"
	} else {
		base_url = viper.GetString("BASE_URL")
	}
	token := viper.GetString("TOKEN")
	chat := viper.GetInt64("CHAT")
	competitie := viper.GetString("COMPETITIE")
	pool_id := viper.GetInt("POOL_ID")

	rootCmd.PersistentFlags().StringVar(&BASE_URL, "BASE_URL", base_url, "Base url")
	rootCmd.PersistentFlags().StringVar(&TOKEN, "TOKEN", token, "Telegram token")
	rootCmd.PersistentFlags().Int64Var(&CHAT, "CHAT", chat, "Telegram chat")
	rootCmd.PersistentFlags().StringVar(&COMPETITIE, "COMPETITIE", competitie, "voetbalpoules competitie name")
	rootCmd.PersistentFlags().IntVar(&POOL_ID, "POOL_ID", pool_id, "voetbalpoules pool id")

}
