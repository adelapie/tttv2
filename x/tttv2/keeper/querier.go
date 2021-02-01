package keeper

import (
  // this line is used by starport scaffolding # 1
	"github.com/adelapie/tttv2/x/tttv2/types"
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for tttv2 clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListMove:
			return listMove(ctx, k)
		case types.QueryGetMove:
			return getMove(ctx, path[1:], k)
		case types.QueryListMatch:
			return listMatch(ctx, k)
		case types.QueryGetMatch:
			return getMatch(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown tttv2 query endpoint")
		}
	}
}
