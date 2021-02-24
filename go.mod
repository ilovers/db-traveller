module github.com/okex/db-traveller

go 1.14

require (
	github.com/cosmos/cosmos-sdk v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.1.3
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.33.9
	github.com/tendermint/tm-db v0.6.4
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/okex/cosmos-sdk v0.39.2-okexchain4
	github.com/tendermint/iavl => github.com/okex/iavl v0.14.1-okexchain1
	github.com/tendermint/tendermint => github.com/okex/tendermint v0.33.9-okexchain2
)
