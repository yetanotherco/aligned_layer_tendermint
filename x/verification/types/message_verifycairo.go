package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVerifycairo{}

func NewMsgVerifycairo(creator string, proof string) *MsgVerifycairo {
	return &MsgVerifycairo{
		Creator: creator,
		Proof:   proof,
	}
}

func (msg *MsgVerifycairo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
