package controllers

import (
	"encoding/json"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, req *http.Request) {
	p := map[string]int{"jerk": 1, "awesome": 2}
	json.NewEncoder(w).Encode(p)

}
