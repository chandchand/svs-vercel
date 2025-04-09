package getrelay

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Relay model
type Relay struct {
	Contact int `json:"contact"`
	Engine  int `json:"engine"`
	Key     int `json:"key"`
}

const firebaseURL = "https://smartvehiclesentinel-2ed68-default-rtdb.asia-southeast1.firebasedatabase.app/relay.json"

// Handler is the entrypoint function Vercel expects
func Handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(firebaseURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get data: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var relay Relay
	if err := json.NewDecoder(resp.Body).Decode(&relay); err != nil {
		http.Error(w, "Failed to parse response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(relay)
}
