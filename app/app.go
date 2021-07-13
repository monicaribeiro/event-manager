package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/monicaribeiro/event-manager/domain"
	"github.com/monicaribeiro/event-manager/logger"
	"github.com/monicaribeiro/event-manager/service"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_ADDRESS",
		"DB_PORT",
		"DB_NAME",
		"DB_SSLMODE",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	eh := EventHandlers{service.NewEventService(domain.NewEventRepositoryDb())}

	router.HandleFunc("/events", eh.getAllEvents).Methods(http.MethodGet)
	router.HandleFunc("/events", eh.createEvent).Methods(http.MethodPost)
	router.HandleFunc("/events/{event_id:[0-9]+}", eh.deleteEvent).Methods(http.MethodDelete)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
