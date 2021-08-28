package tx

import (
	"github.com/tendermint/go-amino"
	"github.com/zlyzol/tendermint-0.32.3/crypto/encoding/amino"

	"github.com/zlyzol/binance-go-sdk/types/msg"
)

// cdc global variable
var Cdc = amino.NewCodec()

func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterInterface((*Tx)(nil), nil)
	cdc.RegisterConcrete(StdTx{}, "auth/StdTx", nil)
	msg.RegisterCodec(cdc)
}

func init() {
	cryptoAmino.RegisterAmino(Cdc)
	RegisterCodec(Cdc)
}
