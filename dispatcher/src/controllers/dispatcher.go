package controllers

import (
	"net/http"

	"pic-collage.com/dispatcher/services"
)

type dispatchResp struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}

func handleDispatchCounter(w http.ResponseWriter, r *http.Request) {
	start, end, err := services.DispatchCounter()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := dispatchResp{
		Start: start,
		End:   end,
	}
	returnOK(w, resp)
}
