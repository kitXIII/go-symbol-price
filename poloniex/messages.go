package poloniex

import (
	"encoding/json"
	"errors"
)

// SubscriptionRequestMessage type
type SubscriptionRequestMessage struct {
	Command string  `json:"command"`
	Channel float64 `json:"channel"`
}

// TickerDataMessage type
type TickerDataMessage struct {
	Channel       int
	SubscriptonID int
	Data          TickerData
}

// TickerData type
type TickerData struct {
	CurrencyPairID                   int
	LastTradePrice                   string
	LowestAsk                        string
	HighestBid                       string
	PercentChangeInLast24Hours       string
	BaseCurrencyVolumeInLast24Hours  string
	QuoteCurrencyVolumeInLast24Hours string
	IsFrozen                         int
	HighestTradePriceInLast24Hours   string
	LowestTradePriceInLast24Hours    string
}

// UnmarshalJSON implements json.Unmarshaler for TickerDataMessage
func (p *TickerDataMessage) UnmarshalJSON(b []byte) error {
	var records []json.RawMessage

	if err := json.Unmarshal(b, &records); err != nil {
		return err
	}

	if len(records) < 1 {
		return errors.New("Empty message")
	}

	switch len(records) {
	case 3:
		var innerRecords []json.RawMessage
		if err := json.Unmarshal(records[2], &innerRecords); err != nil {
			return err
		}

		if err := json.Unmarshal(innerRecords[0], &p.Data.CurrencyPairID); err != nil {
			return err
		}

		if err := json.Unmarshal(innerRecords[1], &p.Data.LastTradePrice); err != nil {
			return err
		}

		if err := json.Unmarshal(innerRecords[2], &p.Data.LowestAsk); err != nil {
			return err
		}

		if err := json.Unmarshal(innerRecords[3], &p.Data.HighestBid); err != nil {
			return err
		}

		if err := json.Unmarshal(innerRecords[4], &p.Data.PercentChangeInLast24Hours); err != nil {
			return err
		}

		if err := json.Unmarshal(innerRecords[5], &p.Data.BaseCurrencyVolumeInLast24Hours); err != nil {
			return err
		}

		if err := json.Unmarshal(innerRecords[6], &p.Data.QuoteCurrencyVolumeInLast24Hours); err != nil {
			return err
		}

		if err := json.Unmarshal(innerRecords[7], &p.Data.IsFrozen); err != nil {
			return err
		}

		if err := json.Unmarshal(innerRecords[8], &p.Data.HighestTradePriceInLast24Hours); err != nil {
			return err
		}

		if err := json.Unmarshal(innerRecords[9], &p.Data.LowestTradePriceInLast24Hours); err != nil {
			return err
		}

		fallthrough
	case 2:
		if err := json.Unmarshal(records[1], &p.SubscriptonID); err != nil {
			return err
		}
		fallthrough
	case 1:
		if err := json.Unmarshal(records[0], &p.Channel); err != nil {
			return err
		}
		return nil
	}

	return nil
}
