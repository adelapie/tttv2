package tttv2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/adelapie/tttv2/x/tttv2/types"
	"github.com/adelapie/tttv2/x/tttv2/keeper"
)

// Handle a message to delete name
func handleMsgDeleteMatch(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteMatch) (*sdk.Result, error) {
	if !k.MatchExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetMatchOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteMatch(ctx, msg.ID)
	return &sdk.Result{}, nil
}
