package app

import (
	"github.com/gorilla/mux"
	"github.com/monicaribeiro/event-manager/domain"
	"github.com/monicaribeiro/event-manager/service"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	eh := EventHandlers{service.NewEventService(domain.NewEventRepositoryDb())}

	router.HandleFunc("/events", eh.getAllEvents).Methods(http.MethodGet)
	router.HandleFunc("/events/{event_id:[0-9]+}", eh.deleteEvent).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
