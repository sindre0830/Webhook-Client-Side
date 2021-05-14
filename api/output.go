package api

import (
	"encoding/json"
	"main/debug"
	"net/http"
)

// OutputHandler will output inputLog to client.
func OutputHandler(w http.ResponseWriter, r *http.Request) {
	if len(inputLog) < 1 {
		http.Error(w, "", http.StatusNoContent)
		return
	}
	//update header to JSON and set HTTP code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//send output to user and branch if an error occured
	err := json.NewEncoder(w).Encode(inputLog)
	if err != nil {
		debug.ErrorMessage.Update(
			http.StatusInternalServerError, 
			"OutputHandler() -> Sending data to user",
			err.Error(),
			"Unknown",
		)
		debug.ErrorMessage.Print(w)
	}
}
