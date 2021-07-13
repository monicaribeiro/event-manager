package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/monicaribeiro/event-manager/dto"
	"github.com/monicaribeiro/event-manager/service"
	"net/http"
	"strconv"
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
	events, appError := eh.service.GetAllEvents()

	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, events)
	}
}

func (eh *EventHandlers) deleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId, _ := strconv.ParseInt(vars["event_id"], 10, 64)

	event, appError := eh.service.Delete(eventId)

	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, event)
	}
}

func (eh *EventHandlers) createEvent(w http.ResponseWriter, r *http.Request) {
	var request dto.NewEventRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		appError := eh.service.Create(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, nil)
		}
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
