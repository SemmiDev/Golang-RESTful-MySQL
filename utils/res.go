package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, data interface{}, status int) {
	changeToByte, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "error nih", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(changeToByte))
}
