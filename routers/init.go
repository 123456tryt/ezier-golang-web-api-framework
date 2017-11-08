package routers

import (
	"my_go_web_framework/controllers"
	"net/http"
)

var HttpServeMux *http.ServeMux

func init() {
	HttpServeMux = http.NewServeMux()

	HttpServeMux.HandleFunc("/jerk", controllers.ApiHandler)

}
