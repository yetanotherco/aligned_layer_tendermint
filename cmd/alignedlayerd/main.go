package main

/*
#cgo darwin LDFLAGS: -L./lib -lsp1_verifier
#cgo linux LDFLAGS: -L./lib -lsp1_verifier -ldl -lrt -lm

#include "lib/sp1.h"
*/
import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"alignedlayer/app"
	"alignedlayer/cmd/alignedlayerd/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
