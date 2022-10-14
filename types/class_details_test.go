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
	"math/big"
	"testing"

	. "github.com/uinb/go-substrate-rpc-client/v4/types"
)

var (
	testClassDetails = ClassDetails{
		Owner:             NewAccountID([]byte("acc_id")),
		Issuer:            NewAccountID([]byte("acc_id2")),
		Admin:             NewAccountID([]byte("acc_id3")),
		Freezer:           NewAccountID([]byte("acc_id4")),
		TotalDeposit:      NewU128(*big.NewInt(123)),
		FreeHolding:       true,
		Instances:         4,
		InstanceMetadatas: 5,
		Attributes:        6,
		IsFrozen:          true,
	}
)

func TestClassDetails_EncodeDecode(t *testing.T) {
	assertRoundTripFuzz[ClassDetails](t, 1000)
	assertDecodeNilData[ClassDetails](t)
	assertEncodeEmptyObj[ClassDetails](t, 158)
}

func TestClassDetails_Encode(t *testing.T) {
	assertEncode(t, []encodingAssert{
		{
			testClassDetails,
			MustHexDecodeString("0x6163635f696400000000000000000000000000000000000000000000000000006163635f696432000000000000000000000000000000000000000000000000006163635f696433000000000000000000000000000000000000000000000000006163635f696434000000000000000000000000000000000000000000000000007b0000000000000000000000000000000104000000050000000600000001"),
		},
	})
}

func TestClassDetails_Decode(t *testing.T) {
	assertDecode(t, []decodingAssert{
		{
			MustHexDecodeString("0x6163635f696400000000000000000000000000000000000000000000000000006163635f696432000000000000000000000000000000000000000000000000006163635f696433000000000000000000000000000000000000000000000000006163635f696434000000000000000000000000000000000000000000000000007b0000000000000000000000000000000104000000050000000600000001"),
			testClassDetails,
		},
	})
}
