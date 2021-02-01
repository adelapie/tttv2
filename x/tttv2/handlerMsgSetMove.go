package tttv2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/adelapie/tttv2/x/tttv2/types"
	"github.com/adelapie/tttv2/x/tttv2/keeper"
)

func handleMsgSetMove(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetMove) (*sdk.Result, error) {
	var move = types.Move{
		Creator: msg.Creator,
		ID:      msg.ID,
    	MatchID: msg.MatchID,
    	Position: msg.Position,
	}
	if !msg.Creator.Equals(k.GetMoveOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetMove(ctx, move)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
