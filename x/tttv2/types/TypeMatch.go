package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Match struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    Opponent sdk.AccAddress `json:"opponent" yaml:"opponent"`
	Status string `json:"status" yaml:"status"`
	Board [9] string `json:"board" yaml:"board"`
	NextPlayer sdk.AccAddress `json:"nextplayer" yaml:"nextplayer"`
	RoleCreator string `json:"rolecreator" yaml:"rolecreator"`
	RoleOpponent string `json:"roleopponent" yaml:"roleopponent"`
	Winner sdk.AccAddress `json:"winner" yaml:"winner"`
}
