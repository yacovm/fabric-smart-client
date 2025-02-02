/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package state

import "github.com/hyperledger-labs/fabric-smart-client/platform/view/view"

type Identities []view.Identity

func (i Identities) Count() int {
	return len(i)
}

func (i Identities) Others(me view.Identity) Identities {
	res := []view.Identity{}
	for _, identity := range i {
		if identity.Equal(me) {
			continue
		}
		res = append(res, identity)
	}
	return res
}

func (i Identities) Match(ids []view.Identity) bool {
	if len(ids) != len(i) {
		return false
	}
	for _, id := range ids {
		found := false
		for _, identity := range i {
			if identity.Equal(id) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func (i Identities) Contain(id view.Identity) bool {
	for _, identity := range i {
		if identity.Equal(id) {
			return true
		}
	}
	return false
}
