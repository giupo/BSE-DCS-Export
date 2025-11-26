package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Vary", "Origin")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h(w, r)
	}
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func StartAPI(store *DataStore, port int) {
	http.HandleFunc("/health", withCORS(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]any{
			"missionRunning":       true,
			"missionServerRunning": true,
		}
		log.Printf("/health called\n")
		writeJSON(w, resp)
	}))

	http.HandleFunc("/mission-data", withCORS(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/mission-data called\n")
		store.mu.RLock()
		defer store.mu.RUnlock()
		writeJSON(w, store.Mission)
	}))

	http.HandleFunc("/position-player", withCORS(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/position-player called\n")
		store.mu.RLock()
		defer store.mu.RUnlock()
		writeJSON(w, store.Position)
	}))

	http.HandleFunc("/player-id", withCORS(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/player-id called\n")
		store.mu.RLock()
		defer store.mu.RUnlock()

		var val any
		if store.PlayerID != nil {
			if v, ok := store.PlayerID["playerId"]; ok {
				val = v
			}
		}
		writeJSON(w, val)
	}))

	http.HandleFunc("/export-world-objects", withCORS(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/export-world-objects\n")
		store.mu.RLock()
		defer store.mu.RUnlock()
		writeJSON(w, store.WorldObjects)
	}))

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Println("REST API in ascolto su", addr)
	http.ListenAndServe(addr, nil)
}
