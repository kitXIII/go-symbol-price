package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

// KeyValueStorage - simple storage
type KeyValueStorage struct {
	sync.Mutex
	data map[string]float64
}

// SymbolPrice - struct for save
type SymbolPrice struct {
	Symbol string
	Price  float64
}

// GetKeyValueStorage - KeyValueStorage factory
func GetKeyValueStorage() KeyValueStorage {
	data := make(map[string]float64)
	return KeyValueStorage{data: data}
}

func (k *KeyValueStorage) set(key string, value float64) {
	k.data[key] = value
}

func (k *KeyValueStorage) get(key string) (float64, error) {
	value, ok := k.data[key]
	if !ok {
		errorMessage := fmt.Sprintf("There is no value for key %s", key)
		return 0, errors.New(errorMessage)
	}

	return value, nil
}

// Set value by key with lock by mutex
func (k *KeyValueStorage) Set(key string, value float64) {
	k.Lock()
	defer k.Unlock()

	k.set(key, value)
}

// Get value by key with lock by mutex
func (k *KeyValueStorage) Get(key string) (float64, error) {
	k.Lock()
	defer k.Unlock()

	return k.get(key)
}

// GetWriteChannel generate channel for write keyValue data
func (k *KeyValueStorage) GetWriteChannel() chan SymbolPrice {
	ch := make(chan SymbolPrice)

	go func() {
		for {
			symbolPrice := <-ch
			k.Set(symbolPrice.Symbol, symbolPrice.Price)
		}
	}()

	return ch
}

// GetJsonData return serialized storage as string
func (k *KeyValueStorage) GetJsonData() ([]byte, error) {
	data := k.data

	return json.Marshal(data)
}
