package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Move struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    MatchID string `json:"matchID" yaml:"matchID"`
    Position string `json:"position" yaml:"position"`
}