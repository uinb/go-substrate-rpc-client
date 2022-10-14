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

	. "github.com/uinb/go-substrate-rpc-client/v4/types"
)

var (
	testMigrationCompute1 = MigrationCompute{
		IsSigned: true,
	}
	testMigrationCompute2 = MigrationCompute{
		IsAuto: true,
	}

	migrationComputeFuzzOpts = []fuzzOpt{
		withFuzzFuncs(func(m *MigrationCompute, c fuzz.Continue) {
			r := c.RandBool()
			m.IsSigned = r
			m.IsAuto = !r
		}),
	}
)

func TestMigrationCompute_EncodeDecode(t *testing.T) {
	assertRoundTripFuzz[MigrationCompute](t, 100, migrationComputeFuzzOpts...)
	assertDecodeNilData[MigrationCompute](t)
	assertEncodeEmptyObj[MigrationCompute](t, 0)
}

func TestMigrationCompute_Encode(t *testing.T) {
	assertEncode(t, []encodingAssert{
		{testMigrationCompute1, MustHexDecodeString("0x00")},
		{testMigrationCompute2, MustHexDecodeString("0x01")},
	})
}

func TestMigrationCompute_Decode(t *testing.T) {
	assertDecode(t, []decodingAssert{
		{MustHexDecodeString("0x00"), testMigrationCompute1},
		{MustHexDecodeString("0x01"), testMigrationCompute2},
	})
}
