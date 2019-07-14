package routes

import (
	"net/http"
)

type Route struct {
	Method   string
	Endpoint string
	Handler  func(http.ResponseWriter, *http.Request)
}

// NewRoute returns the route with default api version: 1.
func NewRoute(method string, endpoint string,
	handler func(http.ResponseWriter, *http.Request)) *Route {
	return &Route{
		Method:   method,
		Endpoint: endpoint,
		Handler:  handler,
	}
}
