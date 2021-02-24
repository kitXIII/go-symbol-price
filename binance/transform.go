package binance

import (
	"price_loader/storage"
	"strconv"
)

// Transform from data type TickerDataMessage to data type SymbolPrice
func Transform(m SymbolPriceBinance) (storage.SymbolPrice, error) {
	price, err := strconv.ParseFloat(m.Price, 64)

	if err != nil {
		return storage.SymbolPrice{}, err
	}

	return storage.SymbolPrice{Symbol: m.Symbol, Price: price}, nil
}
