/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	_ "net/http/pprof"
	"strings"

	node2 "github.com/hyperledger-labs/fabric-smart-client/node"
	"github.com/spf13/viper"
)

const CmdRoot = "core"

func main() {
	// For environment variables.
	viper.SetEnvPrefix(CmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Instantiate node and execute
	node := node2.New()
	node.Execute(nil)
}
