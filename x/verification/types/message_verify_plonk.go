package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVerifyPlonk{}

func NewMsgVerifyPlonk(creator, proof, public_inputs, verifying_key string) *MsgVerifyPlonk {
	return &MsgVerifyPlonk{
		Creator:      creator,
		Proof:        proof,
		PublicInputs: public_inputs,
		VerifyingKey: verifying_key,
	}
}

func (msg *MsgVerifyPlonk) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
