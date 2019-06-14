package handler

import (
	"github.com/jaxmef/booklover-backend/models"
	"github.com/lancer-kit/armory/api/httpx"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/log"
	"net/http"
)

func RegisterAcc(w http.ResponseWriter, r *http.Request) {
	body := &models.Account{}
	err := httpx.ParseJSONBody(r, body)
	if err != nil {
		render.BadRequest(w, "invalid json")
		return
	}

	_, err = models.DB().Account().Insert(body)
	if err != nil {
		log.DefaultForRequest(r).WithError(err).
			Debugln("failed to insert account into db")
		render.ServerError(w)
		return
	}

	render.Success(w, nil)
}

func GetLastAccount(w http.ResponseWriter, r *http.Request) {
	acc, err := models.DB().Account().GetLast()
	if err != nil {
		log.DefaultForRequest(r).WithError(err).
			Debugln("failed to get last acc from db")
		render.ServerError(w)
		return
	}

	render.Success(w, acc)
}
