package controllers

import (
	"encoding/json"
	"ezier/models"
	"fmt"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, req *http.Request) {

	agents := models.GetAllAgents()
	jsonReturn(agents, 1, "success", w)

}
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func jsonReturn(data interface{}, code uint8, msg string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	p := map[string]interface{}{"code": code, "message": msg, "data": data}
	json.NewEncoder(w).Encode(p)
}
