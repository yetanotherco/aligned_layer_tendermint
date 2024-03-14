package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVerify{}

func NewMsgVerify(creator string, proof string, publicInputs string, verifyingKey string) *MsgVerify {
	return &MsgVerify{
		Creator:      creator,
		Proof:        proof,
		PublicInputs: publicInputs,
		VerifyingKey: verifyingKey,
	}
}

func (msg *MsgVerify) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
