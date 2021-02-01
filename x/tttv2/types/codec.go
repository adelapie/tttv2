package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding # 1
		cdc.RegisterConcrete(MsgCreateMove{}, "tttv2/CreateMove", nil)
		cdc.RegisterConcrete(MsgSetMove{}, "tttv2/SetMove", nil)
		cdc.RegisterConcrete(MsgDeleteMove{}, "tttv2/DeleteMove", nil)
		cdc.RegisterConcrete(MsgCreateMatch{}, "tttv2/CreateMatch", nil)
		cdc.RegisterConcrete(MsgSetMatch{}, "tttv2/SetMatch", nil)
		cdc.RegisterConcrete(MsgDeleteMatch{}, "tttv2/DeleteMatch", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
