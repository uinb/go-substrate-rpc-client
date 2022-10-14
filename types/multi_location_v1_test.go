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

	"github.com/stretchr/testify/assert"

	fuzz "github.com/google/gofuzz"

	. "github.com/uinb/go-substrate-rpc-client/v4/types"
)

var (
	testMultiLocation = MultiLocationV1{
		Parents: 1,
		Interior: JunctionsV1{
			IsHere: true,
		},
	}

	optionMultiLocationV1FuzzOpts = combineFuzzOpts(
		junctionsV1FuzzOpts,
		[]fuzzOpt{
			withFuzzFuncs(func(o *OptionMultiLocationV1, c fuzz.Continue) {
				if c.RandBool() {
					*o = NewOptionMultiLocationV1Empty()
					return
				}

				var m MultiLocationV1

				c.Fuzz(&m)

				*o = NewOptionMultiLocationV1(m)
			}),
		},
	)
)

func TestOptionMultiLocation_EncodeDecode(t *testing.T) {
	assertRoundTripFuzz[OptionMultiLocationV1](t, 1000, optionMultiLocationV1FuzzOpts...)
	assertEncodeEmptyObj[OptionMultiLocationV1](t, 1)
}

func TestOptionMultiLocation_Encode(t *testing.T) {
	assertEncode(t, []encodingAssert{
		{NewOptionMultiLocationV1(testMultiLocation), MustHexDecodeString("0x010100")},
		{NewOptionMultiLocationV1Empty(), MustHexDecodeString("0x00")},
	})
}

func TestOptionMultiLocation_Decode(t *testing.T) {
	assertDecode(t, []decodingAssert{
		{MustHexDecodeString("0x010100"), NewOptionMultiLocationV1(testMultiLocation)},
		{MustHexDecodeString("0x00"), NewOptionMultiLocationV1Empty()},
	})
}

func TestOptionMultiLocation_OptionMethods(t *testing.T) {
	o := NewOptionMultiLocationV1Empty()
	o.SetSome(testMultiLocationV1n1)

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, MultiLocationV1{}, v)
}

var (
	testMultiLocationV1n1 = MultiLocationV1{
		Parents:  4,
		Interior: testJunctionsV1n9,
	}

	// Note that we only use the junctionsV1FuzzOpts here, since no explicit
	// opts are needed for the MultiLocationV1 as the other fields are handled by fuzz.
	multiLocationV1FuzzOpts = junctionsV1FuzzOpts
)

func TestMultiLocationV1_EncodeDecode(t *testing.T) {
	assertRoundTripFuzz[MultiLocationV1](t, 100, multiLocationV1FuzzOpts...)
	assertDecodeNilData[MultiLocationV1](t)
	assertEncodeEmptyObj[MultiLocationV1](t, 1)
}

func TestMultiLocationV1_Encode(t *testing.T) {
	assertEncode(t, []encodingAssert{
		{testMultiLocationV1n1, MustHexDecodeString("0x0408002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807")},
	})
}

func TestMultiLocationV1_Decode(t *testing.T) {
	assertDecode(t, []decodingAssert{
		{MustHexDecodeString("0x0408002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807"), testMultiLocationV1n1},
	})
}

var (
	testVersionMultiLocation1 = VersionedMultiLocation{
		IsV0:            true,
		MultiLocationV0: testMultiLocationV0n9,
	}
	testVersionMultiLocation2 = VersionedMultiLocation{
		IsV1:            true,
		MultiLocationV1: testMultiLocationV1n1,
	}

	versionedMultiLocationFuzzOpts = combineFuzzOpts(
		multiLocationV0FuzzOpts,
		multiLocationV1FuzzOpts,
		[]fuzzOpt{
			withFuzzFuncs(func(v *VersionedMultiLocation, c fuzz.Continue) {
				if c.RandBool() {
					v.IsV0 = true
					c.Fuzz(&v.MultiLocationV0)

					return
				}

				v.IsV1 = true
				c.Fuzz(&v.MultiLocationV1)
			}),
		},
	)
)

func TestVersionedMultiLocation_EncodeDecode(t *testing.T) {
	assertRoundTripFuzz[VersionedMultiLocation](t, 1000, versionedMultiLocationFuzzOpts...)
	assertDecodeNilData[VersionedMultiLocation](t)
	assertEncodeEmptyObj[VersionedMultiLocation](t, 0)
}

func TestVersionedMultiLocation_Encode(t *testing.T) {
	assertEncode(t, []encodingAssert{
		{testVersionMultiLocation1, MustHexDecodeString("0x000800010b00000002000c010203030010000000000000000403000504062a00000000000000000000000000000007080608")},
		{testVersionMultiLocation2, MustHexDecodeString("0x010408002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807")},
	})
}

func TestVersionedMultiLocation_Decode(t *testing.T) {
	assertDecode(t, []decodingAssert{
		{MustHexDecodeString("0x000800010b00000002000c010203030010000000000000000403000504062a00000000000000000000000000000007080608"), testVersionMultiLocation1},
		{MustHexDecodeString("0x010408002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807"), testVersionMultiLocation2},
	})
}
