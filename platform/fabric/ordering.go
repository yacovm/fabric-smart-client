/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package fabric

import (
	"github.com/hyperledger-labs/fabric-smart-client/platform/fabric/driver"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/grpc"
)

type Ordering struct {
	network driver.FabricNetworkService
}

// Orderers returns the list of known Orderer nodes
func (n *Ordering) Orderers() []*grpc.ConnectionConfig {
	return n.network.Orderers()
}

func (n *Ordering) Broadcast(blob interface{}) error {
	switch b := blob.(type) {
	case *Envelope:
		return n.network.Broadcast(b.e)
	case *Transaction:
		return n.network.Broadcast(b.tx)
	default:
		return n.network.Broadcast(blob)
	}
}
