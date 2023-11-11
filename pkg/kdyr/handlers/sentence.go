package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/kdyr/services"
	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/models"
	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/rest"
)

func NewGetRandomSentenceHandler(svc services.SentenceSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sentence, err := svc.GetRandomSentence(r.Context())
		if err != nil {
			rest.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		rest.RespondJSON(w, http.StatusOK, sentence)
	}
}

func NewAddSentenceHandler(svc services.SentenceSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			rest.RespondError(w, http.StatusBadRequest, err)
			return
		}
		defer r.Body.Close()

		var sentence *models.Sentence
		err = json.Unmarshal(b, &sentence)
		if err != nil {
			rest.RespondError(w, http.StatusBadRequest, err)
			return
		}

		err = svc.AddSentence(r.Context(), sentence)
		if err != nil {
			rest.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		rest.RespondJSON(w, http.StatusOK, sentence)
	}
}

func NewDeleteSentencesHandler(svc services.SentenceSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			rest.RespondError(w, http.StatusBadRequest, err)
			return
		}
		defer r.Body.Close()

		var ids []uint64
		err = json.Unmarshal(b, &ids)
		if err != nil {
			rest.RespondError(w, http.StatusBadRequest, err)
			return
		}

		err = svc.DeleteSentenceByIDs(r.Context(), ids)
		if err != nil {
			rest.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
