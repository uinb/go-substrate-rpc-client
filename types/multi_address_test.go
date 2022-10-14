// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types_test

import (
	"testing"

	fuzz "github.com/google/gofuzz"

	"github.com/uinb/go-substrate-rpc-client/v4/signature"
	. "github.com/uinb/go-substrate-rpc-client/v4/types"
	"github.com/stretchr/testify/assert"
)

var (
	multiAddressFuzzOpts = []fuzzOpt{
		withFuzzFuncs(func(m *MultiAddress, c fuzz.Continue) {
			switch c.Intn(5) {
			case 0:
				m.IsID = true
				c.Fuzz(&m.AsID)
			case 1:
				m.IsIndex = true
				c.Fuzz(&m.AsIndex)
			case 2:
				m.IsRaw = true
				c.Fuzz(&m.AsRaw)
			case 3:
				m.IsAddress32 = true
				c.Fuzz(&m.AsAddress32)
			case 4:
				m.IsAddress20 = true
				c.Fuzz(&m.AsAddress20)
			}
		}),
	}
)

func TestMultiAddress_EncodeDecode(t *testing.T) {
	assertRoundTripFuzz[MultiAddress](t, 100, multiAddressFuzzOpts...)
	assertDecodeNilData[MultiAddress](t)
}

func TestNewMultiAddressFromAccountID(t *testing.T) {
	assertRoundtrip(t, NewMultiAddressFromAccountID(signature.TestKeyringPairAlice.PublicKey))

	_, err := NewMultiAddressFromHexAccountID("123!")
	assert.Error(t, err)

	addr, err := NewMultiAddressFromHexAccountID(HexEncodeToString(signature.TestKeyringPairAlice.PublicKey))
	assert.NoError(t, err)
	assertRoundtrip(t, addr)
	assertRoundtrip(t, MultiAddress{
		IsIndex: true,
		AsIndex: 100,
	})
	assertRoundtrip(t, MultiAddress{
		IsRaw: true,
		AsRaw: []byte{1, 2, 3},
	})
	assertRoundtrip(t, MultiAddress{
		IsAddress32: true,
		AsAddress32: [32]byte{},
	})
	assertRoundtrip(t, MultiAddress{
		IsAddress20: true,
		AsAddress20: [20]byte{},
	})
}
