package types

import (
	"testing"

	"alignedlayer/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgGnarkPlonk_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgGnarkPlonk
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgGnarkPlonk{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgGnarkPlonk{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
