package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCairoPlatinum{}

func NewMsgCairoPlatinum(creator string, proof string) *MsgCairoPlatinum {
	return &MsgCairoPlatinum{
		Creator: creator,
		Proof:   proof,
	}
}

func (msg *MsgCairoPlatinum) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
