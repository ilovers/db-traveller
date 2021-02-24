package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	dataPath string
)

var rootCmd = &cobra.Command{
	Use:   "db-traveller [command]",
	Short: "db-traveller is a tool for analysis of okexchain db.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dataPath, "data-path", "d", "", "data path of okexchain db")
	rootCmd.AddCommand(queryCmd(), blockTxsCmd)
}
