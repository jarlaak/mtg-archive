package controllers

import (
	model "github.com/jarlaak/mtg-archive/models"
	"github.com/jarlaak/mtg-archive/server"
	"net/http"
)

type IdsParams struct {
	Ids []int `json:"id"`
}

func IdsHandler(w http.ResponseWriter, r *http.Request) {
	var params IdsParams
	err := server.ReadPostBody(r, &params)
	if err != nil {
		server.SendError(w, http.StatusBadRequest, "malformed JSON")
		return
	}
	if len(params.Ids) == 0 {
		server.SendError(w, http.StatusBadRequest, "no ids")
		return
	}

	cards := model.GetCards(params.Ids)
	server.SendJSON(w, cards)
}
func InitCardsController(r *server.AdvancedRouter) {
	cr := r.PathPrefix("/cards").Subrouter()

	cr.HandleFunc("/ids", IdsHandler)
}
