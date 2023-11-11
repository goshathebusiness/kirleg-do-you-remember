package rest

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, code int, response any) {
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}

func RespondError(w http.ResponseWriter, code int, err error) {
	RespondJSON(w, code, &Error{
		Code:    code,
		Message: err.Error(),
	})
}
