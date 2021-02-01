package tttv2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/adelapie/tttv2/x/tttv2/types"
	"github.com/adelapie/tttv2/x/tttv2/keeper"
)

func handleMsgCreateMatch(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateMatch) (*sdk.Result, error) {
	k.CreateMatch(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
