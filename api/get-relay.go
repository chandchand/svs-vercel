package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RelayGet struct {
	Contact int `json:"contact"`
	Engine  int `json:"engine"`
	Key     int `json:"key"`
}

const firebaseURL = "https://smartvehiclesentinel-2ed68-default-rtdb.asia-southeast1.firebasedatabase.app/relay.json"

func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	resp, err := http.Get(firebaseURL)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Failed to fetch from Firebase", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var relay RelayGet
	if err := json.Unmarshal(body, &relay); err != nil {
		http.Error(w, "Failed to parse Firebase response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(relay)
}
