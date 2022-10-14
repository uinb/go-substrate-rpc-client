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

package types

import "github.com/uinb/go-substrate-rpc-client/v4/scale"

type LotteryCallIndex struct {
	PalletIndex U8
	CallIndex   U8
}

func (m *LotteryCallIndex) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.PalletIndex)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.CallIndex)
	if err != nil {
		return err
	}

	return nil
}

func (m LotteryCallIndex) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(m.PalletIndex)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.CallIndex)
	if err != nil {
		return err
	}

	return nil
}
