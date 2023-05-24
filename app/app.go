package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8081", muxRouter))
}
