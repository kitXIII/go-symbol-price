package poloniex

import (
	"fmt"
	"log"
	"net/url"

	"price_loader/storage"

	"github.com/gorilla/websocket"
)

const (
	tickerChannel       = 1002
	serverHost          = "api2.poloniex.com"
	serverScheme        = "wss"
	serverPath          = ""
	subscriptionCommand = "subscribe"
)

// PriceLoader runs ticker data loader and puts symbols price data into the data channel
func PriceLoader(dataChannel *chan storage.SymbolPrice, errorChannel *chan error) {
	done := make(chan struct{})

	u := url.URL{Scheme: serverScheme, Host: serverHost, Path: serverPath}
	log.Printf("connecting to %s\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		*errorChannel <- err
		return
	}

	defer conn.Close()
	defer close(done)

	// Message handler
	go func() {
		defer close(done)
		for {
			data := TickerDataMessage{}
			err := conn.ReadJSON(&data)
			if err != nil {
				*errorChannel <- err
				done <- struct{}{}
			}

			if data.Channel != tickerChannel || data.SubscriptonID != 0 {
				continue
			}

			preparedData, err := Transform(data)

			if err != nil {
				fmt.Printf("Converter error: %s\n", err.Error())
				continue
			}

			*dataChannel <- preparedData
		}
	}()

	// Subscribe to ticker data channel (1002)
	conn.WriteJSON(SubscriptionRequestMessage{Command: subscriptionCommand, Channel: tickerChannel})

	<-done
}
