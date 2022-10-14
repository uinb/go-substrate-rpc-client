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
	"encoding/json"
	"testing"

	. "github.com/uinb/go-substrate-rpc-client/v4/types"
)

func TestGenerateMMRProofResponse_Unmarshal(t *testing.T) {
	jsonData := map[string]interface{}{"blockHash": "0x52d917b53796b671eb6f9f5497878cd7bacd1e8bafd41004687f67d2938b7b13", "leaf": "0xc50100d007000024724d3186265a4c1da058d532f7e806cb7b6d0c5be6eecf753f35addc53b54601000000000000000300000042b63941ec636f52303b3c33f53349830d8a466e9456d25d22b28f4bb0ad03650000000000000000000000000000000000000000000000000000000000000000", "proof": "0xd007000000000000d90800000000000030ef3c78f3febfafa319c77aadb0801099da03a775898f486392106a15688c0354c71aed1edc46767357f1f6dd8701f7974dd2ae39e6f178f9de1f6b8402683ef78f2048bf973a2f8b852a4b4997d0b8637111641d88a5ae8670260417a1f4ccfc2a5a56bc7517b04d9931c82a24660c8fa86fd33611b6964c3201669537d7851a17f814e37d71841d69412816692a8c3b6615572b799674c9b6879e547b1f6511af7be87ea122cb70f6da8718cb017acf56c0dc8fb0f093ed300dbbef98d632329b93aa08f39a2fd0cf2e51c12f8b0370b9f3e50d3395a20017e3f24991c9b603951ed76df7fad6c19cade303795345a7a96729fa337ae82d74b03f2b039ee5e3b0c34f63d8f8e43a0db754f4c0f42dc84d98a26c51edea931971206398a80f5ebee8c4ab19a7358adc89d2c4258fb61d979517d363ee33086cc5a9e63c7eed27674f1ca1d3585eb77c1c281ff21057dfdf89227dbe9f6a522753d26cfcf2f271e27f93422d112ef763d502970da640b416c1d3a04c9373b10beef676068ceeec"}

	marshalled, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	expected := GenerateMMRProofResponse{BlockHash:H256{0x52, 0xd9, 0x17, 0xb5, 0x37, 0x96, 0xb6, 0x71, 0xeb, 0x6f, 0x9f, 0x54, 0x97, 0x87, 0x8c, 0xd7, 0xba, 0xcd, 0x1e, 0x8b, 0xaf, 0xd4, 0x10, 0x4, 0x68, 0x7f, 0x67, 0xd2, 0x93, 0x8b, 0x7b, 0x13}, Leaf:MMRLeaf{Version:0x0, ParentNumberAndHash:ParentNumberAndHash{ParentNumber:0x7d0, Hash:Hash{0x24, 0x72, 0x4d, 0x31, 0x86, 0x26, 0x5a, 0x4c, 0x1d, 0xa0, 0x58, 0xd5, 0x32, 0xf7, 0xe8, 0x6, 0xcb, 0x7b, 0x6d, 0xc, 0x5b, 0xe6, 0xee, 0xcf, 0x75, 0x3f, 0x35, 0xad, 0xdc, 0x53, 0xb5, 0x46}}, BeefyNextAuthoritySet:BeefyNextAuthoritySet{ID:0x1, Len:0x3, Root:H256{0x42, 0xb6, 0x39, 0x41, 0xec, 0x63, 0x6f, 0x52, 0x30, 0x3b, 0x3c, 0x33, 0xf5, 0x33, 0x49, 0x83, 0xd, 0x8a, 0x46, 0x6e, 0x94, 0x56, 0xd2, 0x5d, 0x22, 0xb2, 0x8f, 0x4b, 0xb0, 0xad, 0x3, 0x65}}, ParachainHeads:H256{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}, Proof:MMRProof{LeafIndex:0x7d0, LeafCount:0x8d9, Items:[]H256{H256{0xef, 0x3c, 0x78, 0xf3, 0xfe, 0xbf, 0xaf, 0xa3, 0x19, 0xc7, 0x7a, 0xad, 0xb0, 0x80, 0x10, 0x99, 0xda, 0x3, 0xa7, 0x75, 0x89, 0x8f, 0x48, 0x63, 0x92, 0x10, 0x6a, 0x15, 0x68, 0x8c, 0x3, 0x54}, H256{0xc7, 0x1a, 0xed, 0x1e, 0xdc, 0x46, 0x76, 0x73, 0x57, 0xf1, 0xf6, 0xdd, 0x87, 0x1, 0xf7, 0x97, 0x4d, 0xd2, 0xae, 0x39, 0xe6, 0xf1, 0x78, 0xf9, 0xde, 0x1f, 0x6b, 0x84, 0x2, 0x68, 0x3e, 0xf7}, H256{0x8f, 0x20, 0x48, 0xbf, 0x97, 0x3a, 0x2f, 0x8b, 0x85, 0x2a, 0x4b, 0x49, 0x97, 0xd0, 0xb8, 0x63, 0x71, 0x11, 0x64, 0x1d, 0x88, 0xa5, 0xae, 0x86, 0x70, 0x26, 0x4, 0x17, 0xa1, 0xf4, 0xcc, 0xfc}, H256{0x2a, 0x5a, 0x56, 0xbc, 0x75, 0x17, 0xb0, 0x4d, 0x99, 0x31, 0xc8, 0x2a, 0x24, 0x66, 0xc, 0x8f, 0xa8, 0x6f, 0xd3, 0x36, 0x11, 0xb6, 0x96, 0x4c, 0x32, 0x1, 0x66, 0x95, 0x37, 0xd7, 0x85, 0x1a}, H256{0x17, 0xf8, 0x14, 0xe3, 0x7d, 0x71, 0x84, 0x1d, 0x69, 0x41, 0x28, 0x16, 0x69, 0x2a, 0x8c, 0x3b, 0x66, 0x15, 0x57, 0x2b, 0x79, 0x96, 0x74, 0xc9, 0xb6, 0x87, 0x9e, 0x54, 0x7b, 0x1f, 0x65, 0x11}, H256{0xaf, 0x7b, 0xe8, 0x7e, 0xa1, 0x22, 0xcb, 0x70, 0xf6, 0xda, 0x87, 0x18, 0xcb, 0x1, 0x7a, 0xcf, 0x56, 0xc0, 0xdc, 0x8f, 0xb0, 0xf0, 0x93, 0xed, 0x30, 0xd, 0xbb, 0xef, 0x98, 0xd6, 0x32, 0x32}, H256{0x9b, 0x93, 0xaa, 0x8, 0xf3, 0x9a, 0x2f, 0xd0, 0xcf, 0x2e, 0x51, 0xc1, 0x2f, 0x8b, 0x3, 0x70, 0xb9, 0xf3, 0xe5, 0xd, 0x33, 0x95, 0xa2, 0x0, 0x17, 0xe3, 0xf2, 0x49, 0x91, 0xc9, 0xb6, 0x3}, H256{0x95, 0x1e, 0xd7, 0x6d, 0xf7, 0xfa, 0xd6, 0xc1, 0x9c, 0xad, 0xe3, 0x3, 0x79, 0x53, 0x45, 0xa7, 0xa9, 0x67, 0x29, 0xfa, 0x33, 0x7a, 0xe8, 0x2d, 0x74, 0xb0, 0x3f, 0x2b, 0x3, 0x9e, 0xe5, 0xe3}, H256{0xb0, 0xc3, 0x4f, 0x63, 0xd8, 0xf8, 0xe4, 0x3a, 0xd, 0xb7, 0x54, 0xf4, 0xc0, 0xf4, 0x2d, 0xc8, 0x4d, 0x98, 0xa2, 0x6c, 0x51, 0xed, 0xea, 0x93, 0x19, 0x71, 0x20, 0x63, 0x98, 0xa8, 0xf, 0x5e}, H256{0xbe, 0xe8, 0xc4, 0xab, 0x19, 0xa7, 0x35, 0x8a, 0xdc, 0x89, 0xd2, 0xc4, 0x25, 0x8f, 0xb6, 0x1d, 0x97, 0x95, 0x17, 0xd3, 0x63, 0xee, 0x33, 0x8, 0x6c, 0xc5, 0xa9, 0xe6, 0x3c, 0x7e, 0xed, 0x27}, H256{0x67, 0x4f, 0x1c, 0xa1, 0xd3, 0x58, 0x5e, 0xb7, 0x7c, 0x1c, 0x28, 0x1f, 0xf2, 0x10, 0x57, 0xdf, 0xdf, 0x89, 0x22, 0x7d, 0xbe, 0x9f, 0x6a, 0x52, 0x27, 0x53, 0xd2, 0x6c, 0xfc, 0xf2, 0xf2, 0x71}, H256{0xe2, 0x7f, 0x93, 0x42, 0x2d, 0x11, 0x2e, 0xf7, 0x63, 0xd5, 0x2, 0x97, 0xd, 0xa6, 0x40, 0xb4, 0x16, 0xc1, 0xd3, 0xa0, 0x4c, 0x93, 0x73, 0xb1, 0xb, 0xee, 0xf6, 0x76, 0x6, 0x8c, 0xee, 0xec}}}}

	var unmarshalled GenerateMMRProofResponse
	json.Unmarshal(marshalled, &unmarshalled)

	assertEqual(t, unmarshalled, expected)
}
