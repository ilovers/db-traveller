package cmd

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/store"
	"github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"strconv"
)

var blockTxsCmd = &cobra.Command{
	Use:   "block-txs [height]",
	Short: "query txs in block by height.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		blockHeight := args[0]

		height, err := strconv.ParseInt(blockHeight, 10, 64)
		if err != nil {
			return err
		}

		blockStoreID := "blockstore"
		db, err := dbm.NewDB(blockStoreID, dbm.GoLevelDBBackend, dataPath)
		blockStore := store.NewBlockStore(db)
		block := blockStore.LoadBlock(height)
		txs := block.Txs
		txLen := len(txs)
		txHashes := make([]string, txLen)
		for i, txBytes := range txs {
			txHashes[i] = fmt.Sprintf("%X", tmhash.Sum(txBytes))
		}

		cdc := amino.NewCodec()
		dbId := "tx_index"
		txDB, err := dbm.NewDB(dbId, dbm.GoLevelDBBackend, dataPath)
		validateError(err)
		var txResultList []*types.TxResult
		for _, txHash := range txHashes {
			hash, _ := hex.DecodeString(txHash)

			rawBytes, err := txDB.Get(hash)
			validateError(err)

			txResult := new(types.TxResult)
			err = cdc.UnmarshalBinaryBare(rawBytes, &txResult)
			validateError(err)
			txResultList = append(txResultList, txResult)
		}

		result, err := json.MarshalIndent(txResultList, " ", "")
		validateError(err)
		fmt.Printf(string(result))
		return nil

		return nil
	},
}
