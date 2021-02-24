package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"price_loader/binance"
	"price_loader/poloniex"
	"price_loader/server"
	"price_loader/storage"
)

func main() {
	interruptAppChannel := make(chan os.Signal, 1)
	signal.Notify(interruptAppChannel, os.Interrupt)

	errorChannel := make(chan error)
	defer close(errorChannel)

	go func(ch chan error) {
		log.Fatal(<-ch)
	}(errorChannel)

	s := storage.GetKeyValueStorage()
	saveChannel := s.GetWriteChannel()
	defer close(saveChannel)

	go binance.PriceLoader(saveChannel, errorChannel)
	go poloniex.PriceLoader(saveChannel, errorChannel)

	server.RunServer(&s)

	<-interruptAppChannel
	fmt.Printf("\nBye!\n")
}
