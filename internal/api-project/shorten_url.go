package apiproject

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type shortenURLRequest struct {
	URL string `json:"url"`
}

type shortenURLResponse struct {
	Code string `json:"code"`
}

func handleShortenURL(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body shortenURLRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJson(w, apiRespose{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			sendJson(w, apiRespose{Error: "invalid url passed"}, http.StatusBadRequest)
			return
		}

		code := genCode()
		db[code] = body.URL
		sendJson(w, apiRespose{Data: shortenURLResponse{Code: code}}, http.StatusCreated)
	}
}
