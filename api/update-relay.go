// api/update-relay.go
package updaterelay

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Relay struct {
	Contact int `json:"contact"`
	Engine  int `json:"engine"`
	Key     int `json:"key"`
}

const firebaseURL = "https://smartvehiclesentinel-2ed68-default-rtdb.asia-southeast1.firebasedatabase.app/relay.json"

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var relay Relay
	if err := json.NewDecoder(r.Body).Decode(&relay); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	payload, _ := json.Marshal(relay)
	req, err := http.NewRequest(http.MethodPatch, firebaseURL, strings.NewReader(string(payload)))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	io.Copy(w, resp.Body) // Optional: bisa juga return message sendiri
}
