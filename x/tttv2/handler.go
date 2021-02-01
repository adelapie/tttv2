package tttv2

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/adelapie/tttv2/x/tttv2/keeper"
	"github.com/adelapie/tttv2/x/tttv2/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreateMove:
			return handleMsgCreateMove(ctx, k, msg)
		case types.MsgSetMove:
			return handleMsgSetMove(ctx, k, msg)
		case types.MsgDeleteMove:
			return handleMsgDeleteMove(ctx, k, msg)
		case types.MsgCreateMatch:
			return handleMsgCreateMatch(ctx, k, msg)
		case types.MsgSetMatch:
			return handleMsgSetMatch(ctx, k, msg)
		case types.MsgDeleteMatch:
			return handleMsgDeleteMatch(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
