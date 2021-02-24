package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/types"

	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/tendermint/go-amino"
	dbm "github.com/tendermint/tm-db"
	"log"
)

var cdc = codec.New()

var queryTxCmd = &cobra.Command{
	Use:   "txs [hash]",
	Short: "query tx by hash.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		txHash := args[0]

		hash, err := hex.DecodeString(txHash)
		validateError(err)

		dbId := "tx_index"
		db, err := dbm.NewDB(dbId, dbm.GoLevelDBBackend, dataPath)
		validateError(err)

		rawBytes, err := db.Get(hash)
		validateError(err)

		cdc := amino.NewCodec()
		txResult := new(types.TxResult)
		err = cdc.UnmarshalBinaryBare(rawBytes, &txResult)
		validateError(err)

		fmt.Println(txResult.Tx.String())

		result, err := json.MarshalIndent(txResult, " ", "")
		validateError(err)
		fmt.Printf(string(result))
		return nil
	},
}

func validateError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
