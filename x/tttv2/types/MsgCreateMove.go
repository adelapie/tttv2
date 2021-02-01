package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
    "strconv"
)

var _ sdk.Msg = &MsgCreateMove{}

type MsgCreateMove struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  MatchID string `json:"matchID" yaml:"matchID"`
  Position string `json:"position" yaml:"position"`
}

func NewMsgCreateMove(creator sdk.AccAddress, matchID string, position string) MsgCreateMove {
  return MsgCreateMove{
		Creator: creator,
    MatchID: matchID,
    Position: position,
	}
}

func (msg MsgCreateMove) Route() string {
  return RouterKey
}

func (msg MsgCreateMove) Type() string {
  return "CreateMove"
}

func (msg MsgCreateMove) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateMove) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateMove) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }

  i, _ := strconv.Atoi(msg.Position)

  if i < 0 || i > 8 {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }



  return nil
}
