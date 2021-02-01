package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMatch{}

type MsgCreateMatch struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Opponent string `json:"opponent" yaml:"opponent"`
}

func NewMsgCreateMatch(creator sdk.AccAddress, opponent string) MsgCreateMatch {
  return MsgCreateMatch{
		Creator: creator,
    Opponent: opponent,
	}
}

func (msg MsgCreateMatch) Route() string {
  return RouterKey
}

func (msg MsgCreateMatch) Type() string {
  return "CreateMatch"
}

func (msg MsgCreateMatch) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateMatch) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateMatch) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}