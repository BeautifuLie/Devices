package handlers

import (
	"encoding/json"
	"fmt"

	"net/http"
	"program/joker"

	"github.com/gorilla/mux"
)

type apiHandler struct {
	Server *joker.Server
}

func RetHandler(server *joker.Server) *apiHandler {
	return &apiHandler{
		Server: server,
	}
}

func HandleRequest(h *apiHandler) *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/last", h.Last)
	myRouter.HandleFunc("/time", h.EventsT)

	return myRouter
}

func (h *apiHandler) Last(w http.ResponseWriter, r *http.Request) {

	_, str, err := h.Server.Last()

	if err != nil {
		fmt.Println("err")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *apiHandler) EventsT(w http.ResponseWriter, r *http.Request) {

	res := h.Server.EventsByTime()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
