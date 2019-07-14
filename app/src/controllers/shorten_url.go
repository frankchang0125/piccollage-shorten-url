package controllers

import (
	"encoding/json"
	"net/http"

	"pic-collage.com/shorten_url/models"
	"pic-collage.com/shorten_url/services"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type createShortenURLReq struct {
	URL string `json:"url"`
}

type createShortenURLResp struct {
	URL string `json:"url"`
}

func handleCreateShortenURL(w http.ResponseWriter, r *http.Request) {
	request := createShortenURLReq{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.WithError(err).Error("Fail to unmarshal user request")
		returnBadRequest(w, errorResponse{
			Message: "Invalid request body",
		})
		return
	}

	shorten, err := services.CreateShortenURL(request.URL)
	if err != nil {
		switch err {
		case models.ErrShortenURLNotExists:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	resp := createShortenURLResp{
		URL: shorten,
	}

	returnOK(w, resp)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shorten := vars["shorten-url"]
	if shorten == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := services.GetURL(shorten)
	if err != nil {
		switch err {
		case models.ErrShortenURLNotExists:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
