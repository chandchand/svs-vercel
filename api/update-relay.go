package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const updateURL = "https://smartvehiclesentinel-2ed68-default-rtdb.asia-southeast1.firebasedatabase.app/relay.json"

type RelayUpdate struct {
	Contact int `json:"contact"`
	Engine  int `json:"engine"`
	Key     int `json:"key"`
}

func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Only PATCH allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(http.MethodPatch, updateURL, strings.NewReader(string(body)))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to update Firebase", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Update success",
	})
}
