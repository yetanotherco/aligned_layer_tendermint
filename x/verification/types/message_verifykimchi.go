package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVerifykimchi{}

func NewMsgVerifykimchi(creator string, proof string) *MsgVerifykimchi {
	return &MsgVerifykimchi{
		Creator: creator,
		Proof:   proof,
	}
}

func (msg *MsgVerifykimchi) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
