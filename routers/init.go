package routers

import (
	"ezier/controllers"
	"net/http"
)

var HttpServeMux *http.ServeMux

//其实这个可以省略调controller
func init() {
	HttpServeMux = http.NewServeMux()

	HttpServeMux.HandleFunc("/jerk", controllers.ApiHandler)

}
