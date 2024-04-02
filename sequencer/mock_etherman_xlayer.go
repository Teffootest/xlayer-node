// Code generated by mockery v2.39.0. DO NOT EDIT.

package sequencer

import (
	common "github.com/ethereum/go-ethereum/common"

	coretypes "github.com/ethereum/go-ethereum/core/types"

	types "github.com/0xPolygonHermez/zkevm-node/etherman/types"
)


// BuildSequenceBatchesTxData provides a mock function with given fields: sender, sequences, l2CoinBase
func (_m *EthermanMock) BuildSequenceBatchesTxDataXLayer(sender common.Address, sequences []types.Sequence, l2CoinBase common.Address, committeeSignaturesAndAddrs []byte) (*common.Address, []byte, error) {
	ret := _m.Called(sender, sequences, l2CoinBase)

	if len(ret) == 0 {
		panic("no return value specified for BuildSequenceBatchesTxData")
	}

	var r0 *common.Address
	var r1 []byte
	var r2 error
	if rf, ok := ret.Get(0).(func(common.Address, []types.Sequence, common.Address) (*common.Address, []byte, error)); ok {
		return rf(sender, sequences, l2CoinBase)
	}
	if rf, ok := ret.Get(0).(func(common.Address, []types.Sequence, common.Address) *common.Address); ok {
		r0 = rf(sender, sequences, l2CoinBase)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, []types.Sequence, common.Address) []byte); ok {
		r1 = rf(sender, sequences, l2CoinBase)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	if rf, ok := ret.Get(2).(func(common.Address, []types.Sequence, common.Address) error); ok {
		r2 = rf(sender, sequences, l2CoinBase)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EstimateGasSequenceBatchesXLayer provides a mock function with given fields: sender, sequences, l2CoinBase
func (_m *EthermanMock) EstimateGasSequenceBatchesXLayer(sender common.Address, sequences []types.Sequence, l2CoinBase common.Address, committeeSignaturesAndAddrs []byte) (*coretypes.Transaction, error) {
	ret := _m.Called(sender, sequences, l2CoinBase)

	if len(ret) == 0 {
		panic("no return value specified for EstimateGasSequenceBatches")
	}

	var r0 *coretypes.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, []types.Sequence, common.Address) (*coretypes.Transaction, error)); ok {
		return rf(sender, sequences, l2CoinBase)
	}
	if rf, ok := ret.Get(0).(func(common.Address, []types.Sequence, common.Address) *coretypes.Transaction); ok {
		r0 = rf(sender, sequences, l2CoinBase)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, []types.Sequence, common.Address) error); ok {
		r1 = rf(sender, sequences, l2CoinBase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
