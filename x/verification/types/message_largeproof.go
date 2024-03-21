package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLargeproof{}

func NewMsgLargeproof(creator string, index uint64, proof string, proofHash string, finished bool) *MsgLargeproof {
	return &MsgLargeproof{
		Creator:   creator,
		Index:     index,
		Proof:     proof,
		ProofHash: proofHash,
		Finished:  finished,
	}
}

func (msg *MsgLargeproof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
