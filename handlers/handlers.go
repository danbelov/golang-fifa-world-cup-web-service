package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

const AccessTokenHeader = "X-ACCESS-TOKEN"

// RootHandler returns an empty body status code
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// ListWinners returns winners from the list
func ListWinners(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	year := r.URL.Query().Get("year")
	if year == "" {
		wnr, err := data.ListAllJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(wnr)
	} else {
		fwn, err := data.ListAllByYear(year)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write(fwn)
	}
}

// AddNewWinner adds new winner to the list
func AddNewWinner(w http.ResponseWriter, r *http.Request) {
	tk := r.Header.Get(AccessTokenHeader)
	ok := data.IsAccessTokenValid(tk)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := data.AddNewWinner(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {

}
