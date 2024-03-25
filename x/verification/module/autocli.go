package verification

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "alignedlayer/api/alignedlayer/verification"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "VerifyPlonk",
					Use:            "verify-plonk [proof] [public_inputs] [verifying_key]",
					Short:          "Send a verify-plonk tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof"}, {ProtoField: "public_inputs"}, {ProtoField: "verifying_key"}},
				},
				{
					RpcMethod:      "VerifyCairo",
					Use:            "verify-cairo [proof]",
					Short:          "Send a verify-cairo tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
