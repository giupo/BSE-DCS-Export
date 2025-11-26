package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type DataStore struct {
	mu           sync.RWMutex
	Mission      map[string]any
	Position     map[string]any
	PlayerID     map[string]any
	WorldObjects map[string]any
}

func NewDataStore() *DataStore {
	return &DataStore{
		Mission:      make(map[string]any),
		Position:     make(map[string]any),
		PlayerID:     make(map[string]any),
		WorldObjects: make(map[string]any),
	}
}

func updateData(store *DataStore, data []byte) {
	var parsed map[string]any
	if err := json.Unmarshal(data, &parsed); err != nil {
		fmt.Printf("JSON non valido: %s\n", err)
		return
	}

	// esempio: puoi popolare il datastore con chiavi note
	store.mu.Lock()
	defer store.mu.Unlock()
	if _, ok := parsed["mission"]; ok {
		store.Mission = parsed
	} else if _, ok := parsed["position"]; ok {
		store.Position = parsed
	} else if _, ok := parsed["playerId"]; ok {
		store.PlayerID = parsed
	} else if _, ok := parsed["worldObjects"]; ok {
		store.WorldObjects = parsed
	}
}
