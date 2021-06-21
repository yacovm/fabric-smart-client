/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package start

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/hyperledger-labs/fabric-smart-client/integration"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/fabric"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/fsc"
)

type Topology struct {
	Name string `yaml:"name,omitempty"`
}

type Topologies struct {
	Topologies []Topology `yaml:"topologies,omitempty"`
}

type T struct {
	Topologies []interface{} `yaml:"topologies,omitempty"`
}

var path string

// Cmd returns the Cobra Command for Version
func Cmd() *cobra.Command {
	// Set the flags on the node start command.
	flags := cobraCommand.Flags()
	flags.StringVarP(&path, "path", "p", "", "artifacts path")

	return cobraCommand
}

var cobraCommand = &cobra.Command{
	Use:   "start",
	Short: "Start networks.",
	Long:  `Read topology from artifacts path and start networds.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}
		// Parsing of the command line is done so silence cmd usage
		cmd.SilenceUsage = true
		return start(args)
	},
}

// start read topology and generates artifacts
func start(args []string) error {
	if len(path) == 0 {
		return errors.Errorf("expecting artifacts path")
	}
	raw, err := ioutil.ReadFile(filepath.Join(path, "topologies.yaml"))
	if err != nil {
		return errors.Wrapf(err, "failed reading topology file [%s]", path)
	}
	names := &Topologies{}
	if err := yaml.Unmarshal(raw, names); err != nil {
		return errors.Wrapf(err, "failed unmarshalling topology file [%s]", path)
	}

	t := &T{}
	if err := yaml.Unmarshal(raw, t); err != nil {
		return errors.Wrapf(err, "failed unmarshalling topology file [%s]", path)
	}
	t2 := []nwo.Topology{}
	for i, topology := range names.Topologies {
		switch topology.Name {
		case fabric.TopologyName:
			top := fabric.NewDefaultTopology()
			r, err := yaml.Marshal(t.Topologies[i])
			if err != nil {
				return errors.Wrapf(err, "failed remarshalling topology configuration [%s]", path)
			}
			if err := yaml.Unmarshal(r, top); err != nil {
				return errors.Wrapf(err, "failed unmarshalling topology file [%s]", path)
			}
			t2 = append(t2, top)
		case fsc.TopologyName:
			top := fsc.NewTopology()
			r, err := yaml.Marshal(t.Topologies[i])
			if err != nil {
				return errors.Wrapf(err, "failed remarshalling topology configuration [%s]", path)
			}
			if err := yaml.Unmarshal(r, top); err != nil {
				return errors.Wrapf(err, "failed unmarshalling topology file [%s]", path)
			}
			t2 = append(t2, top)
		}
	}

	ii, err := integration.Load(path, t2...)
	if err != nil {
		return errors.Wrapf(err, "failed starting networks from [%s]", path)
	}
	ii.Start()
	time.Sleep(10 * time.Second)

	return nil
}
