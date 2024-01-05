package cmd

import (
	"github.com/spf13/cobra"
)

var profile string

var rootCmd = &cobra.Command{
	Use:   "quotes",
	Short: "API to interact with a collection of quotes",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "development", "application profile (default is development)")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
