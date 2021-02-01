package tttv2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/adelapie/tttv2/x/tttv2/types"
	"github.com/adelapie/tttv2/x/tttv2/keeper"
)

func handleMsgSetMatch(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetMatch) (*sdk.Result, error) {
      op_tmp, _ := sdk.AccAddressFromBech32(msg.Opponent)                         



	var match = types.Match{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Opponent: op_tmp, //msg.Opponent,
	}
	if !msg.Creator.Equals(k.GetMatchOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetMatch(ctx, match)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
