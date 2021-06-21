/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	gen2 "github.com/hyperledger-labs/fabric-smart-client/integration/nwo/cmd/artifacts/gen"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/cmd/artifacts/start"
	version2 "github.com/hyperledger-labs/fabric-smart-client/integration/nwo/cmd/artifacts/version"
)

const CmdRoot = "core"

// The main command describes the service and
// defaults to printing the help message.
var mainCmd = &cobra.Command{Use: "artifacts"}

func main() {
	// For environment variables.
	viper.SetEnvPrefix(CmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Define command-line flags that are valid for all peer commands and
	// subcommands.
	mainFlags := mainCmd.PersistentFlags()

	mainFlags.String("logging-level", "", "Legacy logging level flag")
	viper.BindPFlag("logging_level", mainFlags.Lookup("logging-level"))
	mainFlags.MarkHidden("logging-level")

	mainCmd.AddCommand(gen2.Cmd())
	mainCmd.AddCommand(start.Cmd())
	mainCmd.AddCommand(version2.Cmd())

	// On failure Cobra prints the usage message and error string, so we only
	// need to exit with a non-0 status
	if mainCmd.Execute() != nil {
		os.Exit(1)
	}
}
