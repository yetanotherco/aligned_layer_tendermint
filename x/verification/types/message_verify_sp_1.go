package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVerifySp1{}

func NewMsgVerifySp1(creator string, proof string) *MsgVerifySp1 {
	return &MsgVerifySp1{
		Creator: creator,
		Proof:   proof,
	}
}

func (msg *MsgVerifySp1) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
