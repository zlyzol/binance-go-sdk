package types

import (
	"github.com/tendermint/go-amino"
	ntypes "github.com/zlyzol/binance-go-sdk/common/types"
	"github.com/zlyzol/binance-go-sdk/types/tx"
	types "github.com/zlyzol/tendermint-0.32.3/rpc/core/types"
)

func NewCodec() *amino.Codec {
	cdc := amino.NewCodec()
	types.RegisterAmino(cdc)
	ntypes.RegisterWire(cdc)
	tx.RegisterCodec(cdc)
	return cdc
}
