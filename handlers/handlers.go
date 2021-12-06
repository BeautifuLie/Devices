package handlers

import (
	"encoding/json"
	"fmt"

	"net/http"
	"program/server"

	"github.com/gorilla/mux"
)

type apiHandler struct {
	Server *server.Server
}

func RetHandler(server *server.Server) *apiHandler {
	return &apiHandler{
		Server: server,
	}
}

func HandleRequest(h *apiHandler) *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/last", h.LastTenEvents)
	myRouter.HandleFunc("/time", h.EventsTime)

	return myRouter
}

func (h *apiHandler) LastTenEvents(w http.ResponseWriter, r *http.Request) {

	res, str, err := h.Server.Last()

	if err != nil {
		fmt.Println("err")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *apiHandler) EventsTime(w http.ResponseWriter, r *http.Request) {

	res := h.Server.EventsByTime()

	w.Header().Set("Content-Type", "application/json")

	e, _ := json.MarshalIndent(res, "/", "   ")
	// err := json.NewEncoder(w).Encode(e)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	w.Write(e)
}
