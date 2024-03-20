package vm

import (
	"github.com/titanchain/titanchain/common"
	"github.com/titanchain/titanchain/log"
	"github.com/titanchain/titanchain/params"
	"github.com/titanchain/titanchain/titanx/tradingstate"
)

const TitanXPriceNumberOfBytesReturn = 32

// titanxPrice implements a pre-compile contract to get token price in titanx

type titanxLastPrice struct {
	tradingStateDB *tradingstate.TradingStateDB
}
type titanxEpochPrice struct {
	tradingStateDB *tradingstate.TradingStateDB
}

func (t *titanxLastPrice) RequiredGas(input []byte) uint64 {
	return params.TitanXPriceGas
}

func (t *titanxLastPrice) Run(input []byte) ([]byte, error) {
	// input includes baseTokenAddress, quoteTokenAddress
	if t.tradingStateDB != nil && len(input) == 64 {
		base := common.BytesToAddress(input[12:32]) // 20 bytes from 13-32
		quote := common.BytesToAddress(input[44:])  // 20 bytes from 45-64
		price := t.tradingStateDB.GetLastPrice(tradingstate.GetTradingOrderBookHash(base, quote))
		if price != nil {
			log.Debug("Run GetLastPrice", "base", base.Hex(), "quote", quote.Hex(), "price", price)
			return common.LeftPadBytes(price.Bytes(), TitanXPriceNumberOfBytesReturn), nil
		}
	}
	return common.LeftPadBytes([]byte{}, TitanXPriceNumberOfBytesReturn), nil
}

func (t *titanxLastPrice) SetTradingState(tradingStateDB *tradingstate.TradingStateDB) {
	if tradingStateDB != nil {
		t.tradingStateDB = tradingStateDB.Copy()
	} else {
		t.tradingStateDB = nil
	}
}

func (t *titanxEpochPrice) RequiredGas(input []byte) uint64 {
	return params.TitanXPriceGas
}

func (t *titanxEpochPrice) Run(input []byte) ([]byte, error) {
	// input includes baseTokenAddress, quoteTokenAddress
	if t.tradingStateDB != nil && len(input) == 64 {
		base := common.BytesToAddress(input[12:32]) // 20 bytes from 13-32
		quote := common.BytesToAddress(input[44:])  // 20 bytes from 45-64
		price := t.tradingStateDB.GetMediumPriceBeforeEpoch(tradingstate.GetTradingOrderBookHash(base, quote))
		if price != nil {
			log.Debug("Run GetEpochPrice", "base", base.Hex(), "quote", quote.Hex(), "price", price)
			return common.LeftPadBytes(price.Bytes(), TitanXPriceNumberOfBytesReturn), nil
		}
	}
	return common.LeftPadBytes([]byte{}, TitanXPriceNumberOfBytesReturn), nil
}

func (t *titanxEpochPrice) SetTradingState(tradingStateDB *tradingstate.TradingStateDB) {
	if tradingStateDB != nil {
		t.tradingStateDB = tradingStateDB.Copy()
	} else {
		t.tradingStateDB = nil
	}
}
