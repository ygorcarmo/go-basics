package api

import (
	"encoding/json"
	"example/museum/data"
	"fmt"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if &id != nil {
		finalId, err := strconv.Atoi(id)
		if err == nil && (finalId > len(data.GetAll())-1) {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else {
			fmt.Println(len(data.GetAll()))
			http.Error(w, "Invalid Request", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
