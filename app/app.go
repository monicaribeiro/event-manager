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

	eh := EventHandlers{service.NewEventService(domain.NewEventRepositoryStub())}

	router.HandleFunc("/events", eh.getAllEvents).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
