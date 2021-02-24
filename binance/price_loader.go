package binance

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"price_loader/storage"
	"time"
)

const (
	updateTimeout = 5 * time.Second
	serverHost    = "api.binance.com"
	serverScheme  = "https"
	serverPath    = "api/v3/ticker/price"
)

// SymbolPriceBinance - binance price data format
type SymbolPriceBinance struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// PriceLoader runs ticker data loader and puts symbols price data into the data channel
func PriceLoader(dataChannel chan storage.SymbolPrice, errorChannel chan error) {
	done := make(chan struct{})
	u := url.URL{Scheme: serverScheme, Host: serverHost, Path: serverPath}

	go func() {
		defer close(done)
		for {
			resp, err := http.Get(u.String())
			if err != nil {
				errorChannel <- err
				done <- struct{}{}
				continue
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				errorChannel <- err
				done <- struct{}{}
				continue
			}

			data := []SymbolPriceBinance{}

			err = json.Unmarshal(body, &data)
			if err != nil {
				errorChannel <- err
				done <- struct{}{}
				continue
			}

			for _, item := range data {
				preparedItem, err := Transform(item)
				if err != nil {
					errorChannel <- err
					done <- struct{}{}
					continue
				}

				dataChannel <- preparedItem
			}
			time.Sleep(updateTimeout)
		}
	}()
}
