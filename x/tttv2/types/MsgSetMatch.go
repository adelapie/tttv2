package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetMatch{}

type MsgSetMatch struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Opponent string `json:"opponent" yaml:"opponent"`
}

func NewMsgSetMatch(creator sdk.AccAddress, id string, opponent string) MsgSetMatch {
  return MsgSetMatch{
    ID: id,
		Creator: creator,
    Opponent: opponent,
	}
}

func (msg MsgSetMatch) Route() string {
  return RouterKey
}

func (msg MsgSetMatch) Type() string {
  return "SetMatch"
}

func (msg MsgSetMatch) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetMatch) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetMatch) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}