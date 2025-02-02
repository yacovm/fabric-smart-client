/*
Copyright IBM Corp All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fsc_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hyperledger-labs/fabric-smart-client/integration"
	"github.com/hyperledger-labs/fabric-smart-client/integration/fabric/atsa/fsc"
	"github.com/hyperledger-labs/fabric-smart-client/integration/fabric/atsa/fsc/views"
)

var _ = Describe("EndToEnd", func() {
	var (
		ii *integration.Infrastructure
	)

	AfterEach(func() {
		// Stop the ii
		ii.Stop()
	})

	Describe("Asset Transfer Secured Agreement (With Approvers)", func() {
		var (
			issuer *fsc.Client
			alice  *fsc.Client
			bob    *fsc.Client
		)

		BeforeEach(func() {
			var err error
			// Create the integration ii
			ii, err = integration.Generate(StartPort(), fsc.Topology()...)
			Expect(err).NotTo(HaveOccurred())
			// Start the integration ii
			ii.Start()

			approver := ii.Identity("approver")

			issuer = fsc.NewClient(ii.Client("issuer"), ii.Identity("issuer"), approver)
			alice = fsc.NewClient(ii.Client("alice"), ii.Identity("alice"), approver)
			bob = fsc.NewClient(ii.Client("bob"), ii.Identity("bob"), approver)
		})

		It("succeeded", func() {
			assetID, err := issuer.Issue(&views.Asset{
				ObjectType:        "coin",
				ID:                "1234",
				Owner:             ii.Identity("alice"),
				PublicDescription: "Coin",
				PrivateProperties: []byte("Hello World!!!"),
			})
			Expect(err).ToNot(HaveOccurred())
			time.Sleep(5 * time.Second)

			agreementID, err := alice.AgreeToSell(&views.AgreementToSell{
				TradeID: "1234",
				ID:      "1234",
				Price:   100,
			})
			Expect(err).ToNot(HaveOccurred())
			time.Sleep(5 * time.Second)

			_, err = bob.AgreeToBuy(&views.AgreementToBuy{
				TradeID: "1234",
				ID:      "1234",
				Price:   100,
			})
			Expect(err).ToNot(HaveOccurred())
			time.Sleep(5 * time.Second)

			err = alice.Transfer(assetID, agreementID, ii.Identity("bob"))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
