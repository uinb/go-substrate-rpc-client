// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	chain "github.com/uinb/go-substrate-rpc-client/v4/rpc/chain"
	mock "github.com/stretchr/testify/mock"

	types "github.com/uinb/go-substrate-rpc-client/v4/types"
)

// Chain is an autogenerated mock type for the Chain type
type Chain struct {
	mock.Mock
}

// GetBlock provides a mock function with given fields: blockHash
func (_m *Chain) GetBlock(blockHash types.Hash) (*types.SignedBlock, error) {
	ret := _m.Called(blockHash)

	var r0 *types.SignedBlock
	if rf, ok := ret.Get(0).(func(types.Hash) *types.SignedBlock); ok {
		r0 = rf(blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Hash) error); ok {
		r1 = rf(blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHash provides a mock function with given fields: blockNumber
func (_m *Chain) GetBlockHash(blockNumber uint64) (types.Hash, error) {
	ret := _m.Called(blockNumber)

	var r0 types.Hash
	if rf, ok := ret.Get(0).(func(uint64) types.Hash); ok {
		r0 = rf(blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHashLatest provides a mock function with given fields:
func (_m *Chain) GetBlockHashLatest() (types.Hash, error) {
	ret := _m.Called()

	var r0 types.Hash
	if rf, ok := ret.Get(0).(func() types.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockLatest provides a mock function with given fields:
func (_m *Chain) GetBlockLatest() (*types.SignedBlock, error) {
	ret := _m.Called()

	var r0 *types.SignedBlock
	if rf, ok := ret.Get(0).(func() *types.SignedBlock); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFinalizedHead provides a mock function with given fields:
func (_m *Chain) GetFinalizedHead() (types.Hash, error) {
	ret := _m.Called()

	var r0 types.Hash
	if rf, ok := ret.Get(0).(func() types.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHeader provides a mock function with given fields: blockHash
func (_m *Chain) GetHeader(blockHash types.Hash) (*types.Header, error) {
	ret := _m.Called(blockHash)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(types.Hash) *types.Header); ok {
		r0 = rf(blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Hash) error); ok {
		r1 = rf(blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHeaderLatest provides a mock function with given fields:
func (_m *Chain) GetHeaderLatest() (*types.Header, error) {
	ret := _m.Called()

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func() *types.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeFinalizedHeads provides a mock function with given fields:
func (_m *Chain) SubscribeFinalizedHeads() (*chain.FinalizedHeadsSubscription, error) {
	ret := _m.Called()

	var r0 *chain.FinalizedHeadsSubscription
	if rf, ok := ret.Get(0).(func() *chain.FinalizedHeadsSubscription); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chain.FinalizedHeadsSubscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeNewHeads provides a mock function with given fields:
func (_m *Chain) SubscribeNewHeads() (*chain.NewHeadsSubscription, error) {
	ret := _m.Called()

	var r0 *chain.NewHeadsSubscription
	if rf, ok := ret.Get(0).(func() *chain.NewHeadsSubscription); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chain.NewHeadsSubscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewChainT interface {
	mock.TestingT
	Cleanup(func())
}

// NewChain creates a new instance of Chain. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChain(t NewChainT) *Chain {
	mock := &Chain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
