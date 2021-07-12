package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/monicaribeiro/event-manager/service"
	"net/http"
)

type Event struct {
	Name          string `json:"name"`
	City          string `json:"city"`
	State         string `json:"state"`
	photoUrl      string `json:"photoUrl"`
	eventDateTime string `json:"eventDateTime"`
}

type EventHandlers struct {
	service service.EventService
}

func (eh *EventHandlers) getAllEvents(w http.ResponseWriter, r *http.Request) {
	events, _ := eh.service.GetAllEvents()
	writeResponse(w, http.StatusOK, events)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "CREATE EVENT: Hello world")
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["event_id"])
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
