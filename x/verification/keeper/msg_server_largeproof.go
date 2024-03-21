package keeper

import (
	"context"
	"strconv"

	"alignedlayer/x/verification/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Largeproof(goCtx context.Context, msg *types.MsgLargeproof) (*types.MsgLargeproofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.storeService.OpenKVStore(ctx).Set([]byte(msg.ProofHash+strconv.Itoa(int(msg.Index))), []byte(msg.Proof))
	if msg.Finished {
		answer := []byte("")
		for i := 0; i <= int(msg.Index); i++ {
			partial, _ := k.storeService.OpenKVStore(ctx).Get([]byte(msg.ProofHash + strconv.Itoa(i)))
			answer = append(answer, partial...)
		}
		event := sdk.NewEvent("message received",
			sdk.NewAttribute("message", string(answer)))

		ctx.EventManager().EmitEvent(event)
	}
	return &types.MsgLargeproofResponse{}, nil
}
