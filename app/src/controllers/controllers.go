package controllers

import (
	"encoding/json"
	"net/http"

	"pic-collage.com/shorten_url/routes"
)

var Routes []*routes.Route

func init() {
	Routes = []*routes.Route{
		routes.NewRoute(http.MethodPost, "shorten", handleCreateShortenURL),
		routes.NewRoute(http.MethodGet, "{shorten-url}", handleRedirect),
	}
}

type errorResponse struct {
	Message string `json:"message"`
}

func returnOK(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeResponseJSON(w, resp)
}

func returnBadRequest(w http.ResponseWriter, errResp errorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	writeResponseJSON(w, errResp)
}

func returnForbiddenRequest(w http.ResponseWriter, errResp errorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	writeResponseJSON(w, errResp)
}

func returnNotFoundRequest(w http.ResponseWriter, errResp errorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	writeResponseJSON(w, errResp)
}

func returnInternalServerError(w http.ResponseWriter, errResp errorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	writeResponseJSON(w, errResp)
}

func writeResponseJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}
