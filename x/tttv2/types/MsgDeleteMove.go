package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteMove{}

type MsgDeleteMove struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteMove(id string, creator sdk.AccAddress) MsgDeleteMove {
  return MsgDeleteMove{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteMove) Route() string {
  return RouterKey
}

func (msg MsgDeleteMove) Type() string {
  return "DeleteMove"
}

func (msg MsgDeleteMove) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteMove) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteMove) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}