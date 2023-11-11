package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/kdyr/handlers"
	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/kdyr/services"
)

func NewRouter(svc *services.Services) *mux.Router {
	r := mux.NewRouter()
	apiV1 := r.PathPrefix("/api/v1").Subrouter()

	apiV1.HandleFunc("/sentence", handlers.NewGetRandomSentenceHandler(svc.SentenceSvc)).
		Methods(http.MethodGet)
	apiV1.HandleFunc("/sentence", handlers.NewAddSentenceHandler(svc.SentenceSvc)).
		Methods(http.MethodPost)
	apiV1.HandleFunc("/sentences/delete", handlers.NewDeleteSentencesHandler(svc.SentenceSvc)).
		Methods(http.MethodPost)

	return r
}
