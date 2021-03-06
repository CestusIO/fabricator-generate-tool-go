// Code generated by fabricator-generate-plugin-go
//
// Modifications in code regions will be lost during regeneration!

package main

import (
	// region CODE_REGION(import)
	"context"
	"os"

	_ "code.cestus.io/tools/fabricator-generate-tool-go"
	fabricatorgeneratetoolgo "code.cestus.io/tools/fabricator-generate-tool-go/pkg/fabricator-generate-tool-go"
	"code.cestus.io/tools/fabricator/pkg/cmd/version"
	"code.cestus.io/tools/fabricator/pkg/fabricator"
	"code.cestus.io/tools/fabricator/pkg/helpers"
	// endregion
)

func main() {
	// region CODE_REGION(Main)
	ctx := context.Background()
	io := fabricator.NewStdIOStreams()
	ctx, cancel := helpers.WithCancelOnSignal(ctx, io, fabricator.TerminationSignals...)
	defer cancel()
	root := fabricatorgeneratetoolgo.NewFabricatorGenerateToolGo(io, helpers.DefaultFlagParser)
	root.AddCommand(version.NewCmdVersion(io))
	if err := root.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
	// endregion
}
