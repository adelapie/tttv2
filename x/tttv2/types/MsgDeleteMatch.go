package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteMatch{}

type MsgDeleteMatch struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteMatch(id string, creator sdk.AccAddress) MsgDeleteMatch {
  return MsgDeleteMatch{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteMatch) Route() string {
  return RouterKey
}

func (msg MsgDeleteMatch) Type() string {
  return "DeleteMatch"
}

func (msg MsgDeleteMatch) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteMatch) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteMatch) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}