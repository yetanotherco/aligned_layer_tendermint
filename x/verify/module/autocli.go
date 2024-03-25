package verify

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "alignedlayer/api/alignedlayer/verify"
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
					RpcMethod:      "GnarkPlonk",
					Use:            "gnark-plonk [proof] [public-inputs] [verifying-key]",
					Short:          "Send a gnark-plonk tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof"}, {ProtoField: "publicInputs"}, {ProtoField: "verifyingKey"}},
				},
				{
					RpcMethod:      "CairoPlatinum",
					Use:            "cairo-platinum [proof]",
					Short:          "Send a cairo-platinum tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof"}},
				},
				{
					RpcMethod:      "Kimchi",
					Use:            "kimchi [proof]",
					Short:          "Send a kimchi tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
