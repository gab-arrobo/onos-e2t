// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/encoder"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-ies"
	"gotest.tools/assert"
)

func TestResetRequest(t *testing.T) {
	newE2apPdu, err := CreateResetRequestE2apPdu(1, &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_Protocol{
			Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
		},
	})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("ResetRequest E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	e2apPdu, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	//per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	//assert.NilError(t, err)
	//t.Logf("ResetRequest E2AP PDU PER\n%v", hex.Dump(per))
	//
	//e2apPdu, err = asn1cgo.PerDecodeE2apPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())
}