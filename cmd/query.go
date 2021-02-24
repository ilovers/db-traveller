package cmd

import (
	"github.com/spf13/cobra"
)

func queryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:   "query [command]",
		Short: "query data from okexchain db.",
	}

	queryCmd.AddCommand(queryTxCmd, blockTxsCmd)
	return queryCmd
}
