package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, req *http.Request) {
	p := map[string]interface{}{"jerk": 1, "awesome": 2, "name": "Eric Zhou"}
	jsonReturn(p, 1, "success", w)

}
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func jsonReturn(data interface{}, code uint8, msg string, w http.ResponseWriter) {
	p := map[string]interface{}{"code": code, "message": msg, "data": data}
	json.NewEncoder(w).Encode(p)
}
