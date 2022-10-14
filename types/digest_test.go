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

	. "github.com/uinb/go-substrate-rpc-client/v4/types"
	. "github.com/uinb/go-substrate-rpc-client/v4/types/codec"
	. "github.com/uinb/go-substrate-rpc-client/v4/types/test_utils"
)

func TestDigest_EncodeDecode(t *testing.T) {
	AssertRoundtrip(t, Digest{testDigestItem1, testDigestItem2})
	AssertRoundTripFuzz[Digest](t, 100, digestItemFuzzOpts...)
}

func TestDigest_JSONMarshalUnmarshal(t *testing.T) {
	d := Digest{testDigestItem1, testDigestItem2}
	AssertJSONRoundTrip(t, &d)
}

func TestDigest_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			Input: Digest{testDigestItem1, testDigestItem2},
			Expected: MustHexDecodeString(
				"0x080004ab020102030000000000000000000000000000000000000000000000000000000000")}, //nolint:lll
	})
}

func TestDigest_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			Input: MustHexDecodeString(
				"0x080004ab020102030000000000000000000000000000000000000000000000000000000000"), //nolint:lll
			Expected: Digest{testDigestItem1, testDigestItem2}},
	})
}
