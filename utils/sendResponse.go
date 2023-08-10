package utils

import (
	"NEWGOLANG/config"
	"encoding/json"
	"net/http"
)

// RespondWithJSON ...
func RespondWithJSON(msg string, code int, w http.ResponseWriter, r *http.Request) {

	body := config.Body{ResponseCode: code, Message: msg}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if code == 400 {
		http.Error(w, msg, 400)
		return
	}
	w.Write(jsonBody)

}
