package tttv2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/adelapie/tttv2/x/tttv2/types"
	"github.com/adelapie/tttv2/x/tttv2/keeper"
)

// Handle a message to delete name
func handleMsgDeleteMove(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteMove) (*sdk.Result, error) {
	if !k.MoveExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetMoveOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteMove(ctx, msg.ID)
	return &sdk.Result{}, nil
}
