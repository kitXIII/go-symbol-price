package main

import (
	"log"

	"price_loader/binance"
	"price_loader/poloniex"
	"price_loader/server"
	"price_loader/storage"
)

func main() {
	errorChannel := make(chan error)
	defer close(errorChannel)

	go func(ch chan error) {
		log.Fatal(<-ch)
	}(errorChannel)

	s := storage.GetKeyValueStorage()
	saveChannel := s.GetWriteChannel()
	defer close(saveChannel)

	go binance.PriceLoader(&saveChannel, &errorChannel)
	go poloniex.PriceLoader(&saveChannel, &errorChannel)

	server.RunServer(&s)
}
