package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetMove{}

type MsgSetMove struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  MatchID string `json:"matchID" yaml:"matchID"`
  Position string `json:"position" yaml:"position"`
}

func NewMsgSetMove(creator sdk.AccAddress, id string, matchID string, position string) MsgSetMove {
  return MsgSetMove{
    ID: id,
		Creator: creator,
    MatchID: matchID,
    Position: position,
	}
}

func (msg MsgSetMove) Route() string {
  return RouterKey
}

func (msg MsgSetMove) Type() string {
  return "SetMove"
}

func (msg MsgSetMove) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetMove) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetMove) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}